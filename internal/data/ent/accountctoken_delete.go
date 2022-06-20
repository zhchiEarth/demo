// Code generated by entc, DO NOT EDIT.

package ent

import (
	"compound/internal/data/ent/accountctoken"
	"compound/internal/data/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountCTokenDelete is the builder for deleting a AccountCToken entity.
type AccountCTokenDelete struct {
	config
	hooks    []Hook
	mutation *AccountCTokenMutation
}

// Where appends a list predicates to the AccountCTokenDelete builder.
func (acd *AccountCTokenDelete) Where(ps ...predicate.AccountCToken) *AccountCTokenDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *AccountCTokenDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(acd.hooks) == 0 {
		affected, err = acd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountCTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acd.mutation = mutation
			affected, err = acd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acd.hooks) - 1; i >= 0; i-- {
			if acd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *AccountCTokenDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *AccountCTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: accountctoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountctoken.FieldID,
			},
		},
	}
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
}

// AccountCTokenDeleteOne is the builder for deleting a single AccountCToken entity.
type AccountCTokenDeleteOne struct {
	acd *AccountCTokenDelete
}

// Exec executes the deletion query.
func (acdo *AccountCTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountctoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *AccountCTokenDeleteOne) ExecX(ctx context.Context) {
	acdo.acd.ExecX(ctx)
}