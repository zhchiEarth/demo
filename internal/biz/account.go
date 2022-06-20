package biz

import (
	v1 "compound/api/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
)

type Account struct {
	Address                   string
	HasBorrowed               bool
	CountLiquidated           uint32
	CountLiquidator           uint32
	Health                    string
	TotalCollateralValueInUsd string
	TotalBorrowValueInUsd     string
	BlockNumber               uint64
	Tokens                    []AccountCToken
}

type AccountReply struct {
	ID int
}

type AccountListReq struct {
	Addresses           []string
	BlockNumber         uint64
	MaxHealth           float64
	MinBorrowValueInEth float64
	PageNumber          int64
	PageSize            int64
}

type AccountListReply struct {
	Count   int64
	Account []*Account
}

type AccountRepo interface {
	List(ctx context.Context, req *AccountListReq) (*AccountListReply, error)
	SaveAccount(ctx context.Context, b *Account) (*AccountReply, error)
	GetAccount(ctx context.Context, address string) (*Account, error)
	GetUsersOraclePriceChanged(ctx context.Context, cToken string, underlyingPriceUsd string, pageNum, pageSize int64) ([]string, error)
}

type AccountUseCase struct {
	log         *log.Helper
	accountRepo AccountRepo
	cTokenRepo  AccountCTokenRepo
	tm          Transaction
}

func NewAccountUseCase(accountRepo AccountRepo, cTokenRepo AccountCTokenRepo, tm Transaction, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{
		accountRepo: accountRepo,
		cTokenRepo:  cTokenRepo,
		tm:          tm,
		log:         log.NewHelper(log.With(logger, "module", "biz/account")),
	}
}

func (uc *AccountUseCase) List(ctx context.Context, in *v1.AccountListRequest) (*AccountListReply, error) {
	// 将address 转换成小写
	var addresses []string
	for _, addr := range in.Addresses {
		addresses = append(addresses, strings.ToLower(addr))
	}
	req := &AccountListReq{
		Addresses: addresses,
		//BlockNumber: in.B
		MaxHealth:           in.MaxHealth,
		MinBorrowValueInEth: in.MinBorrowValueInEth,
		PageNumber:          int64(in.PageNumber),
		PageSize:            int64(in.PageSize),
	}
	return uc.accountRepo.List(ctx, req)
}

func (uc *AccountUseCase) SaveAccountAndCToken(ctx context.Context, a *Account) (*AccountReply, error) {
	var (
		reply *AccountReply
		err   error
	)
	if e := uc.tm.ExecTx(ctx, func(ctx context.Context) error {
		reply, err = uc.accountRepo.SaveAccount(ctx, a)
		if err != nil {
			return err
		}
		for _, cToken := range a.Tokens {
			if _, e := uc.cTokenRepo.SaveCToken(ctx, reply.ID, &cToken); err != nil {
				return e
			}
		}
		return nil
	}); e != nil {
		return reply, err
	}
	return reply, nil
}

func (uc *AccountUseCase) GetUsersOraclePriceChanged(ctx context.Context, cToken string, underlyingPriceUsd string, pageNum,
	pageSize int64) ([]string, error) {
	return uc.accountRepo.GetUsersOraclePriceChanged(ctx, cToken, underlyingPriceUsd, pageNum, pageSize)
}

func (uc *AccountUseCase) GetAccount(ctx context.Context, address string) (*Account, error) {
	return uc.accountRepo.GetAccount(ctx, address)
}
