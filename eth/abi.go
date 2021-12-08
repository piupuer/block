package eth

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/piupuer/go-helper/pkg/utils"
	"strings"
)

func EncodeAbiFun(abiJson, name string, args ...interface{}) (res string, err error) {
	ssAbi, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return
	}
	var data []byte
	data, err = ssAbi.Pack(name, args...)
	if err != nil {
		return
	}
	return hex.EncodeToString(data), nil
}

func EncodeAbiParams(argTypes []string, args ...interface{}) (res string, err error) {
	arguments := make(abi.Arguments, len(argTypes))
	for i, item := range argTypes {
		var t abi.Type
		t, err = abi.NewType(item, "", nil)
		if err != nil {
			return
		}
		arg := abi.Argument{
			Type: t,
		}
		arguments[i] = arg
	}
	var data []byte
	data, err = arguments.Pack(args...)
	if err != nil {
		return
	}
	return hex.EncodeToString(data), nil
}

func DecodeAbiParams(argTypes []string, str string) (res []interface{}, err error) {
	data, err := hex.DecodeString(str)
	if err != nil {
		return
	}
	arguments := make(abi.Arguments, len(argTypes))
	for i, item := range argTypes {
		var t abi.Type
		t, err = abi.NewType(item, "", nil)
		if err != nil {
			return
		}
		arg := abi.Argument{
			Type: t,
		}
		arguments[i] = arg
	}
	res, err = arguments.Unpack(data)
	return
}

func DecodeAbiFun(abiJson, name, encode string) (res string, err error) {
	ssAbi, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return
	}
	var data []interface{}
	data, err = ssAbi.Unpack(name, common.Hex2Bytes(encode))
	if err == nil && len(data) > 0 {
		res = utils.Struct2Json(data[0])
	}
	return
}
