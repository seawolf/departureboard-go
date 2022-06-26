// Code generated by entc, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/service"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/toc"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TOCCreate is the builder for creating a TOC entity.
type TOCCreate struct {
	config
	mutation *TOCMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TOCCreate) SetName(s string) *TOCCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tc *TOCCreate) SetNillableName(s *string) *TOCCreate {
	if s != nil {
		tc.SetName(*s)
	}
	return tc
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (tc *TOCCreate) AddServiceIDs(ids ...int) *TOCCreate {
	tc.mutation.AddServiceIDs(ids...)
	return tc
}

// AddServices adds the "services" edges to the Service entity.
func (tc *TOCCreate) AddServices(s ...*Service) *TOCCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddServiceIDs(ids...)
}

// Mutation returns the TOCMutation object of the builder.
func (tc *TOCCreate) Mutation() *TOCMutation {
	return tc.mutation
}

// Save creates the TOC in the database.
func (tc *TOCCreate) Save(ctx context.Context) (*TOC, error) {
	var (
		err  error
		node *TOC
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TOCMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TOCCreate) SaveX(ctx context.Context) *TOC {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TOCCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TOCCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TOCCreate) defaults() {
	if _, ok := tc.mutation.Name(); !ok {
		v := toc.DefaultName
		tc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TOCCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`entities: missing required field "TOC.name"`)}
	}
	return nil
}

func (tc *TOCCreate) sqlSave(ctx context.Context) (*TOC, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TOCCreate) createSpec() (*TOC, *sqlgraph.CreateSpec) {
	var (
		_node = &TOC{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: toc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: toc.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: toc.FieldName,
		})
		_node.Name = value
	}
	if nodes := tc.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   toc.ServicesTable,
			Columns: []string{toc.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TOCCreateBulk is the builder for creating many TOC entities in bulk.
type TOCCreateBulk struct {
	config
	builders []*TOCCreate
}

// Save creates the TOC entities in the database.
func (tcb *TOCCreateBulk) Save(ctx context.Context) ([]*TOC, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*TOC, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TOCMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TOCCreateBulk) SaveX(ctx context.Context) []*TOC {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TOCCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TOCCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
