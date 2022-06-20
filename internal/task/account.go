package task

import (
	"compound/internal/biz"
	"compound/internal/util"
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hasura/go-graphql-client"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type Account struct {
	Id              graphql.String  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HasBorrowed     graphql.Boolean `protobuf:"varint,2,opt,name=has_borrowed,json=hasBorrowed,proto3" json:"has_borrowed,omitempty"`
	CountLiquidated graphql.Int     `protobuf:"bytes,3,opt,name=count_liquidated,json=countLiquidated,proto3" json:"count_liquidated,omitempty"`
	CountLiquidator graphql.Int     `protobuf:"bytes,4,opt,name=count_liquidator,json=countLiquidator,proto3" json:"count_liquidator,omitempty"`
	//Tokens          []Token         `protobuf:"bytes,5,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

type AccountCToken struct {
	Id                      graphql.String  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Symbol                  graphql.String  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	AccrualBlockNumber      graphql.String  `protobuf:"varint,3,opt,name=accrual_block_number,json=accrualBlockNumber,proto3" json:"accrual_block_number,omitempty"`
	EnteredMarket           graphql.Boolean `protobuf:"varint,4,opt,name=entered_market,json=enteredMarket,proto3" json:"entered_market,omitempty"`
	CTokenBalance           graphql.String  `protobuf:"bytes,5,opt,name=c_token_balance,json=cTokenBalance,proto3" json:"c_token_balance,omitempty"`
	TotalUnderlyingSupplied graphql.String  `protobuf:"bytes,6,opt,name=total_underlying_supplied,json=totalUnderlyingSupplied,proto3" json:"total_underlying_supplied,omitempty"`
	TotalUnderlyingRedeemed graphql.String  `protobuf:"bytes,7,opt,name=total_underlying_redeemed,json=totalUnderlyingRedeemed,proto3" json:"total_underlying_redeemed,omitempty"`
	TotalUnderlyingBorrowed graphql.String  `protobuf:"bytes,8,opt,name=total_underlying_borrowed,json=totalUnderlyingBorrowed,proto3" json:"total_underlying_borrowed,omitempty"`
	AccountBorrowIndex      graphql.String  `protobuf:"bytes,9,opt,name=account_borrow_index,json=accountBorrowIndex,proto3" json:"account_borrow_index,omitempty"`
	TotalUnderlyingRepaid   graphql.String  `protobuf:"bytes,10,opt,name=total_underlying_repaid,json=totalUnderlyingRepaid,proto3" json:"total_underlying_repaid,omitempty"`
	StoredBorrowBalance     graphql.String  `protobuf:"bytes,11,opt,name=stored_borrow_balance,json=storedBorrowBalance,proto3" json:"stored_borrow_balance,omitempty"`
	Account                 Account
}

type AccountTask struct {
	uc   *biz.AccountUseCase
	puc  *biz.PreferenceUseCase
	muc  *biz.MarketUseCase
	log  *log.Helper
	task *Task
	job  *gocron.Job
	p    *ants.PoolWithFunc
}

type HealthPrams struct {
	Markets map[string]*biz.Market
	Address string
}

// NewAccountTask new an account task.
func NewAccountTask(t *Task, uc *biz.AccountUseCase, puc *biz.PreferenceUseCase, muc *biz.MarketUseCase,
	logger log.Logger) (*AccountTask, func()) {

	a := &AccountTask{
		uc:   uc,
		log:  log.NewHelper(log.With(logger, "module", "task/account")),
		task: t,
		muc:  muc,
		puc:  puc,
	}

	// 新建线程池
	p, _ := ants.NewPoolWithFunc(10000, func(i interface{}) {
		p, ok := i.(HealthPrams)
		if ok {
			a.calcAccountHealth(p.Markets, p.Address)
			return
		}
	})
	a.p = p
	return a, func() {
		p.Release()
	}
}

func (t *AccountTask) SetJob(job *gocron.Job) {
	t.job = job
}

func (t *AccountTask) getLastBlockNumber() uint64 {
	lastBlockNumber, err := t.puc.GetAccountCTokenLastBlockNumber(context.Background())
	if err != nil {
		t.log.Errorf("获取 lastBlockNumber 失败, %+v", err)
		return 0
	}
	return lastBlockNumber
}

// Handle 同步 account
func (t *AccountTask) Handle() {
	marketMap, err := t.muc.GetMarketMap(context.Background())
	if err != nil {
		t.log.Errorf("获取 getMarketMap 失败：%+v", err)
		return
	}
	if len(marketMap) == 0 {
		t.log.Info("marketMap 是空的")
		return
	}

	lastBlockNumber := t.getLastBlockNumber()
	var query struct {
		AccountCTokens []AccountCToken `graphql:"accountCTokens(first: 100,orderBy: accrualBlockNumber, orderDirection: asc, where: {accrualBlockNumber_gte: $lastBlockNumber})"`
	}

	variables := map[string]interface{}{
		"lastBlockNumber": graphql.Int(lastBlockNumber),
	}

	err = t.task.graphqlCli.Query(context.Background(), &query, variables)
	if err != nil {
		t.log.Errorf("查询accountCTokens 失败 :  %+v", errors.Wrap(err, ""))
		return
	}

	for _, cToken := range query.AccountCTokens {
		accrualBlockNumber, err := strconv.ParseUint(string(cToken.AccrualBlockNumber), 10, 64)
		if err != nil {
			t.log.Error(errors.Wrap(err, "accrualBlockNumber 转换uint64失败"))
			return
		}
		if lastBlockNumber < accrualBlockNumber {
			lastBlockNumber = accrualBlockNumber
		}
		account, err := t.toBizAccount(cToken)
		if err != nil {
			t.log.Errorf("账户转换失败:%+v", err)
			break
		}
		reply, err := t.uc.SaveAccountAndCToken(context.Background(), account)
		if err != nil {
			t.log.Errorf("SaveAccountAndCToken %+v", err)
			return
		}
		// 更新成功  重新计算，用户健康值
		if reply.ID > 0 {
			t.calcAccountHealth(marketMap, account.Address)
		}
	}
	_, err = t.puc.SaveAccountCTokenLastBlockNumber(context.Background(), lastBlockNumber)
	if err != nil {
		t.log.Errorf("保存SaveAccountCTokenLastBlockNumber 失败:%+v", err)
	}
}

func (t *AccountTask) toBizAccount(cToken AccountCToken) (*biz.Account, error) {
	strArr := strings.Split(string(cToken.Id), "-")
	if len(strArr) != 2 {
		return nil, errors.New("token ID 分割失败" + string(cToken.Id))
	}
	block, err := strconv.ParseUint(string(cToken.AccrualBlockNumber), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "AccrualBlockNumber转换失败")
	}

	ct := biz.AccountCToken{
		User:                          strArr[1],
		Address:                       strArr[0],
		Symbol:                        string(cToken.Symbol),
		EnteredMarket:                 bool(cToken.EnteredMarket),
		CtokenBalance:                 string(cToken.CTokenBalance),
		StoredBorrowBalance:           string(cToken.StoredBorrowBalance),
		BorrowIndex:                   string(cToken.AccountBorrowIndex),
		TotalUnderlyingSupplied:       string(cToken.TotalUnderlyingSupplied),
		TotalUnderlyingRedeemed:       string(cToken.TotalUnderlyingRedeemed),
		TotalUnderlyingBorrowed:       string(cToken.TotalUnderlyingBorrowed),
		TotalUnderlyingRepaid:         string(cToken.TotalUnderlyingRepaid),
		SupplyBalanceUnderlying:       "",
		BorrowBalanceUnderlying:       "",
		CollateralValueInUsd:          "",
		BorrowValueInUsd:              "",
		LifetimeSupplyInterestAccrued: "",
		LifetimeBorrowInterestAccrued: "",
		SafeWithdrawAmountUnderlying:  "",
		BlockNumber:                   block,
	}

	a := &biz.Account{
		Address:                   string(cToken.Account.Id),
		HasBorrowed:               bool(cToken.Account.HasBorrowed),
		CountLiquidated:           uint32(cToken.Account.CountLiquidated),
		CountLiquidator:           uint32(cToken.Account.CountLiquidated),
		Health:                    "",
		TotalCollateralValueInUsd: "",
		TotalBorrowValueInUsd:     "",
		//BlockNumber:               blockNumber,
		Tokens: []biz.AccountCToken{ct},
	}
	return a, nil
}

