package data

import (
	"compound/internal/biz"
	"compound/internal/conf"
	"compound/internal/data/ent"
	"compound/internal/data/ent/migrate"
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTransaction, NewMarketRepo, NewAccountRepo,
	NewAccountCTokenRepo, NewPreferenceRepo)

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
}

type contextTxKey struct{}

func (d *Data) ExecTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := d.db.Tx(ctx)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, contextTxKey{}, tx)
	if err := f(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (d *Data) Account(ctx context.Context) *ent.AccountClient {
	tx, ok := ctx.Value(contextTxKey{}).(*ent.Tx)
	if ok {
		return tx.Account
	}
	return d.db.Account
}

func (d *Data) AccountCToken(ctx context.Context) *ent.AccountCTokenClient {
	tx, ok := ctx.Value(contextTxKey{}).(*ent.Tx)
	if ok {
		return tx.AccountCToken
	}
	return d.db.AccountCToken
}

func (d *Data) Market(ctx context.Context) *ent.MarketClient {
	tx, ok := ctx.Value(contextTxKey{}).(*ent.Tx)
	if ok {
		return tx.Market
	}
	return d.db.Market
}

// "github.com/go-redis/redis/v8"
//func NewRedis(conf *conf.Data, logger log.Logger) *redis.Client {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:         conf.Redis.Addr,
//		Password:     conf.Redis.Password,
//		DB:           int(conf.Redis.Db),
//		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
//		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
//		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
//	})
//	rdb.AddHook(redisotel.TracingHook{})
//	return rdb
//}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	logData := log.NewHelper(log.With(logger, "module", "data"))

	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	client := ent.NewClient(ent.Driver(drv))
	if err != nil {
		logData.Fatalf("failed opening connection to db: %v", err)
	}
	//	// Run the auto migration tool.
	err = client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false), //禁用外健
	)
	if err != nil {
		logData.Fatalf("failed creating schema resources: %v", err)
	}

	d := &Data{
		db:  client,
		log: logData,
	}
	cleanup := func() {
		//db
		logData.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}
