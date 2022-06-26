// Code generated by entc, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/callingpoint"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CallingPointCreate is the builder for creating a CallingPoint entity.
type CallingPointCreate struct {
	config
	mutation *CallingPointMutation
	hooks    []Hook
}

// SetArrivalTime sets the "arrival_time" field.
func (cpc *CallingPointCreate) SetArrivalTime(t time.Time) *CallingPointCreate {
	cpc.mutation.SetArrivalTime(t)
	return cpc
}

// SetNillableArrivalTime sets the "arrival_time" field if the given value is not nil.
func (cpc *CallingPointCreate) SetNillableArrivalTime(t *time.Time) *CallingPointCreate {
	if t != nil {
		cpc.SetArrivalTime(*t)
	}
	return cpc
}

// SetDepartureTime sets the "departure_time" field.
func (cpc *CallingPointCreate) SetDepartureTime(t time.Time) *CallingPointCreate {
	cpc.mutation.SetDepartureTime(t)
	return cpc
}

// SetNillableDepartureTime sets the "departure_time" field if the given value is not nil.
func (cpc *CallingPointCreate) SetNillableDepartureTime(t *time.Time) *CallingPointCreate {
	if t != nil {
		cpc.SetDepartureTime(*t)
	}
	return cpc
}

// Mutation returns the CallingPointMutation object of the builder.
func (cpc *CallingPointCreate) Mutation() *CallingPointMutation {
	return cpc.mutation
}

// Save creates the CallingPoint in the database.
func (cpc *CallingPointCreate) Save(ctx context.Context) (*CallingPoint, error) {
	var (
		err  error
		node *CallingPoint
	)
	cpc.defaults()
	if len(cpc.hooks) == 0 {
		if err = cpc.check(); err != nil {
			return nil, err
		}
		node, err = cpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CallingPointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpc.check(); err != nil {
				return nil, err
			}
			cpc.mutation = mutation
			if node, err = cpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cpc.hooks) - 1; i >= 0; i-- {
			if cpc.hooks[i] == nil {
				return nil, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = cpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cpc *CallingPointCreate) SaveX(ctx context.Context) *CallingPoint {
	v, err := cpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpc *CallingPointCreate) Exec(ctx context.Context) error {
	_, err := cpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpc *CallingPointCreate) ExecX(ctx context.Context) {
	if err := cpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpc *CallingPointCreate) defaults() {
	if _, ok := cpc.mutation.ArrivalTime(); !ok {
		v := callingpoint.DefaultArrivalTime
		cpc.mutation.SetArrivalTime(v)
	}
	if _, ok := cpc.mutation.DepartureTime(); !ok {
		v := callingpoint.DefaultDepartureTime
		cpc.mutation.SetDepartureTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpc *CallingPointCreate) check() error {
	if _, ok := cpc.mutation.ArrivalTime(); !ok {
		return &ValidationError{Name: "arrival_time", err: errors.New(`entities: missing required field "CallingPoint.arrival_time"`)}
	}
	if _, ok := cpc.mutation.DepartureTime(); !ok {
		return &ValidationError{Name: "departure_time", err: errors.New(`entities: missing required field "CallingPoint.departure_time"`)}
	}
	return nil
}

func (cpc *CallingPointCreate) sqlSave(ctx context.Context) (*CallingPoint, error) {
	_node, _spec := cpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cpc *CallingPointCreate) createSpec() (*CallingPoint, *sqlgraph.CreateSpec) {
	var (
		_node = &CallingPoint{config: cpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: callingpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: callingpoint.FieldID,
			},
		}
	)
	if value, ok := cpc.mutation.ArrivalTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: callingpoint.FieldArrivalTime,
		})
		_node.ArrivalTime = value
	}
	if value, ok := cpc.mutation.DepartureTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: callingpoint.FieldDepartureTime,
		})
		_node.DepartureTime = value
	}
	return _node, _spec
}

// CallingPointCreateBulk is the builder for creating many CallingPoint entities in bulk.
type CallingPointCreateBulk struct {
	config
	builders []*CallingPointCreate
}

// Save creates the CallingPoint entities in the database.
func (cpcb *CallingPointCreateBulk) Save(ctx context.Context) ([]*CallingPoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cpcb.builders))
	nodes := make([]*CallingPoint, len(cpcb.builders))
	mutators := make([]Mutator, len(cpcb.builders))
	for i := range cpcb.builders {
		func(i int, root context.Context) {
			builder := cpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CallingPointMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cpcb *CallingPointCreateBulk) SaveX(ctx context.Context) []*CallingPoint {
	v, err := cpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpcb *CallingPointCreateBulk) Exec(ctx context.Context) error {
	_, err := cpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpcb *CallingPointCreateBulk) ExecX(ctx context.Context) {
	if err := cpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
