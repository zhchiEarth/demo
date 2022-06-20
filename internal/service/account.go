package service

import (
	v1 "compound/api/v1"
	"compound/internal/biz"
	"math"

	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AccountService struct {
	v1.UnimplementedAccountServer

	uc *biz.AccountUseCase

	log *log.Helper
}

// NewAccountService new an account service.
func NewAccountService(auc *biz.AccountUseCase, logger log.Logger) *AccountService {
	return &AccountService{
		uc:  auc,
		log: log.NewHelper(log.With(logger, "module", "service/graph")),
	}
}

// List  implements
func (s *AccountService) List(ctx context.Context, in *v1.AccountListRequest) (*v1.AccountListReply, error) {
	if in.PageNumber == 0 {
		in.PageNumber = 1
	}
	if in.PageSize == 0 {
		in.PageSize = 10
	}

	rep, err := s.uc.List(ctx, in)
	if err != nil {
		return nil, err
	}
	rs := make([]*v1.AccountListReply_Account, 0)
	for _, x := range rep.Account {
		var tokens []*v1.Token
		for _, token := range x.Tokens {
			t := &v1.Token{
				Address:                       token.Address,
				Symbol:                        token.Symbol,
				SupplyBalanceUnderlying:       token.SupplyBalanceUnderlying,
				BorrowBalanceUnderlying:       token.BorrowValueInUsd,
				LifetimeSupplyInterestAccrued: token.LifetimeSupplyInterestAccrued,
				LifetimeBorrowInterestAccrued: token.LifetimeBorrowInterestAccrued,
				SafeWithdrawAmountUnderlying:  token.SafeWithdrawAmountUnderlying,
			}
			tokens = append(tokens, t)
		}
		rs = append(rs, &v1.AccountListReply_Account{
			Address:                   x.Address,
			Health:                    x.Health,
			TotalBorrowValueInEth:     x.TotalBorrowValueInUsd,
			TotalCollateralValueInEth: x.TotalCollateralValueInUsd,
			Tokens:                    tokens,
		})
	}
	totalPages := math.Ceil(float64(rep.Count) / float64(in.PageSize))
	return &v1.AccountListReply{
		Accounts: rs,
		Request:  in,
		PaginationSummary: &v1.PaginationSummary{
			PageNumber:   in.PageNumber,
			PageSize:     in.PageSize,
			TotalEntries: uint32(rep.Count),
			TotalPages:   uint32(totalPages),
		},
	}, nil
}

//func (s *AccountService) getMarketMap(ctx context.Context) (map[string]*biz.Market, error) {
//	markets, err := s.muc.All(ctx)
//	if err != nil {
//		return nil, err
//	}
//	marketMap := make(map[string]*biz.Market, len(markets))
//	for _, market := range markets {
//		_, ok := marketMap[market.Address]
//		if !ok {
//			marketMap[market.Address] = market
//		}
//	}
//	return marketMap, nil
//}

// SyncAccounts  同步用户
//func (s *AccountService) Handle(ctx context.Context) {
//	var query struct {
//		Accounts []Account `graphql:"accounts(first: 1000)"`
//	}
//	err := s.cli.Query(context.Background(), &query, nil)
//	if err != nil {
//		s.log.Error(err)
//		return
//	}
//
//	marketMap, err := s.getMarketMap(ctx)
//	if err != nil {
//		s.log.Error(err)
//		return
//	}
//
//	for _, account := range query.Accounts {
//		//var tokens []biz.AccountCToken
//		blockNumber := int64(0)
//		totalCollateralValueInUsd := decimal.NewFromInt(0)
//		totalBorrowValueInUsd := decimal.NewFromInt(0)
//		for _, token := range account.Tokens {
//			ctoken, err := s.getAccountCToken(marketMap, token)
//			if err != nil {
//				s.log.Error(err)
//				break
//			}
//			totalCollateralValueInUsd = totalCollateralValueInUsd.Add(ctoken.CollateralValueInUsd)
//			totalBorrowValueInUsd = totalBorrowValueInUsd.Add(ctoken.BorrowValueInUsd)
//
//			_, err = s.cuc.Save(ctx, ctoken)
//			if err != nil {
//				s.log.Error(err)
//				break
//			}
//			//tokens = append(tokens, t)
//			if ctoken.BlockNumber > blockNumber {
//				blockNumber = ctoken.BlockNumber
//			}
//		}
//		health := decimal.NewFromInt(0)
//		if totalCollateralValueInUsd.Cmp(decimal.NewFromInt(0)) != 0 && totalBorrowValueInUsd.Cmp(decimal.NewFromInt(0)) != 0 {
//			health = totalCollateralValueInUsd.Div(totalBorrowValueInUsd)
//		}
//
//		_, err := s.uc.SaveAccountAndCToken(ctx, &biz.Account{
//			Address:                   string(account.Id),
//			Health:                    health,
//			TotalCollateralValueInUsd: totalCollateralValueInUsd,
//			TotalBorrowValueInUsd:     totalBorrowValueInUsd,
//			BlockNumber:               blockNumber,
//			//Tokens: tokens,
//		})
//		if err != nil {
//			s.log.Error(err)
//			break
//		}
//	}
//}

// 同步预言机价格
//func (s *AccountService) SyncOraclePrice(ctx context.Context) {
//	cli, _ := rpc.NewClient("https://api.s0.b.hmny.io")
//	//cli, _ := rpc.NewClient("https://localhost:8545")
//	defer cli.Close()
//	// harmony
//	oracleContractAddress := common.HexToAddress("0x7105418ABAFbFd54b74954d5c65AE0E9110c078D")
//	//oracleContractAddress := common.HexToAddress("0x6df290fb9ef0ec910bcbe5eb07dd1fbe4ee7c1f3")
//	oracleContract, err := oracle.NewOracle(oracleContractAddress, cli)
//
//	marketMap, err := s.getMarketMap(ctx)
//	if err != nil {
//		s.log.Error(err)
//		return
//	}
//
//	callOpts := bind.CallOpts{Pending: false}
//	for _, market := range marketMap {
//		price, err := oracleContract.GetUnderlyingPrice(&callOpts, common.HexToAddress(market.Address))
//		if err != nil {
//			s.log.Error("", err)
//			return
//		}
//		mantissa := decimal.NewFromBigInt(big.NewInt(10), 18+18-market.UnderlyingDecimals)
//		priceBD := decimal.NewFromBigInt(price, 1)
//		priceInUSD := priceBD.Div(mantissa)
//		//underlyingPriceUSD, err := decimal.NewFromString(market.UnderlyingPriceUSD)
//		//if err != nil {
//		//    s.log.Error("", err)
//		//    return
//		//}
//		// 价格相同 不需要更新
//		//if priceBD.Cmp(underlyingPriceUSD) == 0 {
//		//    break
//		//}
//
//		list, err := s.cuc.GetAccountListByCToken(ctx, market.Address)
//		if err != nil {
//			s.log.Error("", err)
//			break
//		}
//
//		for _, item := range list {
//			err = s.updateAccountAndCToken(market, item, priceInUSD)
//			if err != nil {
//				s.log.Error("", err)
//				break
//			}
//		}
//	}
//}

//func (s *AccountService) getAccountCToken(markets map[string]*biz.Market, token Token) (*biz.AccountCToken, error) {
//	strArr := strings.Split(string(token.Id), "-")
//	if len(strArr) != 2 {
//		return nil, errors.New("token ID 分割失败" + string(token.Id))
//	}
//	block, err := strconv.ParseInt(string(token.AccrualBlockNumber), 10, 64)
//	if err != nil {
//		return nil, err
//	}
//
//	ctoken := &biz.AccountCToken{
//		User:    strArr[1],
//		Address: strArr[0],
//		Symbol:  string(token.Symbol),
//		//SupplyBalanceUnderlying:       supplyBalanceUnderlying.String(),
//		//BorrowBalanceUnderlying:       string(token.StoredBorrowBalance),
//		//CollateralValueInUsd: "",
//		//BorrowValueInUsd: "",
//		LifetimeSupplyInterestAccrued: "",
//		LifetimeBorrowInterestAccrued: "",
//		SafeWithdrawAmountUnderlying:  "",
//		BlockNumber:                   block,
//	}
//	market, ok := markets[ctoken.Address]
//	if !ok {
//		return nil, errors.New("市场获取失败, address: " + string(ctoken.Address))
//	}
//	underlyingPriceUSD, err := decimal.NewFromString(market.UnderlyingPriceUSD)
//	if err != nil {
//		s.log.Error("underlyingPriceUSD 转换float失败, err：", err)
//		return nil, err
//	}
//	exchangeRate, err := decimal.NewFromString(market.ExchangeRate)
//	if err != nil {
//		s.log.Error("exchangeRate 转换float失败, err：", err)
//		return nil, err
//	}
//	collateralFactor, err := decimal.NewFromString(market.CollateralFactor)
//	if err != nil {
//		s.log.Error("collateralFactor 转换float失败, err：", err)
//		return nil, err
//	}
//	ctokenBalance, err := decimal.NewFromString(string(token.CTokenBalance))
//	if err != nil {
//		s.log.Error("ctokenBalance 转换big float失败, err：", err)
//		return nil, err
//	}
//	ctoken.BorrowBalanceUnderlying, err = decimal.NewFromString(string(token.StoredBorrowBalance))
//	if err != nil {
//		s.log.Error("borrowBalanceUnderlying 转换big float失败, err：", err)
//		return nil, err
//	}
//	ctoken.SupplyBalanceUnderlying = ctokenBalance.Mul(exchangeRate)
//	// 转换成 underlyingPriceUSD * supplyBalanceUnderlying
//	collateralValue := ctoken.SupplyBalanceUnderlying.Mul(underlyingPriceUSD)
//	// 抵押多少价值
//	ctoken.CollateralValueInUsd = collateralValue.Mul(collateralFactor)
//	ctoken.BorrowValueInUsd = ctoken.BorrowBalanceUnderlying.Mul(underlyingPriceUSD)
//
//	return ctoken, nil
//}

// 计算抵押物价值
//func (s *AccountService) calculateCollateralValue(supplyBalanceUnderlying decimal.Decimal, collateralFactor string,
//	underlyingPriceUSD decimal.Decimal) (decimal.Decimal, error) {
//	zero := decimal.New(0, 1)
//	collateralFactorBD, err := decimal.NewFromString(collateralFactor)
//	if err != nil {
//		s.log.Error("collateralFactor 转换float失败, err：", err)
//		return zero, err
//	}
//	// 转换成 underlyingPriceUSD * supplyBalanceUnderlying
//	collateralValue := supplyBalanceUnderlying.Mul(underlyingPriceUSD)
//	// 抵押多少价值
//	collateralValueInUsd := collateralValue.Mul(collateralFactorBD)
//	return collateralValueInUsd, nil
//}
//
//func (s *AccountService) updateAccountAndCToken(market *biz.Market, item *biz.AccountAndCToken, priceBD decimal.Decimal) (err error) {
//	if decimal.Zero.Cmp(item.CToken.CollateralValueInUsd) == 0 || decimal.Zero.Cmp(item.CToken.BorrowValueInUsd) == 0 {
//		return nil
//	}
//	collateralValueInUsd, err := s.calculateCollateralValue(item.CToken.SupplyBalanceUnderlying, market.CollateralFactor, priceBD)
//	if err != nil {
//		s.log.Error("collateralValueInUsd 计算失败, err：", err)
//		return err
//	}
//	collateralValueInUsdOld := item.CToken.CollateralValueInUsd
//	borrowValueInUsdOld := item.CToken.BorrowValueInUsd
//
//	item.CToken.CollateralValueInUsd = collateralValueInUsd
//	item.CToken.BorrowValueInUsd = item.CToken.BorrowBalanceUnderlying.Mul(priceBD)
//
//	// 减去旧的 + 新的值
//	item.Account.TotalCollateralValueInUsd = item.Account.TotalCollateralValueInUsd.Sub(collateralValueInUsdOld).Add(item.CToken.CollateralValueInUsd)
//	item.Account.TotalBorrowValueInUsd = item.Account.TotalBorrowValueInUsd.Sub(borrowValueInUsdOld).Add(item.CToken.BorrowValueInUsd)
//
//	item.Account.Health = decimal.Zero
//	if decimal.Zero.Cmp(item.Account.TotalCollateralValueInUsd) != 0 && decimal.Zero.Cmp(item.Account.TotalBorrowValueInUsd) != 0 {
//		item.Account.Health = item.Account.TotalCollateralValueInUsd.Div(item.Account.TotalBorrowValueInUsd)
//	}
//	_, err = s.uc.SaveAccountAndCToken(context.Background(), &item.Account)
//	if err != nil {
//		s.log.Error("SaveAccountAndCToken 保存失败, err：", err)
//		return err
//	}
//	_, err = s.cuc.Save(context.Background(), &item.CToken)
//	if err != nil {
//		s.log.Error("SaveAccountAndCToken 保存失败, err：", err)
//		return err
//	}
//	return nil
//}
