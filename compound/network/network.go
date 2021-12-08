package network

import (
	"embed"
	"fmt"
	"github.com/piupuer/go-helper/pkg/utils"
)

//go:embed conf
var networks embed.FS

const (
	ContractKey = "Contracts"
	BlockKey    = "Blocks"
	TokenKey    = "Tokens"
	CTokenKey   = "cTokens"
)

type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Address  string `json:"address"`
}

type Network struct {
	n string
	m map[string]interface{}
}

func Select(name string) *Network {
	return &Network{
		n: name,
		m: read2Map(name),
	}
}

func (nw Network) Address(contract string) (res string) {
	item := nw.m[ContractKey]
	m := make(map[string]string)
	utils.Struct2StructByJson(item, &m)
	res = m[contract]
	return
}

func (nw Network) Block(contract string) (res int64) {
	item := nw.m[BlockKey]
	m := make(map[string]int64)
	utils.Struct2StructByJson(item, &m)
	res = m[contract]
	return
}

func (nw Network) Token(symbol string) (res Token) {
	item := nw.m[TokenKey]
	m := make(map[string]Token)
	utils.Struct2StructByJson(item, &m)
	res = m[symbol]
	return
}

func (nw Network) CToken(cSymbol string) (res Token) {
	item := nw.m[CTokenKey]
	m := make(map[string]Token)
	utils.Struct2StructByJson(item, &m)
	res = m[cSymbol]
	return
}

func (nw Network) Abi() *Abi {
	return &Abi{
		network: nw,
		m:       read2Map(nw.n + "-abi"),
	}
}

type Abi struct {
	network Network
	m       map[string]interface{}
}

func (ab Abi) Json(contract string) (res string) {
	item := ab.m[contract]
	return utils.Struct2Json(item)
}

func read2Map(name string) map[string]interface{} {
	bs, _ := networks.ReadFile(fmt.Sprintf("conf/%s.json", name))
	m := make(map[string]interface{})
	utils.Json2Struct(string(bs), &m)
	return m
}
