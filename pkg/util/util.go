package util

import (
	"math/big"
	"strconv"
)

func StringToBigFloat(value string) (*big.Float, error) {
	if len(value) == 0 {
		return big.NewFloat(0), nil
	}
	underlyingPriceUSD, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}
	return big.NewFloat(underlyingPriceUSD), nil
}
