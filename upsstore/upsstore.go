package upsstore

import (
	"bytes"
	"sync"
	"errors"
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"bufio"
	"path/filepath"
	// "math/big"
	// "time"
	"github.com/truechain/ups/common"
	"github.com/truechain/ups/core/types"
	// "github.com/truechain/ups/crypto"
	"github.com/truechain/ups/log"
	// "github.com/truechain/ups/params"
	// "github.com/truechain/ups/rlp"
	shell "github.com/ipfs/go-ipfs-api"
)
var upsFileMgr = NewFileMgr()

func init() {
	cfg := getDefaultIpfsConfig()
	upsFileMgr.LoadFromCache(cfg)
}

func FileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func isDir(dir string) bool {
	info, err := os.Stat(dir)
	if err == nil {
		return false
	}
	return info.IsDir()
}
func isFile(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

type ipfsConfig struct {
	url 	string
	dir 	string
}
func getDefaultIpfsConfig() *ipfsConfig {
	return &ipfsConfig{
		url:	"localhost:5001",
		dir:	"./cacheFile",
	}
}
type UpsFile struct {
	data 	[]byte
	name 	string
	hash 	string
	address common.Address
	wg 		*sync.WaitGroup
	cache 	bool 
	fileHash common.Hash
}
func NewUpsFile(name string,addr common.Address,data []byte) *UpsFile {
	return &UpsFile{
		name:	name,
		address: addr,
		data:	data,
		wg:		&sync.WaitGroup{},
		cache:	false,
	}
}
func (u *UpsFile) UpdateFileHash() *UpsFile {
	u.fileHash = types.RlpHash([]interface{}{
		u.name,
		u.data,
	})
	return u
}
func (u *UpsFile) GetFileHash() common.Hash {
	empty := common.Hash{}
	if u.fileHash == empty {
		u.UpdateFileHash()
	}
	return u.fileHash
}

func (u *UpsFile) Finish() {
	u.wg.Done()
}
func (u *UpsFile) Wait() {
	u.wg.Wait()
}
func (u *UpsFile) Event() *UpsFile {
	u.wg.Add(1)
	return u
}
func (u *UpsFile) getFileNameInCache(cfg *ipfsConfig) string {
	if cfg == nil {
		cfg = getDefaultIpfsConfig()
	}
	filename := u.name + "_" + u.hash + "_" + u.address.String()
	filename = filepath.Join(cfg.dir,filename)
	return filename
}
func (u *UpsFile) setFileHashCode(hash string) {
	u.hash = hash
}
func (u *UpsFile) GetFileHashCode() string {
	return u.hash
}
func (u *UpsFile) equal(o *UpsFile) bool {
	return u.GetFileHash() == o.GetFileHash()
}
func (u *UpsFile) fileCache(c bool) {
	u.cache = c
}
func (u *UpsFile) isFileCache() bool {
	return u.cache
}
func (u *UpsFile) baseCopy(o *UpsFile) {
	u.data,u.name,u.hash = o.data,o.name,o.hash
	u.cache,u.fileHash = o.cache,o.fileHash
}
func (u *UpsFile) GetData() []byte {
	return u.data
}


type FileMgr struct {
	caches 	map[common.Hash][]*UpsFile	
	max 	int
	lock 	sync.RWMutex
}
func NewFileMgr() *FileMgr {
	return &FileMgr{
		max:	10,
		caches: make(map[common.Hash][]*UpsFile),
	}
} 
func GetGlobalFileMgr() *FileMgr {
	return upsFileMgr
}
func (m *FileMgr) LoadFromCache(cfg *ipfsConfig) error {
	return filepath.Walk(cfg.dir,func(path string, info os.FileInfo, err error) error{
		if err != nil {
            return err
		}
		if info.IsDir() {
			return nil
		}
		filename := filepath.Base(path)
		strs := strings.Split(filename,"_")
		if len(strs) != 3 {
			return errors.New("wrong format of the file name")
		}
		name,hashcode,hexAddr := strs[0],strs[1],strs[2]
		
		if f, err := os.Open(filename); err != nil {
			return err
		} else {
			defer f.Close()
			if fd, err := ioutil.ReadAll(f); err != nil {
				return err
			} else {
				uf := NewUpsFile(name,common.HexToAddress(hexAddr),fd)
				uf.setFileHashCode(hashcode)
				m.addFile(uf)
			}
		}
		return nil
	})
}
func (m *FileMgr) FileExist(hash common.Hash) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_,ok := m.caches[hash]
	return ok
}
func (m *FileMgr) GetFileByHashCode(hash string) *UpsFile {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _,val := range m.caches {
		v := val[0]
		if v.hash == hash {
			return v
		}
	}
	return nil
}
func (m *FileMgr) GetFile(h common.Hash,addr common.Address) (*UpsFile,bool){
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val,ok := m.caches[h]; ok {
		for _,v :=range val {
			if v.address == addr {
				return v,true
			}
		}
		return val[0],false
	}
	return nil,false
}
func (m *FileMgr) addFile(o *UpsFile) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val,ok := m.caches[o.GetFileHash()]; ok {
		val = append(val,o)
		m.caches[o.GetFileHash()] = val
	} else {
		m.caches[o.GetFileHash()] = []*UpsFile{o}
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////
// filename: name_codehash_address
func cacheFileToHard(cfg *ipfsConfig,file *UpsFile) error {
	if file == nil {
		return errors.New("file is nil")
	}
	if file.isFileCache() {
		return nil
	}
	filename := file.getFileNameInCache(cfg)
	dstFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Error("open file failed", "name", filename, "err", err)
		return err
	}
	bufWriter := bufio.NewWriter(dstFile)
	defer func() {
		bufWriter.Flush()
		dstFile.Close()
		file.fileCache(true)
	}()
	bufWriter.Write(file.data)
	return nil
}
func removeFileInCache(cfg *ipfsConfig,name string) error {
	if cfg == nil {
		cfg = getDefaultIpfsConfig()
	}
	filename := filepath.Join(cfg.dir,name)
	os.Remove(filename)
	return nil
}
func executeUpload(cfg *ipfsConfig,file *UpsFile) error {
	fmt.Println("begin upload....")
	if cfg == nil {
		cfg = getDefaultIpfsConfig()
	}
	sh := shell.NewShell(cfg.url)
	cid, err := sh.Add(bytes.NewReader(file.data))
	if err != nil {
		log.Error("executeUpload", "name", file.name, "err", err)
		return err
	}
	file.setFileHashCode(cid)
	file.Finish()
	return nil
}
func executeDownload(cfg *ipfsConfig,file *UpsFile) error {
	if cfg == nil {
		cfg = getDefaultIpfsConfig()
	}
	sh := shell.NewShell(cfg.url)
	read, err := sh.Cat(file.hash)
	if err != nil {
		log.Error("executeDownload", "hash", file.hash, "err", err)
		return err
	}
	defer read.Close()
	data, err := ioutil.ReadAll(read)
	if err != nil {
		log.Error("executeDownload:ReadAll", "hash", file.hash, "err", err)
		return err
	}
	file.data = data
	file.Finish()
	return nil
}
/////////////////////////////////////////////////////////////////////////////////
func AddFile(file *UpsFile) error {
	mgr := GetGlobalFileMgr()
	f1,user := mgr.GetFile(file.GetFileHash(),file.address)
	if f1 != nil {
		file.baseCopy(f1)
		if !user {
			mgr.addFile(file)
		} 
	} else {
		// 1. encryption the file
		// 2. cache the file in the local node
		// 3. upload the file to ipfs
		cfg := getDefaultIpfsConfig()
		var err error
		file.Event()
		go func(){
			err = executeUpload(cfg,file)
		}()
		file.Wait()
		fmt.Println("finish wait...")
		if err != nil {
			return err
		}
		go func() error {
			if err := cacheFileToHard(cfg,file); err != nil {
				return err
			}
			return nil
		}()
	}
	return nil
}
func GetFile(name,hash string,addr common.Address) (*UpsFile,error) {
	// 1. check the file in the cache with the file's hash
	// 2. encryption the file
	// 3. get file from ipfs
	mgr := GetGlobalFileMgr()
	f := mgr.GetFileByHashCode(hash)
	if f != nil {
		return f,nil
	} else {
		file := NewUpsFile(name,addr,nil)
		file.setFileHashCode(hash)
		cfg := getDefaultIpfsConfig()
		var err error
		file.Event()
		go func(){
			err = executeUpload(cfg,file)
		}()
		file.Wait()
		if err != nil {
			return nil,err
		}
		mgr.addFile(file.UpdateFileHash())
		return file,nil
	}
}

