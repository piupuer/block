package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestEncodeAbiParams(t *testing.T) {
	s, err := EncodeAbiParams(
		[]string{"address", "address", "address"},
		common.HexToAddress("0xf76D4a441E4ba86A923ce32B89AFF89dBccAA075"),
		common.HexToAddress("0xcfa7b0e37f5AC60f3ae25226F5e39ec59AD26152"),
		common.HexToAddress("0x07865c6e87b9f70255377e024ace6630c1eaa37f"),
	)
	fmt.Println(s, err)
}

func TestEncodeAbiParams2(t *testing.T) {
	s, err := EncodeAbiParams(
		[]string{"address", "address[]"},
		common.HexToAddress("0xC6A4E105c6ebe68ae2B50b5b53b1a0498381b132"),
		[]common.Address{
			common.HexToAddress("0x2973e69b20563bcc66dC63Bde153072c33eF37fe"),
		},
	)
	fmt.Println(s, err)
}

func TestEncodeAbiFun(t *testing.T) {
	s, err := EncodeAbiFun(
		`[{"constant":false,"inputs":[{"internalType":"contract Comp","name":"comp","type":"address"},{"internalType":"contract ComptrollerLensInterface","name":"comptroller","type":"address"},{"internalType":"address","name":"account","type":"address"}],"name":"getCompBalanceMetadataExt","outputs":[{"components":[{"internalType":"uint256","name":"balance","type":"uint256"},{"internalType":"uint256","name":"votes","type":"uint256"},{"internalType":"address","name":"delegate","type":"address"},{"internalType":"uint256","name":"allocated","type":"uint256"}],"internalType":"struct CompoundLens.CompBalanceMetadataExt","name":"","type":"tuple"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`,
		"getCompBalanceMetadataExt",
		common.HexToAddress("0xf76D4a441E4ba86A923ce32B89AFF89dBccAA075"),
		common.HexToAddress("0xcfa7b0e37f5AC60f3ae25226F5e39ec59AD26152"),
		common.HexToAddress("0xC6A4E105c6ebe68ae2B50b5b53b1a0498381b132"),
	)
	fmt.Println(s, err)
}

func TestEncodeAbiFun2(t *testing.T) {
	s, err := EncodeAbiFun(
		`[{"constant":false,"inputs":[{"internalType":"address","name":"holder","type":"address"},{"internalType":"contract CToken[]","name":"cTokens","type":"address[]"}],"name":"claimComp","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`,
		"claimComp",
		common.HexToAddress("0xC6A4E105c6ebe68ae2B50b5b53b1a0498381b132"),
		[]common.Address{
			common.HexToAddress("0x2973e69b20563bcc66dC63Bde153072c33eF37fe"),
		},
	)
	fmt.Println(s, err)
}

func TestDecodeAbiFun(t *testing.T) {
	s, err := DecodeAbiFun(
		`[{"constant":false,"inputs":[{"internalType":"contract Comp","name":"comp","type":"address"},{"internalType":"contract ComptrollerLensInterface","name":"comptroller","type":"address"},{"internalType":"address","name":"account","type":"address"}],"name":"getCompBalanceMetadataExt","outputs":[{"components":[{"internalType":"uint256","name":"balance","type":"uint256"},{"internalType":"uint256","name":"votes","type":"uint256"},{"internalType":"address","name":"delegate","type":"address"},{"internalType":"uint256","name":"allocated","type":"uint256"}],"internalType":"struct CompoundLens.CompBalanceMetadataExt","name":"","type":"tuple"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`,
		"getCompBalanceMetadataExt",
		"000000000000000000000000000000000000000000000000032dcba5230fa4300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010c24e5b4bd21472",
	)
	fmt.Println(s, err)
}

func TestDecodeAbiParams(t *testing.T) {
	s, err := DecodeAbiParams(
		[]string{"address", "address[]"},
		"000000000000000000000000c6a4e105c6ebe68ae2b50b5b53b1a0498381b132000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000010000000000000000000000002973e69b20563bcc66dc63bde153072c33ef37fe",
	)
	fmt.Println(s, err)
}
