// Code generated by entc, DO NOT EDIT.

package ent

import (
	"compound/internal/data/ent/market"
	"compound/internal/data/ent/predicate"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MarketQuery is the builder for querying Market entities.
type MarketQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Market
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MarketQuery builder.
func (mq *MarketQuery) Where(ps ...predicate.Market) *MarketQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MarketQuery) Limit(limit int) *MarketQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MarketQuery) Offset(offset int) *MarketQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MarketQuery) Unique(unique bool) *MarketQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *MarketQuery) Order(o ...OrderFunc) *MarketQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// First returns the first Market entity from the query.
// Returns a *NotFoundError when no Market was found.
func (mq *MarketQuery) First(ctx context.Context) (*Market, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{market.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MarketQuery) FirstX(ctx context.Context) *Market {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Market ID from the query.
// Returns a *NotFoundError when no Market ID was found.
func (mq *MarketQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{market.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MarketQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Market entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Market entity is found.
// Returns a *NotFoundError when no Market entities are found.
func (mq *MarketQuery) Only(ctx context.Context) (*Market, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{market.Label}
	default:
		return nil, &NotSingularError{market.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MarketQuery) OnlyX(ctx context.Context) *Market {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Market ID in the query.
// Returns a *NotSingularError when more than one Market ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MarketQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = &NotSingularError{market.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MarketQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Markets.
func (mq *MarketQuery) All(ctx context.Context) ([]*Market, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MarketQuery) AllX(ctx context.Context) []*Market {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Market IDs.
func (mq *MarketQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(market.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MarketQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MarketQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MarketQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MarketQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MarketQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MarketQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MarketQuery) Clone() *MarketQuery {
	if mq == nil {
		return nil
	}
	return &MarketQuery{
		config:     mq.config,
		limit:      mq.limit,
		offset:     mq.offset,
		order:      append([]OrderFunc{}, mq.order...),
		predicates: append([]predicate.Market{}, mq.predicates...),
		// clone intermediate query.
		sql:    mq.sql.Clone(),
		path:   mq.path,
		unique: mq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Market.Query().
//		GroupBy(market.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mq *MarketQuery) GroupBy(field string, fields ...string) *MarketGroupBy {
	group := &MarketGroupBy{config: mq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Market.Query().
//		Select(market.FieldCreateTime).
//		Scan(ctx, &v)
//
func (mq *MarketQuery) Select(fields ...string) *MarketSelect {
	mq.fields = append(mq.fields, fields...)
	return &MarketSelect{MarketQuery: mq}
}

func (mq *MarketQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !market.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MarketQuery) sqlAll(ctx context.Context) ([]*Market, error) {
	var (
		nodes = []*Market{}
		_spec = mq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Market{config: mq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mq *MarketQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.fields
	if len(mq.fields) > 0 {
		_spec.Unique = mq.unique != nil && *mq.unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MarketQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mq *MarketQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   market.Table,
			Columns: market.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: market.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, market.FieldID)
		for i := range fields {
			if fields[i] != market.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MarketQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(market.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = market.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.unique != nil && *mq.unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MarketGroupBy is the group-by builder for Market entities.
type MarketGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MarketGroupBy) Aggregate(fns ...AggregateFunc) *MarketGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *MarketGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgb *MarketGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MarketGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MarketGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgb *MarketGroupBy) StringsX(ctx context.Context) []string {
	v, err := mgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MarketGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = fmt.Errorf("ent: MarketGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mgb *MarketGroupBy) StringX(ctx context.Context) string {
	v, err := mgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MarketGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MarketGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgb *MarketGroupBy) IntsX(ctx context.Context) []int {
	v, err := mgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MarketGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = fmt.Errorf("ent: MarketGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mgb *MarketGroupBy) IntX(ctx context.Context) int {
	v, err := mgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MarketGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MarketGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgb *MarketGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MarketGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = fmt.Errorf("ent: MarketGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mgb *MarketGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MarketGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MarketGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgb *MarketGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MarketGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = fmt.Errorf("ent: MarketGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mgb *MarketGroupBy) BoolX(ctx context.Context) bool {
	v, err := mgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgb *MarketGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mgb.fields {
		if !market.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MarketGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql.Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
		for _, f := range mgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgb.fields...)...)
}

// MarketSelect is the builder for selecting fields of Market entities.
type MarketSelect struct {
	*MarketQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MarketSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.MarketQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ms *MarketSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ms *MarketSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MarketSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ms *MarketSelect) StringsX(ctx context.Context) []string {
	v, err := ms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ms *MarketSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = fmt.Errorf("ent: MarketSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ms *MarketSelect) StringX(ctx context.Context) string {
	v, err := ms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ms *MarketSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MarketSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ms *MarketSelect) IntsX(ctx context.Context) []int {
	v, err := ms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ms *MarketSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = fmt.Errorf("ent: MarketSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ms *MarketSelect) IntX(ctx context.Context) int {
	v, err := ms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ms *MarketSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MarketSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ms *MarketSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ms *MarketSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = fmt.Errorf("ent: MarketSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ms *MarketSelect) Float64X(ctx context.Context) float64 {
	v, err := ms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ms *MarketSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MarketSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ms *MarketSelect) BoolsX(ctx context.Context) []bool {
	v, err := ms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ms *MarketSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{market.Label}
	default:
		err = fmt.Errorf("ent: MarketSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ms *MarketSelect) BoolX(ctx context.Context) bool {
	v, err := ms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ms *MarketSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sql.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
