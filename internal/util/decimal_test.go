package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestClient(t *testing.T) {

	addr := common.HexToAddress("0x3b2bd810Ade98f19A29C05bC9350A1FD01694CF2")
	fmt.Println(addr.Hex())
	fmt.Println(addr.Hash().TerminalString())
	fmt.Println(addr.String())
	fmt.Println(addr.Value())
}
