package task

import (
	"compound/internal/util"
	"compound/pkg/contract/oracle"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"math/big"
)

// SyncOraclePrice 同步预言机价格

type OraclePriceTask struct {
	log *log.Helper
	at  *AccountTask
	job *gocron.Job
}

// NewOraclePriceTask new an account task.
func NewOraclePriceTask(at *AccountTask, logger log.Logger) *OraclePriceTask {
	return &OraclePriceTask{
		at:  at,
		log: log.NewHelper(log.With(logger, "module", "task/oracle_price")),
	}
}

func (t *OraclePriceTask) SetJob(job *gocron.Job) {
	t.job = job
}

func (t *OraclePriceTask) Handle() {
	marketMap, err := t.at.muc.GetMarketMap(context.Background())
	if err != nil {
		t.log.Errorf("获取 getMarketMap 失败：%+v", err)
		return
	}
	if len(marketMap) == 0 {
		t.log.Info("marketMap 是空的")
		return
	}

	oracleContract, err := oracle.NewOracle(t.at.task.priceOracleAddress, t.at.task.ethCli)

	var allUsers []string

	callOpts := bind.CallOpts{Pending: false}
	for _, market := range marketMap {
		price, err := oracleContract.GetUnderlyingPrice(&callOpts, common.HexToAddress(market.Address))
		if err != nil {
			t.log.Errorf("获取预言机价格 GetUnderlyingPrice 失败：%+v", err)
			return
		}
		mantissa := decimal.NewFromBigInt(big.NewInt(10), int32(18+18-market.UnderlyingDecimals))
		priceBD := decimal.NewFromBigInt(price, 1)
		priceInUSD := priceBD.Div(mantissa)
		underlyingPrice, _ := decimal.NewFromString(market.UnderlyingPriceUSD)
		// 新获取的价格 和 资金池存的价格不一样 需要更新
		if !priceInUSD.Equal(underlyingPrice) {
			market.UnderlyingPrice = priceInUSD.String()
			_, err := t.at.muc.UpdateMarketOraclePrice(context.Background(), market.ID, market.UnderlyingPrice)
			if err != nil {
				t.log.Errorf("market.UpdateMarketOraclePrice(id:%,underlyingPrice%s) 价格更新失败, err:%+v",
					market.ID, market.UnderlyingPrice, err)
				return
			}
		}
		// 获取预言机价格有所改变的用户
		users, err := t.GetUsersOraclePriceChanged(market.Address, market.UnderlyingPriceUSD)
		if err != nil {
			t.log.Errorf("GetUsersOraclePriceChanged 请求失败, err %+v", err)
			break
		}
		allUsers = append(allUsers, users...)
	}

	// 用户去重
	allUsers = util.RemoveDuplicateElement(allUsers)

	//
	for _, addr := range allUsers {
		t.at.InvokePool(marketMap, addr)
	}
}

// GetUsersOraclePriceChanged 慢慢获取， 一次获取太多可能会造成拥赛
func (t *OraclePriceTask) GetUsersOraclePriceChanged(address string, underlyingPriceUsd string) ([]string, error) {
	page := 1
	pageSize := 1000
	var users []string
	flag := true
	for flag {
		list, err := t.at.uc.GetUsersOraclePriceChanged(context.Background(), address, underlyingPriceUsd, int64(page),
			int64(pageSize))
		if err != nil {
			return nil, err
		}
		users = append(users, list...)
		if len(list) < pageSize {
			flag = false
		} else {
			page++
		}
	}
	return users, nil
}
