package data

import (
	"compound/internal/biz"
	"compound/internal/data/ent"
	"compound/internal/data/ent/accountctoken"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var _ biz.AccountCTokenRepo = (*accountCTokenRepo)(nil)

type accountCTokenRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountCTokenRepo(data *Data, logger log.Logger) biz.AccountCTokenRepo {
	return &accountCTokenRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/accountctoken")),
	}
}

func (r *accountCTokenRepo) SaveCToken(ctx context.Context, accountId int, b *biz.AccountCToken) (*biz.AccountCTokenReply, error) {
	entity, err := r.data.AccountCToken(ctx).Query().
		Where(
			accountctoken.Address(b.Address),
			accountctoken.User(b.User)).
		First(ctx)
	if !ent.IsNotFound(err) && err != nil {
		return nil, errors.Wrap(err, "")
	}
	// 区块高度 大于数据库的区块高度
	if entity == nil {
		return r.CreateAccountCToken(ctx, accountId, b)
	}
	//if b.BlockNumber > entity.BlockNumber {
	return r.UpdateOneAccountCToken(ctx, entity.ID, b)
	//}
	//return &biz.AccountCTokenReply{
	//	ID: entity.ID,
	//}, nil
}

func (r *accountCTokenRepo) CreateAccountCToken(ctx context.Context, accountId int, b *biz.AccountCToken) (*biz.AccountCTokenReply, error) {
	m, err := r.data.AccountCToken(ctx).Create().
		SetUser(b.User).
		SetAccountID(accountId).
		SetAddress(b.Address).
		SetSymbol(b.Symbol).
		SetEnteredMarket(b.EnteredMarket).
		SetCtokenBalance(b.CtokenBalance).
		SetStoredBorrowBalance(b.StoredBorrowBalance).
		SetBorrowIndex(b.BorrowIndex).
		SetTotalUnderlyingSupplied(b.TotalUnderlyingSupplied).
		SetTotalUnderlyingRedeemed(b.TotalUnderlyingRedeemed).
		SetTotalUnderlyingBorrowed(b.TotalUnderlyingBorrowed).
		SetTotalUnderlyingRepaid(b.TotalUnderlyingRepaid).
		SetSupplyBalanceUnderlying(b.SupplyBalanceUnderlying).
		SetBorrowBalanceUnderlying(b.BorrowBalanceUnderlying).
		SetLifetimeSupplyInterestAccrued(b.LifetimeSupplyInterestAccrued).
		SetLifetimeBorrowInterestAccrued(b.LifetimeBorrowInterestAccrued).
		SetSafeWithdrawAmountUnderlying(b.SafeWithdrawAmountUnderlying).
		SetCollateralValueInUsd(b.CollateralValueInUsd).
		SetBorrowValueInUsd(b.BorrowValueInUsd).
		SetUnderlyingPriceUsd(b.UnderlyingPriceUsd).
		SetBlockNumber(b.BlockNumber).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &biz.AccountCTokenReply{
		ID: m.ID,
	}, nil
}

func (r *accountCTokenRepo) UpdateOneAccountCToken(ctx context.Context, id int, b *biz.AccountCToken) (*biz.AccountCTokenReply, error) {
	m, err := r.data.AccountCToken(ctx).UpdateOneID(id).
		SetUser(b.User).
		SetAddress(b.Address).
		SetSymbol(b.Symbol).
		SetEnteredMarket(b.EnteredMarket).
		SetCtokenBalance(b.CtokenBalance).
		SetStoredBorrowBalance(b.StoredBorrowBalance).
		SetBorrowIndex(b.BorrowIndex).
		SetTotalUnderlyingSupplied(b.TotalUnderlyingSupplied).
		SetTotalUnderlyingRedeemed(b.TotalUnderlyingRedeemed).
		SetTotalUnderlyingBorrowed(b.TotalUnderlyingBorrowed).
		SetTotalUnderlyingRepaid(b.TotalUnderlyingRepaid).
		SetSupplyBalanceUnderlying(b.SupplyBalanceUnderlying).
		SetBorrowBalanceUnderlying(b.BorrowBalanceUnderlying).
		SetLifetimeSupplyInterestAccrued(b.LifetimeSupplyInterestAccrued).
		SetLifetimeBorrowInterestAccrued(b.LifetimeBorrowInterestAccrued).
		SetSafeWithdrawAmountUnderlying(b.SafeWithdrawAmountUnderlying).
		SetCollateralValueInUsd(b.CollateralValueInUsd).
		SetBorrowValueInUsd(b.BorrowValueInUsd).
		SetUnderlyingPriceUsd(b.UnderlyingPriceUsd).
		SetBlockNumber(b.BlockNumber).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &biz.AccountCTokenReply{
		ID: m.ID,
	}, nil
}

func (r *accountCTokenRepo) bizToData(ctoken *ent.AccountCToken, biz *biz.AccountCToken) {
	ctoken.ID = biz.ID
	ctoken.User = biz.User
	ctoken.Address = biz.Address
	ctoken.Symbol = biz.Symbol
	ctoken.EnteredMarket = biz.EnteredMarket
	ctoken.CtokenBalance = biz.CtokenBalance
	ctoken.StoredBorrowBalance = biz.StoredBorrowBalance
	ctoken.BorrowIndex = biz.BorrowIndex
	ctoken.TotalUnderlyingSupplied = biz.TotalUnderlyingSupplied
	ctoken.TotalUnderlyingRedeemed = biz.TotalUnderlyingRedeemed
	ctoken.TotalUnderlyingBorrowed = biz.TotalUnderlyingBorrowed
	ctoken.TotalUnderlyingRepaid = biz.TotalUnderlyingRepaid
	ctoken.SupplyBalanceUnderlying = biz.SupplyBalanceUnderlying
	ctoken.BorrowBalanceUnderlying = biz.BorrowBalanceUnderlying
	ctoken.LifetimeSupplyInterestAccrued = biz.LifetimeSupplyInterestAccrued
	ctoken.LifetimeBorrowInterestAccrued = biz.LifetimeBorrowInterestAccrued
	ctoken.SafeWithdrawAmountUnderlying = biz.SafeWithdrawAmountUnderlying
	ctoken.CollateralValueInUsd = biz.CollateralValueInUsd
	ctoken.BorrowValueInUsd = biz.BorrowValueInUsd
	ctoken.UnderlyingPriceUsd = biz.UnderlyingPriceUsd
	ctoken.BlockNumber = biz.BlockNumber
}
