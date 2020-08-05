package upsstore

import (
	"bytes"
	"sync"
	"errors"
	"fmt"
	"os"
	"bufio"
	"path/filepath"
	// "math/big"
	// "time"
	"github.com/truechain/ups/common"
	// "github.com/truechain/ups/common/hexutil"
	// "github.com/truechain/ups/common/math"
	// "github.com/truechain/ups/core"
	// "github.com/truechain/ups/core/rawdb"
	"github.com/truechain/ups/core/types"
	// "github.com/truechain/ups/core/vm"
	// "github.com/truechain/ups/crypto"
	"github.com/truechain/ups/log"
	// "github.com/truechain/ups/p2p"
	// "github.com/truechain/ups/params"
	// "github.com/truechain/ups/rlp"
	// "github.com/truechain/ups/rpc"
	shell "github.com/ipfs/go-ipfs-api"
)

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
func (u *UpsFile) UpdataFileHash() {
	u.fileHash = types.RlpHash([]interface{}{
		u.name,
		u.data,
	})
}
func (u *UpsFile) GetFileHash() common.Hash {
	if u.fileHash == common.Hash{} {
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
	filename := filepath.Join(cfg.dir,file.name)
	return filename
}
func (u *UpsFile) setFileHash(hash string) {
	u.hash = hash
}
func (u *UpsFile) equal(o *UpsFile) bool {
	return u.GetFileHash() == o.GetFileHash()
}
func (u *UpsFile) fileCache(c bool) {
	u.cache = c
}
func (u *UpsFile) baseCopy(o *UpsFile) {
	u.data,u.name,u.hash = o.data,o.name,o.hash
	u.cache,u.fileHash = o.cache,o.fileHash
}


type FileMgr struct {
	caches 	map[common.Hash][]*UpsFile	
	max 	int
}
func NewFileMgr() *FileMgr {
	return &FileMgr{
		max:	10,
		caches: make(map[common.Hash][]*UpsFile),
	}
} 
func GetGlobalFileMgr() *FileMgr {
	return nil
}
func (m *FileMgr) LoadFromCache(cfg *ipfsConfig) error {
	return nil
}
func (m *FileMgr) FileExist(hash common.Hash) bool {
	return false
}
func (m *FileMgr) GetFileByHashCode(hash string) *UpsFile {
	for _,val := range m.caches {
		v := val[0]
		if v.hash == hash {
			return v
		}
	}
	return nil
}
func (m *FileMgr) GetFile(h common.Hash,addr common.Address) (*UpsFile,bool){
	if val,ok := m.caches[h]; ok {
		for _,v :=range val {
			if v.address == addr {
				return v,true
			}
		}
		return v,false
	}
	return nil,false
}
func (m *FileMgr) addFile(o *UpsFile) error {
	if val,ok := m.caches[o.GetFileHash()]; ok {
		val = append(val,o)
		m.caches[o.GetFileHash()] = val
	} else {
		m.caches[o.GetFileHash()] = []*UpsFile{o}
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////
func cacheFileToHard(cfg *ipfsConfig,file *UpsFile) error {
	if file == nil {
		return errors.New("file is nil")
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
	filename := filepath.Join(cfg.dir,file.name)
	os.Remove(filename)
	return nil
}
func executeUpload(cfg *ipfsConfig,file *UpsFile) error {
	if cfg == nil {
		cfg = getDefaultIpfsConfig()
	}
	sh := shell.NewShell(cfg.url)
	cid, err := sh.Add(bytes.NewReader(file.data))
	if err != nil {
		log.Error("executeUpload", "name", filename, "err", err)
		return err
	}
	file.setFileHash(cid)
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
	f1,user := mgr.GetFile(file.GetFileHash())
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
		if err := cacheFileToHard(cfg,file); err != nil {
			return err
		}
		go func(){
			executeUpload(cfg,file.Event())
		}()
		file.Wait()
	}
	return nil
}
func GetFile(name,hash string,addr common.Address) *UpsFile {
	
	return nil
}

