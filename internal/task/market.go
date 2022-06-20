package task

import (
	"compound/internal/biz"
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hasura/go-graphql-client"
)

type Market struct {
	Id                       graphql.String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BorrowRate               graphql.String `protobuf:"bytes,2,opt,name=borrow_rate,json=borrowRate,proto3" json:"borrow_rate,omitempty"`
	Cash                     graphql.String `protobuf:"bytes,3,opt,name=cash,proto3" json:"cash,omitempty"`
	CollateralFactor         graphql.String `protobuf:"bytes,4,opt,name=collateral_factor,json=collateralFactor,proto3" json:"collateral_factor,omitempty"`
	ExchangeRate             graphql.String `protobuf:"bytes,5,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"`
	InterestRateModelAddress graphql.String `protobuf:"bytes,6,opt,name=interest_rate_model_address,json=interestRateModelAddress,proto3" json:"interest_rate_model_address,omitempty"`
	Name                     graphql.String `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Reserves                 graphql.String `protobuf:"bytes,8,opt,name=reserves,proto3" json:"reserves,omitempty"`
	SupplyRate               graphql.String `protobuf:"bytes,9,opt,name=supply_rate,json=supplyRate,proto3" json:"supply_rate,omitempty"`
	Symbol                   graphql.String `protobuf:"bytes,10,opt,name=symbol,proto3" json:"symbol,omitempty"`
	TotalBorrows             graphql.String `protobuf:"bytes,11,opt,name=total_borrows,json=totalBorrows,proto3" json:"total_borrows,omitempty"`
	TotalSupply              graphql.String `protobuf:"bytes,12,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply,omitempty"`
	UnderlyingAddress        graphql.String `protobuf:"bytes,13,opt,name=underlying_address,json=underlyingAddress,proto3" json:"underlying_address,omitempty"`
	UnderlyingName           graphql.String `protobuf:"bytes,14,opt,name=underlying_name,json=underlyingName,proto3" json:"underlying_name,omitempty"`
	UnderlyingPrice          graphql.String `protobuf:"bytes,15,opt,name=underlying_price,json=underlyingPrice,proto3" json:"underlying_price,omitempty"`
	UnderlyingSymbol         graphql.String `protobuf:"bytes,16,opt,name=underlying_symbol,json=underlyingSymbol,proto3" json:"underlying_symbol,omitempty"`
	AccrualBlockNumber       graphql.Int    `protobuf:"varint,17,opt,name=accrual_block_number,json=accrualBlockNumber,proto3" json:"accrual_block_number,omitempty"`
	BlockTimestamp           graphql.Int    `protobuf:"varint,18,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`
	BorrowIndex              graphql.String `protobuf:"bytes,19,opt,name=borrow_index,json=borrowIndex,proto3" json:"borrow_index,omitempty"`
	ReserveFactor            graphql.String `protobuf:"bytes,20,opt,name=reserve_factor,json=reserveFactor,proto3" json:"reserve_factor,omitempty"`
	UnderlyingPriceUSD       graphql.String `graphql:"underlyingPriceUSD" json:"underlying_price_usd"`
	UnderlyingDecimals       graphql.Int    `protobuf:"varint,22,opt,name=underlying_decimals,json=underlyingDecimals,proto3" json:"underlying_decimals,omitempty"`
}

type MarketTask struct {
	uc   *biz.MarketUseCase
	log  *log.Helper
	task *Task
	job  *gocron.Job
}

// NewMarketTask new a market task.
func NewMarketTask(t *Task, uc *biz.MarketUseCase, logger log.Logger) *MarketTask {
	return &MarketTask{
		uc:   uc,
		log:  log.NewHelper(log.With(logger, "module", "task/market")),
		task: t,
	}
}

func (t *MarketTask) SetJob(job *gocron.Job) {
	t.job = job
}

// Handle 同步 market
func (t *MarketTask) Handle() {
	//if t.job.RunCount() > 1 {
	//	t.log.Infof("MarketTask... 还有任务在执行, RunCount:%d", t.job.RunCount())
	//}

	var query struct {
		Markets []Market `graphql:"markets(first: 10)"`
	}
	err := t.task.graphqlCli.Query(context.Background(), &query, nil)
	if err != nil {
		t.log.Error(err)
		return
	}
	for _, market := range query.Markets {
		_, err := t.uc.SaveMarket(context.Background(), &biz.Market{
			Address:            string(market.Id),
			Name:               string(market.Name),
			Symbol:             string(market.Symbol),
			BorrowIndex:        string(market.BorrowIndex),
			BorrowRate:         string(market.BorrowRate),
			SupplyRate:         string(market.SupplyRate),
			Cash:               string(market.Cash),
			CollateralFactor:   string(market.CollateralFactor),
			ExchangeRate:       string(market.ExchangeRate),
			ReserveFactor:      string(market.ReserveFactor),
			Reserves:           string(market.Reserves),
			TotalBorrows:       string(market.TotalBorrows),
			TotalSupply:        string(market.TotalSupply),
			UnderlyingAddress:  string(market.UnderlyingAddress),
			UnderlyingName:     string(market.UnderlyingName),
			UnderlyingSymbol:   string(market.UnderlyingSymbol),
			UnderlyingPrice:    string(market.UnderlyingPrice),
			UnderlyingDecimals: uint32(market.UnderlyingDecimals),
			UnderlyingPriceUSD: string(market.UnderlyingPriceUSD),
			BlockNumber:        uint64(market.AccrualBlockNumber),
			BlockTimestamp:     uint32(market.BlockTimestamp),
		})
		if err != nil {
			t.log.Error(err)
			break
		}
	}
}
