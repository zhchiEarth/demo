package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Market struct {
	ID                 int
	Address            string `json:"address"`
	Name               string
	Symbol             string
	BorrowIndex        string
	BorrowRate         string
	SupplyRate         string
	Cash               string
	CollateralFactor   string
	ExchangeRate       string
	ReserveFactor      string
	Reserves           string
	TotalBorrows       string
	TotalSupply        string
	UnderlyingAddress  string
	UnderlyingName     string
	UnderlyingSymbol   string
	UnderlyingPrice    string
	UnderlyingDecimals uint32
	UnderlyingPriceUSD string
	BlockNumber        uint64
	BlockTimestamp     uint32
}

type MarketReply struct {
	ID int
}

type MarketRepo interface {
	ExistMarket(ctx context.Context, address string) (bool, error)
	SaveMarket(ctx context.Context, c *Market) (*MarketReply, error)
	AllMarket(ctx context.Context) ([]*Market, error)
	UpdateMarketOraclePrice(ctx context.Context, id int, underlyingPriceUsd string) (*MarketReply, error)
}

type MarketUseCase struct {
	repo MarketRepo
	log  *log.Helper
}

func NewMarketUseCase(repo MarketRepo, logger log.Logger) *MarketUseCase {
	return &MarketUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/market")),
	}
}

func (uc *MarketUseCase) SaveMarket(ctx context.Context, m *Market) (*MarketReply, error) {
	return uc.repo.SaveMarket(ctx, m)
}

func (uc *MarketUseCase) All(ctx context.Context) ([]*Market, error) {
	return uc.repo.AllMarket(ctx)
}

func (uc *MarketUseCase) UpdateMarketOraclePrice(ctx context.Context, id int, underlyingPriceUsd string) (*MarketReply, error) {
	return uc.repo.UpdateMarketOraclePrice(ctx, id, underlyingPriceUsd)
}

func (uc *MarketUseCase) GetMarketMap(ctx context.Context) (map[string]*Market, error) {
	markets, err := uc.repo.AllMarket(ctx)
	if err != nil {
		return nil, err
	}
	marketMap := make(map[string]*Market, len(markets))
	for _, market := range markets {
		_, ok := marketMap[market.Address]
		if !ok {
			marketMap[market.Address] = market
		}
	}
	return marketMap, nil
}
