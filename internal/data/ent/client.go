// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"compound/internal/data/ent/migrate"

	"compound/internal/data/ent/account"
	"compound/internal/data/ent/accountctoken"
	"compound/internal/data/ent/market"
	"compound/internal/data/ent/preference"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// AccountCToken is the client for interacting with the AccountCToken builders.
	AccountCToken *AccountCTokenClient
	// Market is the client for interacting with the Market builders.
	Market *MarketClient
	// Preference is the client for interacting with the Preference builders.
	Preference *PreferenceClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.AccountCToken = NewAccountCTokenClient(c.config)
	c.Market = NewMarketClient(c.config)
	c.Preference = NewPreferenceClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Account:       NewAccountClient(cfg),
		AccountCToken: NewAccountCTokenClient(cfg),
		Market:        NewMarketClient(cfg),
		Preference:    NewPreferenceClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Account:       NewAccountClient(cfg),
		AccountCToken: NewAccountCTokenClient(cfg),
		Market:        NewMarketClient(cfg),
		Preference:    NewPreferenceClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.AccountCToken.Use(hooks...)
	c.Market.Use(hooks...)
	c.Preference.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a create builder for Account.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id int) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AccountClient) DeleteOneID(id int) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id int) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id int) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTokens queries the tokens edge of a Account.
