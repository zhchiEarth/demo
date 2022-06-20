package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AccountCToken struct {
	ID int `json:"id,omitempty"`
	// 用户地址
	User string `json:"user,omitempty"`
	// ctoken地址
	Address string `json:"address,omitempty"`
	// ctoken symbol
	Symbol string `json:"symbol,omitempty"`
	// 进入市场的标志 true 进入市场，允许借贷
	EnteredMarket bool `json:"entered_market,omitempty"`
	// ctoken 余额
	CtokenBalance string `json:"ctoken_balance,omitempty"`
	// 借款的余额
	StoredBorrowBalance string `json:"stored_borrow_balance,omitempty"`
	// 用户上一次的贷款指数
	BorrowIndex string `json:"borrow_index,omitempty"`
	// 总的token 供应量
	TotalUnderlyingSupplied string `json:"total_underlying_supplied,omitempty"`
	// 累计提现
	TotalUnderlyingRedeemed string `json:"total_underlying_redeemed,omitempty"`
	// 累计借款
	TotalUnderlyingBorrowed string `json:"total_underlying_borrowed,omitempty"`
	// 累计还款
	TotalUnderlyingRepaid string `json:"total_underlying_repaid,omitempty"`
	// 转换为基础代币的 cTokenBalance * market.exchangeRate
	SupplyBalanceUnderlying string `json:"supply_balance_underlying,omitempty"`
	// token的借款 borrowBalanceUnderlying = storedBorrowBalance * market.borrowIndex / accountBorrowIndex
	BorrowBalanceUnderlying string `json:"borrow_balance_underlying,omitempty"`
	// 累积的供应利息 lifetimeSupplyInterestAccrued = supplyBalanceUnderlying - totalUnderlyingSupplied + totalUnderlyingRedeemed
	LifetimeSupplyInterestAccrued string `json:"lifetime_supply_interest_accrued,omitempty"`
	// 生命周期内应计的借款利息金额 lifetimeSupplyInterestAccrued = supplyBalanceUnderlying - totalUnderlyingSupplied + totalUnderlyingRedeemed
	LifetimeBorrowInterestAccrued string `json:"lifetime_borrow_interest_accrued,omitempty"`
	// 可以提取的供应量，以使用户的健康保持在 1.25 或更高
	SafeWithdrawAmountUnderlying string `json:"safe_withdraw_amount_underlying,omitempty"`
	// 抵押物价值，usd计价
	CollateralValueInUsd string `json:"collateral_value_in_usd,omitempty"`
	// 贷款价值，usd计价
	BorrowValueInUsd string `json:"borrow_value_in_usd,omitempty"`
	// usd价格
	UnderlyingPriceUsd string `json:"underlying_price_usd,omitempty"`
	// 区块高度
	BlockNumber uint64 `json:"block_number,omitempty"`
}

type AccountCTokenReply struct {
	ID int
}

type AccountCTokenRepo interface {
	SaveCToken(ctx context.Context, aid int, ac *AccountCToken) (*AccountCTokenReply, error)
	//GetUsersOraclePriceChanged(ctx context.Context, cToken AccountCToken,  pageNum, pageSize int64) ([]string, error)
}

type AccountCTokenUseCase struct {
	repo AccountCTokenRepo
	log  *log.Helper
}

func NewAccountCTokenUseCase(repo AccountCTokenRepo, logger log.Logger) *AccountCTokenUseCase {
	return &AccountCTokenUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/accountctoken")),
	}
}

func (uc *AccountCTokenUseCase) Save(ctx context.Context, accountId int, u *AccountCToken) (*AccountCTokenReply, error) {
	return uc.repo.SaveCToken(ctx, accountId, u)
}

//func (uc *AccountCTokenUseCase) GetUsersOraclePriceChanged(ctx context.Context, cToken AccountCToken,  pageNum,
//	pageSize int64) ([]string, error) {
//	return uc.repo.GetUsersOraclePriceChanged(ctx, cToken, pageNum, pageSize)
//}
