package task

import (
	"compound/internal/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/hasura/go-graphql-client"
)

// ProviderSet is task providers.
var ProviderSet = wire.NewSet(NewTask, NewMarketTask, NewAccountTask, NewOraclePriceTask)

type Task struct {
	graphqlCli         *graphql.Client
	log                *log.Helper
	ethCli             *ethclient.Client
	priceOracleAddress common.Address
	comptrollerAddress common.Address
	cEtherAddress      common.Address
}

// NewTask .
func NewTask(conf *conf.Task, logger log.Logger) (*Task, func(), error) {
	logData := log.NewHelper(log.With(logger, "module", "task"))

	cli := graphql.NewClient(conf.Graph.Source, nil)

	client, err := ethclient.Dial(conf.Contract.Source)
	if err != nil {
		logData.Fatalf("failed opening connection to ethclient: %v", err)
	}

	t := &Task{
		graphqlCli:         cli,
		log:                logData,
		ethCli:             client,
		priceOracleAddress: common.HexToAddress(conf.Contract.PriceOracle),
		comptrollerAddress: common.HexToAddress(conf.Contract.Comptroller),
		cEtherAddress:      common.HexToAddress(conf.Contract.Cether),
	}
	cleanup := func() {
		client.Close()
		logData.Info("closing the ethclient resources")
	}
	return t, cleanup, nil
}
