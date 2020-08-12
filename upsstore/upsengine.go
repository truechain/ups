package upsstore

import (
	// "fmt"
	"math/big"
	"crypto/rand"
	// "time"
	"bytes"
	"errors"
	"github.com/truechain/ups/common"
	// "github.com/truechain/ups/core/state"
	// "github.com/truechain/ups/core/types"
	"github.com/truechain/ups/crypto"
	"github.com/truechain/ups/crypto/ecies"
	// "github.com/truechain/ups/log"
	// "github.com/truechain/ups/params"
	// "github.com/truechain/ups/rlp"
	// shell "github.com/ipfs/go-ipfs-api"
)

// 交互式密钥交互过程,卖方设置价格,买方出价，卖方给出加密密钥(以买方公钥加密)，验证该加密密钥。
// 本服务为单向服务合约

var (
	UpsEngineAddress = common.BytesToAddress([]byte{90})
	CleanDealFlag = false
)

var (
	ErrNotFoundProvider = errors.New("not found the provider from the key")
	ErrNotMatchPrice = errors.New("not match the price from the provider")
	ErrInvalidParams = errors.New("invalid params")
)

func matchPrice(p *provider,c *consumer) error {
	if p == nil || c == nil {
		return ErrInvalidParams
	}
	e := p.getService(c.getKey())
	if e == nil {
		return ErrNotFoundProvider
	}
	if e.getPrice().Cmp(c.getPrice()) == 0 {
		return nil
 	} else {
		 return ErrNotMatchPrice
	}
}
// will be store to state
func storeBalance(val *big.Int) error {
	// add balance to state
	return nil
}
func transToProvider(val *big.Int) error {
	// trans to the provider from the upsEngineAddress 
	return nil
} 
func allBalance() *big.Int {
	// all balance of the upsEngineAddress
	return nil
}
func Encrypt(pk,msg []byte) ([]byte,error) {
	if p,err := crypto.UnmarshalPubkey(pk); err != nil {
		return nil,err
	} else {
		if cr, err := ecies.Encrypt(rand.Reader, ecies.ImportECDSAPublic(p), msg, nil, nil); err != nil {
			return nil,err
		} else {
			return cr,nil
		}
	}
}
func Decrypt(priv,msg []byte) ([]byte,error) {
	if p,err := crypto.ToECDSA(priv);err != nil {
		return nil,err
	} else {
		priKey := ecies.ImportECDSA(p)
		if dcr, err := priKey.Decrypt(msg, nil, nil); err != nil {
			return nil,err
		} else {
			return dcr,nil
		}
	}
}

type deal struct {
	key string
	buyer common.Address
	buyerPk []byte
	password []byte
	Height 	uint64
}
func (d *deal) MakePassword() error {
	
	return nil
}

type Entry struct {
	key string
	price *big.Int
	buyer bool
}
func (e *Entry) getPrice() *big.Int {
	return e.price
}

type provider struct {
	Addr 		common.Address
	Service 	map[string]*Entry
	DealList 	[]*deal
}
func (p *provider) getService(key string) *Entry {
	v,ok := p.Service[key]
	if !ok {
		return nil
	}
	return v
}
func (p *provider) getPrice(key string) *big.Int {
	e := p.getService(key)
	if e == nil {
		return nil
	}
	return e.getPrice()
}
func (p *provider) getDealResultBy(key string,addr common.Address) *deal {
	for _,v := range p.DealList {
		if key == v.key && bytes.Equal(addr[:],v.buyer[:]) {
			return v
		}
	}
	return nil
}

type consumer struct {
	Key 	string
	Price   *big.Int
	Pk 		[]byte	
}
func (c *consumer) getKey() string {
	return c.Key
}
func (c *consumer) getPrice() *big.Int {
	return new(big.Int).Set(c.Price)
}

type Engine struct {
	
}

func (en *Engine) getProviderByKey(key string) *provider {
	return nil
}
func (en *Engine) Seller() {
	
}

func (en *Engine) LockedAmount() error {
	return nil
}
func (en *Engine) UnlockAmount() error {
	return nil
}
func (en *Engine) matchByConsumer(c *consumer) (*big.Int,error) {
	p := en.getProviderByKey(c.getKey())
	if p == nil {
		return nil,ErrNotFoundProvider
	}
	err := matchPrice(p,c)
	
	return c.getPrice(),err
}
// try to add balance to contract address and locked it
func (en *Engine) tryAddConsumer(c *consumer) error {
	// 1. add consumer and match the provider
	// 2. the consumer store the money to the contract and locked it
	// 3. the provider set the password encrypted by consumer's pk
	
	return nil
}




