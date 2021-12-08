package network

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	fmt.Println(Select("ropsten").Address("CompoundLens"))
	fmt.Println(Select("ropsten").Block("CompoundLens"))
	fmt.Println(Select("ropsten").Abi().Json("CompoundLens"))
}
