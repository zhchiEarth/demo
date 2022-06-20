package util

import "github.com/ethereum/go-ethereum/common"

func AddressToHex(addr string) string {
	return common.HexToAddress(addr).Hex()
}

func AddressesToHex(list []string) []string {
	var l []string
	for _, addr := range list {
		l = append(l, common.HexToAddress(addr).Hex())
	}
	return l
}
