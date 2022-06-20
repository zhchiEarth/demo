package util

import (
	"github.com/shopspring/decimal"
)

type ErrDecimal struct {
	err error
}

func (ed *ErrDecimal) NewFrom(value interface{}) decimal.Decimal {
	var (
		val decimal.Decimal
		err error
	)

	switch value.(type) {
	case string:
		val, err = decimal.NewFromString(value.(string))
		break
	}
	if err != nil {
		ed.err = err
		return decimal.Zero
	}
	return val
}

// 通过cTokenBalance计算 Token的金额
// CalcSupplyBalanceUnderlying SupplyBalanceUnderlying = cTokenBalance * exchangeRate
func CalcSupplyBalanceUnderlying(cTokenBalance string, exchangeRate string) (decimal.Decimal, error) {
	cTokenBalanceBD, err := decimal.NewFromString(cTokenBalance)
	if err != nil {
		return decimal.Zero, err
	}
	exchangeRateBD, err := decimal.NewFromString(exchangeRate)
	if err != nil {
		return decimal.Zero, err
	}
	if cTokenBalanceBD.Equal(decimal.Zero) {
		return decimal.Zero, nil
	}

	return cTokenBalanceBD.Mul(exchangeRateBD), nil
}

// 计算抵押物可以抵押多少 underlying token
func CalcCollateralValue(underlyingBalance string, collateralFactor string) (decimal.Decimal, error) {
	collateralFactorBD, err := decimal.NewFromString(collateralFactor)
	if err != nil {
		return decimal.Zero, err
	}
	underlyingBalanceBD, err := decimal.NewFromString(underlyingBalance)
	if err != nil {
		return decimal.Zero, err
	}
	return underlyingBalanceBD.Mul(collateralFactorBD), nil
}

// CalcCollateralValueInUsd 计算抵押物可以抵押多少USD
func CalcCollateralValueInUsd(underlyingBalance string, collateralFactor string, underlyingPriceUsd string) (decimal.Decimal, error) {
	collateralValue, err := CalcCollateralValue(underlyingBalance, collateralFactor)
	if err != nil {
		return decimal.Zero, err
	}
	underlyingPriceUsdBD, err := decimal.NewFromString(underlyingPriceUsd)
	if err != nil {
		return decimal.Zero, err
	}
	return collateralValue.Mul(underlyingPriceUsdBD), nil
}

func CalcBorrowValueInUsd(underlyingBalance string, collateralFactor string) (decimal.Decimal, error) {
	collateralFactorBD, err := decimal.NewFromString(collateralFactor)
	if err != nil {
		return decimal.Zero, err
	}
	underlyingBalanceBD, err := decimal.NewFromString(underlyingBalance)
	if err != nil {
		return decimal.Zero, err
	}
	return underlyingBalanceBD.Mul(collateralFactorBD), nil
}

// CalcBorrowBalanceUnderlying  借款
//borrowBalanceUnderlying = storedBorrowBalance * market.borrowIndex / accountBorrowIndex
func CalcBorrowBalanceUnderlying(storedBorrowBalance string, borrowIndex string, accountBorrowIndex string) (decimal.Decimal, error) {
	storedBorrowBalanceBD, err := decimal.NewFromString(storedBorrowBalance)
	if storedBorrowBalanceBD.Equal(decimal.Zero) {
		return decimal.Zero, nil
	}
	if err != nil {
		return decimal.Zero, err
	}
	borrowIndexBD, err := decimal.NewFromString(borrowIndex)
	if err != nil {
		return decimal.Zero, err
	}
	accountBorrowIndexBD, err := decimal.NewFromString(accountBorrowIndex)
	if err != nil {
		return decimal.Zero, err
	}
	return storedBorrowBalanceBD.Mul(borrowIndexBD).Div(accountBorrowIndexBD), nil
}

// RemoveDuplicateElement 切片去重
func RemoveDuplicateElement(list []string) []string {
	result := make([]string, 0, len(list))
	temp := map[string]struct{}{}
	for _, item := range list {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