func (t *AccountTask) InvokePool(markets map[string]*biz.Market, address string) {
	err := t.p.Invoke(HealthPrams{
		markets, address,
	})
	if err != nil {
		t.log.Errorf("InvokePool 执行失败:err:%+v", errors.Wrap(err, ""))
	}
}

// CalcAccountHealth  根据用户地址 计算用户的健康度
// 获取用户的所有资产，进行计算。每次都要重写计算，计算部分行不行，比如 只计算预言机改变了的资产呢?
// 部分计算是不对的，因为市场的兑换利率一直在变化。就算这个用户的存款和贷款没有发生变化，但是存款是根据兑换率变化的，
// 而借款每次都会有利息的， 用户的存款和借款相对应的会发生变化。重新计算相对于会准确写.
func (t *AccountTask) calcAccountHealth(markets map[string]*biz.Market, address string) {
	account, err := t.uc.GetAccount(context.Background(), address)
	if err != nil {
		t.log.Errorf(" GetAccount失败 address:%, err:%+v", address, err)
		return
	}

	var cTokens []biz.AccountCToken
	totalCollateralValueInUsd := decimal.NewFromInt(0)
	totalBorrowValueInUsd := decimal.NewFromInt(0)
	lastBlockNumber := uint64(0)
	for _, cToken := range account.Tokens {
		if lastBlockNumber < cToken.BlockNumber {
			lastBlockNumber = cToken.BlockNumber
		}

		m, ok := markets[cToken.Address]
		if !ok {
			t.log.Errorf("markets 不存在, cToken address:%s", cToken.Address)
			return
		}

		supplyBalanceUnderlying, err := util.CalcSupplyBalanceUnderlying(cToken.CtokenBalance, m.ExchangeRate)
		if err != nil {
			t.log.Errorf("CalcSupplyBalanceUnderlying 计算失败:ctokenBalance%s, ExchangeRate:% , err:",
				cToken.CtokenBalance, m.ExchangeRate, err)
			return
		}

		collateralValueInUsd, err := util.CalcCollateralValueInUsd(supplyBalanceUnderlying.String(), m.CollateralFactor, m.UnderlyingPriceUSD)
		if err != nil {
			t.log.Errorf("CalcCollateralValueInUsd 计算失败:supplyBalanceUnderlying:%s  CollateralFactor:% UnderlyingPriceUSD:%s, err:%+v",
				supplyBalanceUnderlying.String(), m.CollateralFactor, m.UnderlyingPriceUSD, err)
			return
		}

		borrowBalanceUnderlying, err := util.CalcBorrowBalanceUnderlying(cToken.StoredBorrowBalance, cToken.BorrowIndex, m.BorrowIndex)
		if err != nil {
			t.log.Errorf("CalcBorrowBalanceUnderlying 计算失败:StoredBorrowBalance:%s  BorrowIndex:% BorrowIndex:%s, err:%+v",
				cToken.StoredBorrowBalance, cToken.BorrowIndex, m.BorrowIndex, err)
			return
		}

		borrowValueInUsd, err := util.CalcBorrowValueInUsd(borrowBalanceUnderlying.String(), m.UnderlyingPriceUSD)
		if err != nil {
			t.log.Errorf("CalcBorrowValueInUsd 计算失败:borrowBalanceUnderlying:%s  UnderlyingPriceUSD:% err:%+v",
				borrowBalanceUnderlying.String(), m.UnderlyingPriceUSD, err)
			return
		}

		cToken.SupplyBalanceUnderlying = supplyBalanceUnderlying.String()
		cToken.BorrowBalanceUnderlying = borrowBalanceUnderlying.String()
		cToken.UnderlyingPriceUsd = m.UnderlyingPriceUSD
		cToken.CollateralValueInUsd = collateralValueInUsd.String()
		cToken.BorrowValueInUsd = borrowValueInUsd.String()
		cTokens = append(cTokens, cToken)

		// 进入了市场 才允许借贷
		if cToken.EnteredMarket {
			totalCollateralValueInUsd = totalCollateralValueInUsd.Add(collateralValueInUsd)
			totalBorrowValueInUsd = totalBorrowValueInUsd.Add(borrowValueInUsd)
		}
	}
	health := decimal.NewFromInt(0)
	if totalCollateralValueInUsd.Cmp(decimal.NewFromInt(0)) != 0 && totalBorrowValueInUsd.Cmp(decimal.NewFromInt(0)) != 0 {
		health = totalCollateralValueInUsd.Div(totalBorrowValueInUsd)
	}
	account.Health = health.String()
	account.TotalBorrowValueInUsd = totalBorrowValueInUsd.String()
	account.TotalCollateralValueInUsd = totalCollateralValueInUsd.String()
	account.Tokens = cTokens
	account.BlockNumber = lastBlockNumber

	_, err = t.uc.SaveAccountAndCToken(context.Background(), account)
	if err != nil {
		t.log.Errorf("计算用户健康值失败: 调用 SaveAccountAndCToken %+v", err)
	}
}
