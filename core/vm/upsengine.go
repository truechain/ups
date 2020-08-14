package vm

import (
	"fmt"
	"math/big"
	"crypto/rand"
	"strings"
	"io"
	"bytes"
	"sort"
	"errors"
	"github.com/truechain/ups/common"
	"github.com/truechain/ups/accounts/abi"
	// "github.com/truechain/ups/core/state"
	// "github.com/truechain/ups/core/types"
	"github.com/truechain/ups/crypto"
	"github.com/truechain/ups/crypto/ecies"
	"github.com/truechain/ups/log"
	// "github.com/truechain/ups/params"
	"github.com/truechain/ups/rlp"
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

var abiEngine abi.ABI

func init() {
	abiEngine, _ = abi.JSON(strings.NewReader(ABIENGINEJSON))
}

///////////////API///////////////////////////////////////////////////////////////////////////
// addProvider
func addProvider(evm *EVM, contract *Contract, input []byte) (ret []byte, err error) {
	args := struct {
		Key	 []byte
		Price  *big.Int
	}{}
	method, _ := abiStaking.Methods["addProvider"]

	err = method.Inputs.Unpack(&args, input)
	if err != nil {
		log.Error("Unpack addProvider error", "err", err)
		return nil, ErrStakingInvalidInput
	}

	from := contract.caller.Address()
	engine := NewEngine()
	err = engine.Load(evm.StateDB)
	if err != nil {
		log.Error("engine load error", "error", err)
		return nil, err
	}
	err = engine.tryAddProvider(string(args.Key), from, args.Price)
	if err != nil {
		log.Error("tryAddProvider", "address", contract.caller.Address(),"Key",args.Key, "Price", args.Price, "error", err)
		return nil, err
	}

	err = engine.Save(evm.StateDB)
	if err != nil {
		log.Error("engine save state error", "error", err)
		return nil, err
	}

	event := abiStaking.Events["AddProvider"]
	logData, err := event.Inputs.PackNonIndexed(args.Key, args.Price)
	if err != nil {
		log.Error("Pack AddProvider log error", "error", err)
		return nil, err
	}
	topics := []common.Hash{
		event.ID(),
		common.BytesToHash(from[:]),
	}
	logN(evm, contract, topics, logData)
	return nil, nil
}
// postRequestKey
func postRequestKey(evm *EVM, contract *Contract, input []byte) (ret []byte, err error) {
	args := struct {
		Key	 []byte
		Pk   []byte
	}{}
	method, _ := abiStaking.Methods["postRequestKey"]

	err = method.Inputs.Unpack(&args, input)
	if err != nil {
		log.Error("Unpack postRequestKey error", "err", err)
		return nil, ErrStakingInvalidInput
	}

	from := contract.caller.Address()
	engine := NewEngine()
	err = engine.Load(evm.StateDB)
	if err != nil {
		log.Error("engine load error", "error", err)
		return nil, err
	}
	err = engine.tryAddConsumer(newConsumer(string(args.Key),contract.Value(),args.Pk), evm.Context.BlockNumber.Uint64())
	if err != nil {
		log.Error("postRequestKey", "address", from,"Key",string(args.Key), "pk", args.Pk, "error", err)
		return nil, err
	}

	err = engine.Save(evm.StateDB)
	if err != nil {
		log.Error("engine save state error", "error", err)
		return nil, err
	}

	event := abiStaking.Events["PostRequestKey"]
	logData, err := event.Inputs.PackNonIndexed(args.Key, args.Pk)
	if err != nil {
		log.Error("Pack PostRequestKey log error", "error", err)
		return nil, err
	}
	topics := []common.Hash{
		event.ID(),
		common.BytesToHash(from[:]),
	}
	logN(evm, contract, topics, logData)
	return nil, nil
}
// getPubKeyFromDeal
func getPubKeyFromDeal(evm *EVM, contract *Contract, input []byte) (ret []byte, err error) {
	args := struct {
		Key	 	 []byte
		Addr 	 common.Address
	}{}
	method, _ := abiStaking.Methods["getPubKeyFromDeal"]

	err = method.Inputs.Unpack(&args, input)
	if err != nil {
		log.Error("Unpack getPubKeyFromDeal error", "err", err)
		return nil, ErrStakingInvalidInput
	}

	from := contract.caller.Address()
	engine := NewEngine()
	err = engine.Load(evm.StateDB)
	if err != nil {
		log.Error("engine load error", "error", err)
		return nil, err
	}
	if pk,err := engine.GetPubKeyByProvider(string(args.Key),args.Addr,from); err != nil {
		log.Error("getPubKeyFromDeal", "address", from,"Key",string(args.Key), "consumerAddr", args.Addr, "error", err)
		return nil, err
	} else {
		return method.Outputs.Pack(pk)
	}
}
// SetFileKey
func SetFileKey(evm *EVM, contract *Contract, input []byte) (ret []byte, err error) {
	args := struct {
		Key	 	 []byte
		EcPass   []byte
		Addr 	 common.Address
	}{}
	method, _ := abiStaking.Methods["setFileKey"]

	err = method.Inputs.Unpack(&args, input)
	if err != nil {
		log.Error("Unpack setFileKey error", "err", err)
		return nil, ErrStakingInvalidInput
	}

	from := contract.caller.Address()
	engine := NewEngine()
	err = engine.Load(evm.StateDB)
	if err != nil {
		log.Error("engine load error", "error", err)
		return nil, err
	}
	err = engine.SetPasswordByProvider(string(args.Key),args.EcPass,args.Addr,from)
	if err != nil {
		log.Error("setFileKey", "address", from,"Key",string(args.Key), "EcPass", args.EcPass, "error", err)
		return nil, err
	}

	err = engine.Save(evm.StateDB)
	if err != nil {
		log.Error("engine save state error", "error", err)
		return nil, err
	}

	event := abiStaking.Events["SetFileKey"]
	logData, err := event.Inputs.PackNonIndexed(args.Key, args.EcPass,args.Addr)
	if err != nil {
		log.Error("Pack SetFileKey log error", "error", err)
		return nil, err
	}
	topics := []common.Hash{
		event.ID(),
		common.BytesToHash(from[:]),
	}
	logN(evm, contract, topics, logData)
	return nil, nil
}
// GetFileKey
func GetFileKey(evm *EVM, contract *Contract, input []byte) (ret []byte, err error) {	
	args := struct {
		Key	 	 []byte
		Addr 	 common.Address
	}{}
	method, _ := abiStaking.Methods["getFileKey"]

	err = method.Inputs.Unpack(&args, input)
	if err != nil {
		log.Error("Unpack setFileKey error", "err", err)
		return nil, ErrStakingInvalidInput
	}

	from := contract.caller.Address()
	engine := NewEngine()
	err = engine.Load(evm.StateDB)
	if err != nil {
		log.Error("engine load error", "error", err)
		return nil, err
	}
	if ec,err := engine.GetPasswordByConsumer(string(args.Key),args.Addr,from); err != nil {
		log.Error("setFileKey", "address", from,"Key",string(args.Key), "providerAddress", args.Addr, "error", err)
		return nil, err
	} else {
		return method.Outputs.Pack(ec)
	}
}
///////////////API///////////////////////////////////////////////////////////////////////////

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
func (d *deal) getPassword() []byte {
	ps := make([]byte,len(d.password))
	copy(ps,d.password)
	return ps
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
func (d *deal) canClear() bool {
	if d.state == DealRedeem || d.state == DealFinish {
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
func (p *provider) clearDealList() {
	tmp := make([]*deal,0,0)
	for _,d := range p.DealList {
		if !d.canClear() {
			tmp = append(tmp,d)
		}
	}
	p.DealList = tmp
}

type consumer struct {
	Key 	string
	Price   *big.Int
	Pk 		[]byte	
}
func newConsumer(key string,price *big.Int,pk []byte) *consumer {
	return &consumer{
		Key:	key,
		Price:	new(big.Int).Set(price),
		Pk:		pk,
	}
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
////////////////////////////////////////////////////////////////////////////////////////////////
type StoreEngine struct {
	Address common.Address
	Pov      *provider
}

type SortStoreEngine []*StoreEngine

func (vs SortStoreEngine) Len() int {
	return len(vs)
}
func (vs SortStoreEngine) Less(i, j int) bool {
	return bytes.Compare(vs[i].Address[:], vs[j].Address[:]) == -1
}
func (vs SortStoreEngine) Swap(i, j int) {
	it := vs[i]
	vs[i] = vs[j]
	vs[j] = it
}

type Engine struct {
	maker  map[common.Address]*provider
}
func NewEngine() *Engine {
	return &Engine{
		maker: 	make(map[common.Address]*provider),
	}
}

func (en *Engine) Load(state StateDB) error {
	key := common.BytesToHash(UpsEngineAddress[:])
	data := state.GetPOSState(UpsEngineAddress, key)
	lenght := len(data)
	if lenght == 0 {
		return errors.New("Load data = 0")
	}
	var temp Engine
	if err := rlp.DecodeBytes(data, &temp); err != nil {
		log.Error("Invalid Engine entry RLP", "err", err)
		return errors.New(fmt.Sprintf("Invalid Engine entry RLP %s", err.Error()))
	}
	en.maker = temp.maker
	return nil
}
func (en *Engine) Save(state StateDB) error {
	key := common.BytesToHash(UpsEngineAddress[:])
	data, err := rlp.EncodeToBytes(en)
	if err != nil {
		log.Crit("Failed to RLP encode Engine", "err", err)
	}
	state.SetPOSState(UpsEngineAddress, key, data)
	return err
}
/////////////////////////////////////////////////////////////////
func (en *Engine) toSlice() SortStoreEngine {
	v1 := make([]*StoreEngine, 0, 0)
	for k, v := range en.maker {
		v1 = append(v1, &StoreEngine{
			Address: 	k,
			Pov:      	v,
		})
	}
	sort.Sort(SortStoreEngine(v1))
	return SortStoreEngine(v1)
}
func (en *Engine) fromSlice(v1 SortStoreEngine) {
	enInfos := make(map[common.Address]*provider)
	for _, v := range v1 {
		enInfos[v.Address] = v.Pov
	}
	en.maker = enInfos
}
func (en *Engine) DecodeRLP(s *rlp.Stream) error {
	eb := struct {
		Value1 SortStoreEngine
	}{}
	if err := s.Decode(&eb); err != nil {
		return err
	}
	en.fromSlice(eb.Value1)
	return nil
}
func (en *Engine) EncodeRLP(w io.Writer) error {
	tmp := struct {
		Value1 SortStoreEngine
	}{ en.toSlice() }

	return rlp.Encode(w, &tmp)
}

/////////////////////////////////////////////////////////////////

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
func (en *Engine) GetPasswordByConsumer(key string, addr, own common.Address) ([]byte,error) {
	p := en.getProviderByAddr(addr)
	if p == nil {
		return nil,ErrNotFoundProvider
	}
	d := p.getDealResult(key,own)
	if d == nil {
		return nil,ErrNotFoundDeal
	}
	return d.getPassword(),nil
}
// the balance will auto translate to the consumer for no deal when more than MaxAutoRedeemHeight height 
func (en *Engine) actionRedeemForConsumer(ch uint64) error {
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
func (en *Engine) actionUnlockedForProvider(ch uint64) error {
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
func (en *Engine) actionClearDeal(ch uint64) error {
	if ch % 200 == 0 {
		for _,p := range en.maker {
			p.clearDealList()
		}
	}
	return nil 
}
func (en *Engine) DoAction(ch uint64) error {
	err := en.actionRedeemForConsumer(ch)
	if err != nil {
		return err
	}
	err = en.actionUnlockedForProvider(ch)
	if err != nil {
		return err
	}
	return en.actionClearDeal(ch)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
const ABIENGINEJSON = `

`

