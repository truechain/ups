package upsstore

import (
	"bytes"
	// "context"
	// "errors"
	"fmt"
	"os"
	"bufio"
	"path/filepath"
	// "math/big"
	// "time"
	// "github.com/truechain/ups/accounts"
	// "github.com/truechain/ups/accounts/keystore"
	"github.com/truechain/ups/common"
	// "github.com/truechain/ups/common/hexutil"
	// "github.com/truechain/ups/common/math"
	// "github.com/truechain/ups/core"
	// "github.com/truechain/ups/core/rawdb"
	// "github.com/truechain/ups/core/types"
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
}
func (u *UpsFile) Event() {
}
func (u *UpsFile) Wait() {

}
func (u *UpsFile) setFileHash(hash string) {
	u.hash = hash
}

type FileMgr struct {
	
}

func cacheFileToHard(cfg *ipfsConfig,file *UpsFile) error {
	if cfg == nil {
		cfg = getDefaultIpfsConfig()
	}
	filename := filepath.Join(cfg.dir,file.name)
	dstFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Error("open file failed", "name", filename, "err", err)
		return err
	}
	bufWriter := bufio.NewWriter(dstFile)
	defer func() {
		bufWriter.Flush()
		dstFile.Close()
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
func AddFile(file *UpsFile) error {
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
	file.Event()
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
	file.Event()
	return nil
}