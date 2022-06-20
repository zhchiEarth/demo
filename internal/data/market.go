package data

import (
	"compound/internal/biz"
	"compound/internal/data/ent"
	"compound/internal/data/ent/market"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var _ biz.MarketRepo = (*marketRepo)(nil)

type marketRepo struct {
	data *Data
	log  *log.Helper
}

func NewMarketRepo(data *Data, logger log.Logger) biz.MarketRepo {
	return &marketRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/market")),
	}
}

func (r *marketRepo) SaveMarket(ctx context.Context, b *biz.Market) (*biz.MarketReply, error) {
	m, err := r.data.Market(ctx).Query().
		Where(market.Address(b.Address)).First(ctx)
	if !ent.IsNotFound(err) && err != nil {
		return nil, errors.Wrap(err, "方法SaveMarket ")
	}
	// 区块高度 大于数据库的区块高度
	if m == nil {
		return r.CreateMarket(ctx, b)
	}
	if b.BlockNumber > m.BlockNumber {
		return r.UpdateOneMarket(ctx, m.ID, b)
	}
	return &biz.MarketReply{
		ID: m.ID,
	}, nil
}

func (r *marketRepo) CreateMarket(ctx context.Context, b *biz.Market) (*biz.MarketReply, error) {
	m, err := r.data.Market(ctx).Create().
		SetAddress(b.Address).
		SetName(b.Name).
		SetSymbol(b.Symbol).
		SetBorrowIndex(b.BorrowIndex).
		SetBorrowRate(b.BorrowRate).
		SetSupplyRate(b.SupplyRate).
		SetCash(b.Cash).
		SetCollateralFactor(b.CollateralFactor).
		SetExchangeRate(b.ExchangeRate).
		SetReserveFactor(b.ReserveFactor).
		SetReserves(b.Reserves).
		SetTotalBorrows(b.TotalBorrows).
		SetTotalSupply(b.TotalSupply).
		SetUnderlyingAddress(b.UnderlyingAddress).
		SetUnderlyingName(b.UnderlyingName).
		SetUnderlyingSymbol(b.UnderlyingSymbol).
		SetUnderlyingPrice(b.UnderlyingPrice).
		SetUnderlyingDecimals(b.UnderlyingDecimals).
		SetUnderlyingPriceUsd(b.UnderlyingPriceUSD).
		SetBlockNumber(b.BlockNumber).
		SetBlockTimestamp(b.BlockTimestamp).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.MarketReply{
		ID: m.ID,
	}, nil
}

func (r *marketRepo) UpdateOneMarket(ctx context.Context, id int, b *biz.Market) (*biz.MarketReply, error) {
	m, err := r.data.Market(ctx).UpdateOneID(id).
		SetAddress(b.Address).
		SetName(b.Name).
		SetSymbol(b.Symbol).
		SetBorrowIndex(b.BorrowIndex).
		SetBorrowRate(b.BorrowRate).
		SetSupplyRate(b.SupplyRate).
		SetCash(b.Cash).
		SetCollateralFactor(b.CollateralFactor).
		SetExchangeRate(b.ExchangeRate).
		SetReserveFactor(b.ReserveFactor).
		SetReserves(b.Reserves).
		SetTotalBorrows(b.TotalBorrows).
		SetTotalSupply(b.TotalSupply).
		SetUnderlyingAddress(b.UnderlyingAddress).
		SetUnderlyingName(b.UnderlyingName).
		SetUnderlyingSymbol(b.UnderlyingSymbol).
		SetUnderlyingPrice(b.UnderlyingPrice).
		SetUnderlyingDecimals(b.UnderlyingDecimals).
		SetUnderlyingPriceUsd(b.UnderlyingPriceUSD).
		SetBlockNumber(b.BlockNumber).
		SetBlockTimestamp(b.BlockTimestamp).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.MarketReply{
		ID: m.ID,
	}, nil
}

func (r *marketRepo) ExistMarket(ctx context.Context, address string) (bool, error) {
	return r.data.Market(ctx).Query().Where(market.AddressEQ(address)).Exist(ctx)
}

func (r *marketRepo) AllMarket(ctx context.Context) ([]*biz.Market, error) {
	list, err := r.data.Market(ctx).Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return r.listToBiz(list), nil
}

func (r *marketRepo) UpdateMarketOraclePrice(ctx context.Context, id int, underlyingPriceUsd string) (*biz.MarketReply, error) {
	m, err := r.data.Market(ctx).UpdateOneID(id).
		SetUnderlyingPriceUsd(underlyingPriceUsd).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &biz.MarketReply{
		ID: m.ID,
	}, nil
}

func (r *marketRepo) bizToData(market *ent.Market, bizMarket *biz.Market) {
	market.ID = bizMarket.ID
	market.Address = bizMarket.Address
	market.Name = bizMarket.Name
	market.Symbol = bizMarket.Symbol
	market.BorrowIndex = bizMarket.BorrowIndex
	market.BorrowRate = bizMarket.BorrowRate
	market.SupplyRate = bizMarket.SupplyRate
	market.Cash = bizMarket.Cash
	market.CollateralFactor = bizMarket.CollateralFactor
	market.ExchangeRate = bizMarket.ExchangeRate
	market.ReserveFactor = bizMarket.ReserveFactor
	market.Reserves = bizMarket.Reserves
	market.TotalBorrows = bizMarket.TotalBorrows
	market.TotalSupply = bizMarket.TotalSupply
	market.UnderlyingAddress = bizMarket.UnderlyingAddress
	market.UnderlyingName = bizMarket.UnderlyingName
	market.UnderlyingSymbol = bizMarket.UnderlyingSymbol
	market.UnderlyingPrice = bizMarket.UnderlyingPrice
	market.UnderlyingDecimals = bizMarket.UnderlyingDecimals
	market.UnderlyingPriceUsd = bizMarket.UnderlyingPriceUSD
	market.BlockNumber = bizMarket.BlockNumber
	market.BlockTimestamp = bizMarket.BlockTimestamp
}

func (r *marketRepo) dataToBiz(entity *ent.Market) *biz.Market {
	return &biz.Market{
		ID:                 entity.ID,
		Address:            entity.Address,
		Name:               entity.Name,
		Symbol:             entity.Symbol,
		BorrowIndex:        entity.BorrowIndex,
		BorrowRate:         entity.BorrowRate,
		SupplyRate:         entity.SupplyRate,
		Cash:               entity.Cash,
		CollateralFactor:   entity.CollateralFactor,
		ExchangeRate:       entity.ExchangeRate,
		ReserveFactor:      entity.ReserveFactor,
		Reserves:           entity.Reserves,
		TotalBorrows:       entity.TotalBorrows,
		TotalSupply:        entity.TotalSupply,
		UnderlyingAddress:  entity.UnderlyingAddress,
		UnderlyingName:     entity.UnderlyingName,
		UnderlyingSymbol:   entity.UnderlyingSymbol,
		UnderlyingPrice:    entity.UnderlyingPrice,
		UnderlyingDecimals: entity.UnderlyingDecimals,
		UnderlyingPriceUSD: entity.UnderlyingPriceUsd,
		BlockNumber:        entity.BlockNumber,
		BlockTimestamp:     entity.BlockTimestamp,
	}
}

func (r *marketRepo) listToBiz(list []*ent.Market) []*biz.Market {
	var markets []*biz.Market
	for _, bizMarket := range list {
		m := r.dataToBiz(bizMarket)
		markets = append(markets, m)
	}
	return markets
}
