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
	MaxAutoRedeemHeight = 200
	MaxAutoUnlockedHieght = 500
)

var (
	ErrNotFoundProvider = errors.New("not found the provider from the key")
	ErrNotFoundDeal = errors.New("not found the deal in the provider")
	ErrNotMatchPrice = errors.New("not match the price from the provider")
	ErrInvalidParams = errors.New("invalid params")
	ErrInvalidPk = errors.New("uninitialized pubkey in deal")
)
const (
	DealUnPaid byte = iota
	DealPayed
	DealRedeem
	DealFinish
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
func transToProvider(val *big.Int,addr common.Address) error {
	// trans to the provider from the upsEngineAddress 
	return nil
} 
func redeemToConsumer(val *big.Int,addr common.Address) error {
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
func MakePassword() error {
	
	return nil
}

type deal struct {
	key string
	buyer common.Address
	buyerPk []byte
	password []byte
	Height 	uint64
	price 	*big.Int
	state byte
}
func (d *deal) setPassword(ec []byte) {
	d.password = make([]byte,len(ec))
	copy(d.password,ec)
}
func (d *deal) getPk() []byte{
	if len(d.buyerPk) > 0 {
		pk := make([]byte,len(d.buyerPk))
		copy(pk,d.buyerPk)
		return pk
	}
	return nil
}
func (d *deal) getHeight() uint64 {
	return d.Height
}
func (d *deal) getPrice() *big.Int {
	return new(big.Int).Set(d.price)
}
func (d *deal) isPayed() bool {
	return d.state == DealPayed
}
func (d *deal) canRedeem(ch uint64) bool {
	if int64(ch - d.Height) > int64(MaxAutoRedeemHeight) && d.state == DealUnPaid {
		return true
	}
	return false
}
func (d *deal) redeemed() {
	d.state = DealRedeem
}
func (d *deal) payed() {
	d.state = DealPayed
}
func (d *deal) finish() {
	d.state = DealFinish
}
func (d *deal) doRedeem() error {
	return redeemToConsumer(d.price,d.buyer)
}

type Entry struct {
	key string
	price *big.Int
	buyer bool
}
func newEntry(key string, price *big.Int) *Entry {
	return &Entry{
		key:	key,
		price: new(big.Int).Set(price),
	}
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
func (p *provider) getDealResult(key string,addr common.Address) *deal {
	for _,v := range p.DealList {
		if key == v.key && bytes.Equal(addr[:],v.buyer[:]) {
			return v
		}
	}
	return nil
}
func (p *provider) addDealResult(height uint64,key string,addr common.Address,pk []byte,price *big.Int) {
	d := p.getDealResult(key,addr)
	if d != nil {
		return 
	}
	p.DealList = append(p.DealList,&deal{
		key:	key,
		buyer:	addr,
		buyerPk:	pk,
		Height:	height,
		price:	new(big.Int).Set(price),
		state:	DealUnPaid,
	})
	return 
}
func (p *provider) setPassword(key string, addr common.Address,ec []byte) error {
	d := p.getDealResult(key,addr)
	if d == nil {
		return ErrNotFoundDeal
	}
	d.setPassword(ec)
	d.payed()
	return nil
}
func (p *provider) addEntry(key string,price *big.Int) error {
	// disable to change the price 
	_,ok := p.Service[key]
	if !ok {
		p.Service[key] = newEntry(key,price)
	}
	return nil
}
func (p *provider) getAddress() common.Address {
	return p.Addr
}
func (p *provider) isSuppy(key string) bool {
	for k := range p.Service {
		if k == key {
			return true
		}
	}
	return false
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
func (c *consumer) getAddr() common.Address {
	return common.BytesToAddress(crypto.Keccak256(c.Pk[1:])[12:])
}

type Engine struct {
	maker  map[common.Address]*provider
}

func (en *Engine) getProviderByKey(key string) *provider {
	for _,v := range en.maker {
		if v.isSuppy(key) {
			return v
		}
	}
	return nil
}
func (en *Engine) getProviderByAddr(own common.Address) *provider {
	v,ok := en.maker[own]
	if ok {
		return v
	}
	return nil
}

func (en *Engine) LockedAmount() error {
	return nil
}
func (en *Engine) UnlockAmount() error {
	return nil
}
func (en *Engine) addProvider(p *provider) {
	en.maker[p.getAddress()] = p
}

func (en *Engine) matchByConsumer(c *consumer) (*provider,error) {
	p := en.getProviderByKey(c.getKey())
	if p == nil {
		return nil,ErrNotFoundProvider
	}
	err := matchPrice(p,c)
	
	return p,err
}

// try to add balance to contract address and locked it
func (en *Engine) tryAddConsumer(c *consumer,height uint64) error {
	// 1. add consumer and match the provider
	// 2. the consumer store the money to the contract and locked it
	// 3. the provider set the password encrypted by consumer's pk
	p,err := en.matchByConsumer(c)
	if err != nil {
		return err
	}
	p.addDealResult(height,c.Key,c.getAddr(),c.Pk,c.getPrice())
	return nil
}
// make sure the caller match the provider
func (en *Engine) tryAddProvider(key string,own common.Address,price *big.Int) error {
	p := en.getProviderByAddr(own)
	if p == nil {
		service := make(map[string]*Entry)
		service[key] = newEntry(key,price)
		en.addProvider(&provider{
			Addr: own,
			Service: service,
			DealList:	make([]*deal,0,0),
		})
		return nil
	} 
	return p.addEntry(key,price)
}
// check provider's address by caller in contract,make sure the caller match the provider
func (en *Engine) batchGetPkFromDeals(keys []string,addrs []common.Address,own common.Address) ([][]byte,error) {
	l := len(keys)
	if len(keys) != len(addrs) {
		return nil, ErrInvalidParams
	}
	p := en.getProviderByAddr(own)
	if p == nil {
		return nil,ErrNotFoundProvider
	}
	res := make([][]byte,0,0)
	for i:=0;i<l;i++ {
		k,a := keys[i],addrs[i]
		d := p.getDealResult(k,a)
		if d != nil {
			pk := d.getPk()
			if pk != nil {
				res = append(res,pk)
				continue
			}
		} 
		res = append(res,[]byte{0})		// invalid pk
	}
	return res,nil
}
func (en *Engine) batchSetPassword(keys []string,addrs []common.Address,ecPass [][]byte,own common.Address) (int,error) {
	l := len(keys)
	if len(keys) != len(addrs) || l != len(ecPass) {
		return 0, ErrInvalidParams
	}
	p := en.getProviderByAddr(own)
	if p == nil {
		return 0,ErrNotFoundProvider
	}
	for i:=0;i<l;i++ {
		k,a,ec := keys[i],addrs[i],ecPass[i]
		if err := p.setPassword(k,a,ec); err != nil {
			return i,err
		}
	}
	return 0,nil
}

func (en *Engine) GetPubKeyByProvider(key string, addr,own common.Address) ([]byte,error) {
	p := en.getProviderByAddr(own)
	if p == nil {
		return nil,ErrNotFoundProvider
	}
	d := p.getDealResult(key,addr)
	if d == nil {
		return nil,ErrNotFoundDeal
	}
	pk := d.getPk()
	if pk == nil {
		return nil,ErrInvalidPk
	}
	return pk,nil
}
func (en *Engine) SetPasswordByProvider(key string, ecPass []byte,addr,own common.Address) error {
	p := en.getProviderByAddr(own) 
	if p == nil {
		return ErrNotFoundProvider
	}
	
	return p.setPassword(key,addr,ecPass)
}
// the balance will auto translate to the consumer for no deal when more than MaxAutoRedeemHeight height 
func (en *Engine) ActionRedeemForConsumer(ch uint64) error {
	// check the not finished deal
	for _,p := range en.maker {
		for _,d := range p.DealList {
			if d.canRedeem(ch) {
				d.doRedeem()
				d.redeemed()
			}
		}
	}
	return nil
}
// 
func (en *Engine) ActionUnlockedForProvider(ch uint64) error {
	for _,p := range en.maker {
		for _,d := range p.DealList {
			if d.isPayed() && int64(ch - d.getHeight()) > int64(MaxAutoUnlockedHieght) {
				transToProvider(d.getPrice(),p.getAddress())
				d.finish()
			}
		}
	}
	return nil
}

///////////////API///////////////////////////////////////////////////////////////////////////
///////////////API///////////////////////////////////////////////////////////////////////////

