package eth

import (
	"fmt"
	"github.com/piupuer/go-helper/pkg/utils"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

type Eth struct {
	decimals int
	val      string
	wei      *big.Int
}

// https://github.com/ethereum/go-ethereum/issues/21221

func New(decimals int) Eth {
	return Eth{
		decimals: decimals,
		val:      "0",
	}
}

// Gt >
func (et Eth) Gt(s2 Eth) bool {
	// convert to wei
	r := et.Amount2Wei().wei.Cmp(s2.Amount2Wei().wei)
	if r == 1 {
		return true
	}
	return false
}

// Gte >=
func (et Eth) Gte(s2 Eth) bool {
	r := et.Amount2Wei().wei.Cmp(s2.Amount2Wei().wei)
	if r >= 0 {
		return true
	}
	return false
}

// Add s + s2
func (et Eth) Add(s2 Eth) Eth {
	return operation("add", et, s2)
}

// Sub s - s2
func (et Eth) Sub(s2 Eth) Eth {
	return operation("sub", et, s2)
}

// Mul s * s2
func (et Eth) Mul(s2 Eth) Eth {
	return operation("mul", et, s2)
}

// Div s / s2
func (et Eth) Div(s2 Eth) Eth {
	return operation("div", et, s2)
}

func operation(operator string, s1, s2 Eth) Eth {
	var w3 *big.Int
	switch operator {
	case "add":
		w3 = new(big.Int).Add(s1.Amount2Wei().wei, s2.Amount2Wei().wei)
	case "sub":
		w3 = new(big.Int).Sub(s1.Amount2Wei().wei, s2.Amount2Wei().wei)
	case "mul":
		w3 = new(big.Int).Mul(s1.Amount2Wei().wei, s2.Amount2Wei().wei)
	case "div":
		w3 = new(big.Int).Div(s1.Amount2Wei().wei, s2.Amount2Wei().wei)
	}
	// new eth
	s3 := Eth{
		decimals: s1.decimals,
		val:      w3.String(),
	}
	// convert to amount
	return s3.Wei2Amount()
}

// Set amount value
func (et Eth) Set(amount string) Eth {
	et.val = fmt.Sprintf("%d", utils.Str2Int64(amount))
	return et
}

// SetWei set wei value
func (et Eth) SetWei(wei string) Eth {
	et.val = fmt.Sprintf("%d", utils.Str2Int64(wei))
	return et.Wei2Amount()
}

func (et Eth) Amount2Wei() Eth {
	if et.decimals > 100 || et.decimals < 0 {
		et.wei = new(big.Int).SetInt64(0)
		et.val = et.wei.String()
		return et
	}
	if et.decimals == 0 {
		et.wei = new(big.Int).SetInt64(utils.Str2Int64(et.val))
		et.val = et.wei.String()
		return et
	}
	f, _ := et.parseBigFloat()
	truncInt, _ := f.Int(nil)

	d := new(big.Float)
	d.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	d.SetMode(big.ToNearestEven)
	// 10^decimals
	d.SetString(fmt.Sprintf("1e%d", et.decimals))
	dInt, _ := d.Int(nil)
	// value * 10^decimals
	truncInt = new(big.Int).Mul(truncInt, dInt)
	fracStr := strings.Split(fmt.Sprintf(fmt.Sprintf("%%.%df", et.decimals), f), ".")[1]
	fracStr += strings.Repeat("0", et.decimals-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	et.wei = wei
	et.val = et.wei.String()
	return et
}

func (et Eth) Wei2Amount() Eth {
	if et.decimals > 100 || et.decimals < 0 {
		et.val = "0"
		return et
	}
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei, _ := et.parseBigFloat()
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	y := new(big.Float)
	y.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	y.SetMode(big.ToNearestEven)
	y.SetString(fmt.Sprintf("1e%d", et.decimals))
	r := f.Quo(fWei, y)
	decimalNum, err := decimal.NewFromString(fmt.Sprintf("%v", r))
	if err != nil {
		et.val = "0"
		return et
	}
	et.val = decimalNum.String()
	return et
}

func (et Eth) BigInt() *big.Int {
	i, _ := new(big.Int).SetString(et.val, 10)
	return i
}

func (et Eth) String() string {
	return et.val
}

// parseBigFloat parse string value to big.Float
func (et Eth) parseBigFloat() (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(et.val, f)
	return f, err
}
