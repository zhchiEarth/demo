package data

import (
	"compound/internal/biz"
	"compound/internal/data/ent"
	"compound/internal/data/ent/account"
	"compound/internal/data/ent/accountctoken"
	"compound/internal/data/ent/predicate"
	"compound/pkg/util/pagination"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"strconv"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/account")),
	}
}

func (r *accountRepo) SaveAccount(ctx context.Context, b *biz.Account) (*biz.AccountReply, error) {
	a, err := r.data.Account(ctx).Query().
		Where(account.Address(b.Address)).
		First(ctx)
	if !ent.IsNotFound(err) && err != nil {
		return nil, errors.Wrap(err, "SaveAccount 保存错误")
	}
	// 区块高度 大于数据库的区块高度
	if a == nil {
		return r.CreateAccount(ctx, b)
	}
	return r.UpdateOneAccount(ctx, a.ID, b)
	//if b.BlockNumber > a.BlockNumber {
	//	return r.UpdateOneAccount(ctx, a.ID, b)
	//}
	//return &biz.AccountReply{
	//	ID: a.ID,
	//}, nil
}

func (r *accountRepo) CreateAccount(ctx context.Context, b *biz.Account) (*biz.AccountReply, error) {
	m, err := r.data.Account(ctx).Create().
		SetAddress(b.Address).
		SetHasBorrowed(b.HasBorrowed).
		SetCountLiquidator(b.CountLiquidator).
		SetCountLiquidated(b.CountLiquidated).
		SetHealth(b.Health).
		SetTotalCollateralValueInUsd(b.TotalCollateralValueInUsd).
		SetTotalBorrowValueInUsd(b.TotalBorrowValueInUsd).
		SetBlockNumber(b.BlockNumber).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "CreateAccount 新增错误")
	}
	return &biz.AccountReply{
		ID: m.ID,
	}, nil
}

func (r *accountRepo) UpdateOneAccount(ctx context.Context, id int, b *biz.Account) (*biz.AccountReply, error) {
	m, err := r.data.Account(ctx).UpdateOneID(id).
		SetAddress(b.Address).
		SetHasBorrowed(b.HasBorrowed).
		SetCountLiquidator(b.CountLiquidator).
		SetCountLiquidated(b.CountLiquidated).
		SetHealth(b.Health).
		SetTotalCollateralValueInUsd(b.TotalCollateralValueInUsd).
		SetTotalBorrowValueInUsd(b.TotalBorrowValueInUsd).
		SetBlockNumber(b.BlockNumber).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateOneAccount 错误")
	}
	return &biz.AccountReply{
		ID: m.ID,
	}, nil
}

func (r *accountRepo) List(ctx context.Context, req *biz.AccountListReq) (*biz.AccountListReply, error) {
	query := r.data.Account(ctx).Query()
	var predicates []predicate.Account
	if req.MaxHealth > 0 {
		maxHealth := strconv.FormatFloat(req.MaxHealth, 'f', -1, 64)
		// 小于MaxHealth 大于0
		predicates = append(predicates, account.HealthLTE(maxHealth), account.HealthGT("0"))
	}
	if len(req.Addresses) > 0 {
		predicates = append(predicates, account.AddressIn(req.Addresses...))
	}
	count, err := query.Where(predicates...).Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "List 错误")
	}
	rv := make([]*biz.Account, 0)

	if count > 0 {
		os, err := query.
			Where(predicates...).
			WithTokens().
			Limit(int(req.PageSize)).Offset(int(pagination.GetPageOffset(req.PageNumber, req.PageSize))).All(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "query 失败")
		}

		for _, o := range os {
			ac := r.dataToBiz(o)
			for _, token := range o.Edges.Tokens {
				t := r.tokenToBiz(token)
				ac.Tokens = append(ac.Tokens, t)
			}
			rv = append(rv, &ac)
		}
	}
	return &biz.AccountListReply{
		Count:   int64(count),
		Account: rv,
	}, nil
}

