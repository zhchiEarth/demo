// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"compound/internal/biz"
	"compound/internal/conf"
	"compound/internal/data"
	"compound/internal/server"
	"compound/internal/service"
	"compound/internal/task"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, confTask *conf.Task, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	accountRepo := data.NewAccountRepo(dataData, logger)
	accountCTokenRepo := data.NewAccountCTokenRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	accountUseCase := biz.NewAccountUseCase(accountRepo, accountCTokenRepo, transaction, logger)
	accountService := service.NewAccountService(accountUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, accountService, logger)
	taskTask, cleanup2, err := task.NewTask(confTask, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	marketRepo := data.NewMarketRepo(dataData, logger)
	marketUseCase := biz.NewMarketUseCase(marketRepo, logger)
	marketTask := task.NewMarketTask(taskTask, marketUseCase, logger)
	preferenceRepo := data.NewPreferenceRepo(dataData, logger)
	preferenceUseCase := biz.NewPreferenceUseCase(preferenceRepo, logger)
	accountTask, cleanup3 := task.NewAccountTask(taskTask, accountUseCase, preferenceUseCase, marketUseCase, logger)
	oraclePriceTask := task.NewOraclePriceTask(accountTask, logger)
	cronServer := server.NewCronServer(confServer, logger, marketTask, accountTask, oraclePriceTask)
	app := newApp(logger, httpServer, cronServer)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}