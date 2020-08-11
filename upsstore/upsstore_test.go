package upsstore

import (
	"fmt"
	"testing"
	// "math/big"
	"math/rand"
    "time"
	"github.com/truechain/ups/common"
	// "github.com/truechain/ups/core/types"
	// "github.com/truechain/ups/crypto"
	// "github.com/truechain/ups/log"
	// "github.com/truechain/ups/params"
	// "github.com/truechain/ups/rlp"
	// shell "github.com/ipfs/go-ipfs-api"
)

func get_file(name string,addr common.Address) *UpsFile {
	rand.Seed(time.Now().Unix())
	data := []byte{}
	for i:=0;i<2000;i++ {
		data = append(data,byte(rand.Intn(255)))
	}
	return NewUpsFile(name,addr,data)
}
func Test_01(t *testing.T) {
	name := "test1"
	addr := common.BytesToAddress([]byte{0,0,0,0,1})
	uf := get_file(name,addr)
	err := AddFile(uf)
	time.Sleep(5 * time.Second)
	if err != nil {
		fmt.Println("err",err)
		return 
	}
	hashcode := uf.GetFileHashCode()
	fmt.Println("hashcode",hashcode)

	if uf2,err := GetFile(name,hashcode,addr); err != nil {
		fmt.Println("err",err)
	} else {
		fmt.Println("name:",uf2.name)
	}
}

func run_01() {
	fmt.Println("begin upload....")
	// if cfg == nil {
	// 	cfg = getDefaultIpfsConfig()
	// }
	// file.Finish()
	// sh := shell.NewShell(cfg.url)
	// cid, err := sh.Add(bytes.NewReader(file.data))
	// if err != nil {
	// 	log.Error("executeUpload", "name", file.name, "err", err)
	// 	return err
	// }
	// file.setFileHashCode(cid)
	fmt.Println("end upload....")
}
func Test_02(t *testing.T) {
	fmt.Println("begin wait...")
	go run_01()
	time.Sleep(5 * time.Second)
	fmt.Println("finish wait...")
}