func (r *accountRepo) GetAccount(ctx context.Context, address string) (*biz.Account, error) {
	a, err := r.data.Account(ctx).Query().
		WithTokens().
		Where(account.Address(address)).
		First(ctx)
	if !ent.IsNotFound(err) && err != nil {
		r.log.Error(err)
		return nil, errors.Wrap(err, "")
	}
	ac := r.dataToBiz(a)
	for _, token := range a.Edges.Tokens {
		t := r.tokenToBiz(token)
		ac.Tokens = append(ac.Tokens, t)
	}
	return &ac, nil
}

// GetUsersOraclePriceChanged 获取预言机价格有改变的用户 /**
func (r *accountRepo) GetUsersOraclePriceChanged(ctx context.Context, cToken string, underlyingPriceUsd string, pageNum, pageSize int64) ([]string, error) {
	os, err := r.data.Account(ctx).Query().
		Select("address").
		Where(
			account.HasBorrowed(true),
			account.HasTokensWith(
				accountctoken.Address(cToken),
				accountctoken.UnderlyingPriceUsdNEQ(underlyingPriceUsd),
				accountctoken.BorrowBalanceUnderlyingGT("0"),
			),
		).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		GroupBy(account.FieldAddress).
		Strings(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return os, nil
}

func (r *accountRepo) tokenToBiz(ctoken *ent.AccountCToken) biz.AccountCToken {
	return biz.AccountCToken{
		ID:                            ctoken.ID,
		User:                          ctoken.User,
		Address:                       ctoken.Address,
		Symbol:                        ctoken.Symbol,
		EnteredMarket:                 ctoken.EnteredMarket,
		CtokenBalance:                 ctoken.CtokenBalance,
		StoredBorrowBalance:           ctoken.StoredBorrowBalance,
		BorrowIndex:                   ctoken.BorrowIndex,
		TotalUnderlyingSupplied:       ctoken.TotalUnderlyingSupplied,
		TotalUnderlyingRedeemed:       ctoken.TotalUnderlyingRedeemed,
		TotalUnderlyingBorrowed:       ctoken.TotalUnderlyingBorrowed,
		TotalUnderlyingRepaid:         ctoken.TotalUnderlyingRepaid,
		SupplyBalanceUnderlying:       ctoken.SupplyBalanceUnderlying,
		BorrowBalanceUnderlying:       ctoken.BorrowBalanceUnderlying,
		LifetimeSupplyInterestAccrued: ctoken.LifetimeSupplyInterestAccrued,
		LifetimeBorrowInterestAccrued: ctoken.LifetimeBorrowInterestAccrued,
		SafeWithdrawAmountUnderlying:  ctoken.SafeWithdrawAmountUnderlying,
		CollateralValueInUsd:          ctoken.CollateralValueInUsd,
		BorrowValueInUsd:              ctoken.BorrowValueInUsd,
		UnderlyingPriceUsd:            ctoken.UnderlyingPriceUsd,
		BlockNumber:                   ctoken.BlockNumber,
	}
}

func (r *accountRepo) bizToData(account *ent.Account, biz *biz.Account) {
	account.Address = biz.Address
	account.HasBorrowed = biz.HasBorrowed
	account.CountLiquidator = biz.CountLiquidator
	account.CountLiquidated = biz.CountLiquidated
	account.Health = biz.Health
	account.TotalCollateralValueInUsd = biz.TotalCollateralValueInUsd
	account.TotalBorrowValueInUsd = biz.TotalBorrowValueInUsd
	account.BlockNumber = biz.BlockNumber
}

func (r *accountRepo) dataToBiz(account *ent.Account) biz.Account {
	b := biz.Account{
		Address:                   account.Address,
		Health:                    account.Health,
		HasBorrowed:               account.HasBorrowed,
		CountLiquidator:           account.CountLiquidator,
		CountLiquidated:           account.CountLiquidated,
		TotalCollateralValueInUsd: account.TotalCollateralValueInUsd,
		TotalBorrowValueInUsd:     account.TotalBorrowValueInUsd,
		BlockNumber:               account.BlockNumber,
	}
	return b
}
