package eth

import (
	"fmt"
	"testing"
)

func TestEth_String(t *testing.T) {
	fmt.Println(New(8).Set("1").String())
	fmt.Println(New(18).Set("1").String())
	fmt.Println(New(50).Set("1").String())
	fmt.Println(New(120).Set("1").String())

	fmt.Println(New(8).Set("1").Amount2Wei().String())
	fmt.Println(New(18).Set("1").Amount2Wei().String())
	fmt.Println(New(50).Set("1").Amount2Wei().String())
	fmt.Println(New(120).Set("1").Amount2Wei().String())

	fmt.Println(New(8).SetWei("1").String())
	fmt.Println(New(18).SetWei("1").String())
	fmt.Println(New(50).SetWei("1").String())
	fmt.Println(New(120).SetWei("1").String())

	fmt.Println(New(8).SetWei("1").Amount2Wei().String())
	fmt.Println(New(18).SetWei("1").Amount2Wei().String())
	fmt.Println(New(50).SetWei("1").Amount2Wei().String())
	fmt.Println(New(120).SetWei("1").Amount2Wei().String())

	fmt.Println(New(8).Set("100").Gt(New(8).Set("200")))
	fmt.Println(New(8).Set("200").Gt(New(8).Set("200")))
	fmt.Println(New(8).Set("200").Gte(New(8).Set("200")))

	fmt.Println(New(8).Set("1").Add(New(8).Set("2")))
	fmt.Println(New(8).Set("1000000000").Add(New(8).Set("2000000000")))
	fmt.Println(New(8).Set("1").Add(New(8).Set("-2")))

	fmt.Println(New(8).Set("3").Sub(New(8).Set("1")))
	fmt.Println(New(8).Set("3000000000").Sub(New(8).Set("1000000000")))
	fmt.Println(New(8).Set("3").Sub(New(8).Set("-1")))
}