func (c *AccountClient) QueryTokens(a *Account) *AccountCTokenQuery {
	query := &AccountCTokenQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(accountctoken.Table, accountctoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.TokensTable, account.TokensColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// AccountCTokenClient is a client for the AccountCToken schema.
type AccountCTokenClient struct {
	config
}

// NewAccountCTokenClient returns a client for the AccountCToken from the given config.
func NewAccountCTokenClient(c config) *AccountCTokenClient {
	return &AccountCTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `accountctoken.Hooks(f(g(h())))`.
func (c *AccountCTokenClient) Use(hooks ...Hook) {
	c.hooks.AccountCToken = append(c.hooks.AccountCToken, hooks...)
}

// Create returns a create builder for AccountCToken.
func (c *AccountCTokenClient) Create() *AccountCTokenCreate {
	mutation := newAccountCTokenMutation(c.config, OpCreate)
	return &AccountCTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AccountCToken entities.
func (c *AccountCTokenClient) CreateBulk(builders ...*AccountCTokenCreate) *AccountCTokenCreateBulk {
	return &AccountCTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AccountCToken.
func (c *AccountCTokenClient) Update() *AccountCTokenUpdate {
	mutation := newAccountCTokenMutation(c.config, OpUpdate)
	return &AccountCTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountCTokenClient) UpdateOne(ac *AccountCToken) *AccountCTokenUpdateOne {
	mutation := newAccountCTokenMutation(c.config, OpUpdateOne, withAccountCToken(ac))
	return &AccountCTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountCTokenClient) UpdateOneID(id int) *AccountCTokenUpdateOne {
	mutation := newAccountCTokenMutation(c.config, OpUpdateOne, withAccountCTokenID(id))
	return &AccountCTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AccountCToken.
func (c *AccountCTokenClient) Delete() *AccountCTokenDelete {
	mutation := newAccountCTokenMutation(c.config, OpDelete)
	return &AccountCTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AccountCTokenClient) DeleteOne(ac *AccountCToken) *AccountCTokenDeleteOne {
	return c.DeleteOneID(ac.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AccountCTokenClient) DeleteOneID(id int) *AccountCTokenDeleteOne {
	builder := c.Delete().Where(accountctoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountCTokenDeleteOne{builder}
}

// Query returns a query builder for AccountCToken.
func (c *AccountCTokenClient) Query() *AccountCTokenQuery {
	return &AccountCTokenQuery{
		config: c.config,
	}
}

// Get returns a AccountCToken entity by its id.
func (c *AccountCTokenClient) Get(ctx context.Context, id int) (*AccountCToken, error) {
	return c.Query().Where(accountctoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountCTokenClient) GetX(ctx context.Context, id int) *AccountCToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a AccountCToken.
func (c *AccountCTokenClient) QueryAccount(ac *AccountCToken) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ac.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(accountctoken.Table, accountctoken.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accountctoken.AccountTable, accountctoken.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(ac.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountCTokenClient) Hooks() []Hook {
	return c.hooks.AccountCToken
}

// MarketClient is a client for the Market schema.
type MarketClient struct {
	config
}

// NewMarketClient returns a client for the Market from the given config.
func NewMarketClient(c config) *MarketClient {
	return &MarketClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `market.Hooks(f(g(h())))`.
func (c *MarketClient) Use(hooks ...Hook) {
	c.hooks.Market = append(c.hooks.Market, hooks...)
}

// Create returns a create builder for Market.
func (c *MarketClient) Create() *MarketCreate {
	mutation := newMarketMutation(c.config, OpCreate)
	return &MarketCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Market entities.
func (c *MarketClient) CreateBulk(builders ...*MarketCreate) *MarketCreateBulk {
	return &MarketCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Market.
func (c *MarketClient) Update() *MarketUpdate {
	mutation := newMarketMutation(c.config, OpUpdate)
	return &MarketUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MarketClient) UpdateOne(m *Market) *MarketUpdateOne {
	mutation := newMarketMutation(c.config, OpUpdateOne, withMarket(m))
	return &MarketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MarketClient) UpdateOneID(id int) *MarketUpdateOne {
	mutation := newMarketMutation(c.config, OpUpdateOne, withMarketID(id))
	return &MarketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Market.
func (c *MarketClient) Delete() *MarketDelete {
	mutation := newMarketMutation(c.config, OpDelete)
	return &MarketDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MarketClient) DeleteOne(m *Market) *MarketDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MarketClient) DeleteOneID(id int) *MarketDeleteOne {
	builder := c.Delete().Where(market.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MarketDeleteOne{builder}
}

// Query returns a query builder for Market.
func (c *MarketClient) Query() *MarketQuery {
	return &MarketQuery{
		config: c.config,
	}
}

// Get returns a Market entity by its id.
func (c *MarketClient) Get(ctx context.Context, id int) (*Market, error) {
	return c.Query().Where(market.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MarketClient) GetX(ctx context.Context, id int) *Market {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MarketClient) Hooks() []Hook {
	return c.hooks.Market
}

// PreferenceClient is a client for the Preference schema.
type PreferenceClient struct {
	config
}

// NewPreferenceClient returns a client for the Preference from the given config.
func NewPreferenceClient(c config) *PreferenceClient {
	return &PreferenceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `preference.Hooks(f(g(h())))`.
func (c *PreferenceClient) Use(hooks ...Hook) {
	c.hooks.Preference = append(c.hooks.Preference, hooks...)
}

// Create returns a create builder for Preference.
func (c *PreferenceClient) Create() *PreferenceCreate {
	mutation := newPreferenceMutation(c.config, OpCreate)
	return &PreferenceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Preference entities.
func (c *PreferenceClient) CreateBulk(builders ...*PreferenceCreate) *PreferenceCreateBulk {
	return &PreferenceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Preference.
func (c *PreferenceClient) Update() *PreferenceUpdate {
	mutation := newPreferenceMutation(c.config, OpUpdate)
	return &PreferenceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PreferenceClient) UpdateOne(pr *Preference) *PreferenceUpdateOne {
	mutation := newPreferenceMutation(c.config, OpUpdateOne, withPreference(pr))
	return &PreferenceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PreferenceClient) UpdateOneID(id int) *PreferenceUpdateOne {
	mutation := newPreferenceMutation(c.config, OpUpdateOne, withPreferenceID(id))
	return &PreferenceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Preference.
func (c *PreferenceClient) Delete() *PreferenceDelete {
	mutation := newPreferenceMutation(c.config, OpDelete)
	return &PreferenceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PreferenceClient) DeleteOne(pr *Preference) *PreferenceDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PreferenceClient) DeleteOneID(id int) *PreferenceDeleteOne {
	builder := c.Delete().Where(preference.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PreferenceDeleteOne{builder}
}

// Query returns a query builder for Preference.
func (c *PreferenceClient) Query() *PreferenceQuery {
	return &PreferenceQuery{
		config: c.config,
	}
}

// Get returns a Preference entity by its id.
func (c *PreferenceClient) Get(ctx context.Context, id int) (*Preference, error) {
	return c.Query().Where(preference.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PreferenceClient) GetX(ctx context.Context, id int) *Preference {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PreferenceClient) Hooks() []Hook {
	return c.hooks.Preference
}