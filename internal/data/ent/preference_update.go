// Code generated by entc, DO NOT EDIT.

package ent

import (
	"compound/internal/data/ent/predicate"
	"compound/internal/data/ent/preference"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PreferenceUpdate is the builder for updating Preference entities.
type PreferenceUpdate struct {
	config
	hooks    []Hook
	mutation *PreferenceMutation
}

// Where appends a list predicates to the PreferenceUpdate builder.
func (pu *PreferenceUpdate) Where(ps ...predicate.Preference) *PreferenceUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *PreferenceUpdate) SetUpdateTime(t time.Time) *PreferenceUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// SetValue sets the "value" field.
func (pu *PreferenceUpdate) SetValue(s string) *PreferenceUpdate {
	pu.mutation.SetValue(s)
	return pu
}

// Mutation returns the PreferenceMutation object of the builder.
func (pu *PreferenceUpdate) Mutation() *PreferenceMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PreferenceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PreferenceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PreferenceUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PreferenceUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PreferenceUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PreferenceUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := preference.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

func (pu *PreferenceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   preference.Table,
			Columns: preference.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: preference.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: preference.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: preference.FieldValue,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{preference.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PreferenceUpdateOne is the builder for updating a single Preference entity.
type PreferenceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PreferenceMutation
}

// SetUpdateTime sets the "update_time" field.
func (puo *PreferenceUpdateOne) SetUpdateTime(t time.Time) *PreferenceUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// SetValue sets the "value" field.
func (puo *PreferenceUpdateOne) SetValue(s string) *PreferenceUpdateOne {
	puo.mutation.SetValue(s)
	return puo
}

// Mutation returns the PreferenceMutation object of the builder.
func (puo *PreferenceUpdateOne) Mutation() *PreferenceMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PreferenceUpdateOne) Select(field string, fields ...string) *PreferenceUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Preference entity.
func (puo *PreferenceUpdateOne) Save(ctx context.Context) (*Preference, error) {
	var (
		err  error
		node *Preference
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PreferenceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PreferenceUpdateOne) SaveX(ctx context.Context) *Preference {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PreferenceUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PreferenceUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PreferenceUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := preference.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

func (puo *PreferenceUpdateOne) sqlSave(ctx context.Context) (_node *Preference, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   preference.Table,
			Columns: preference.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: preference.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Preference.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, preference.FieldID)
		for _, f := range fields {
			if !preference.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != preference.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: preference.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: preference.FieldValue,
		})
	}
	_node = &Preference{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{preference.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
