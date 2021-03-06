// Code generated by entc, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/day"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/predicate"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/service"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DayUpdate is the builder for updating Day entities.
type DayUpdate struct {
	config
	hooks    []Hook
	mutation *DayMutation
}

// Where appends a list predicates to the DayUpdate builder.
func (du *DayUpdate) Where(ps ...predicate.Day) *DayUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetDate sets the "date" field.
func (du *DayUpdate) SetDate(t time.Time) *DayUpdate {
	du.mutation.SetDate(t)
	return du
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (du *DayUpdate) SetNillableDate(t *time.Time) *DayUpdate {
	if t != nil {
		du.SetDate(*t)
	}
	return du
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (du *DayUpdate) AddServiceIDs(ids ...int) *DayUpdate {
	du.mutation.AddServiceIDs(ids...)
	return du
}

// AddServices adds the "services" edges to the Service entity.
func (du *DayUpdate) AddServices(s ...*Service) *DayUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.AddServiceIDs(ids...)
}

// Mutation returns the DayMutation object of the builder.
func (du *DayUpdate) Mutation() *DayMutation {
	return du.mutation
}

// ClearServices clears all "services" edges to the Service entity.
func (du *DayUpdate) ClearServices() *DayUpdate {
	du.mutation.ClearServices()
	return du
}

// RemoveServiceIDs removes the "services" edge to Service entities by IDs.
func (du *DayUpdate) RemoveServiceIDs(ids ...int) *DayUpdate {
	du.mutation.RemoveServiceIDs(ids...)
	return du
}

// RemoveServices removes "services" edges to Service entities.
func (du *DayUpdate) RemoveServices(s ...*Service) *DayUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.RemoveServiceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DayUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DayUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DayUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DayUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DayUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   day.Table,
			Columns: day.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: day.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: day.FieldDate,
		})
	}
	if du.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   day.ServicesTable,
			Columns: []string{day.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedServicesIDs(); len(nodes) > 0 && !du.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   day.ServicesTable,
			Columns: []string{day.ServicesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   day.ServicesTable,
			Columns: []string{day.ServicesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{day.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DayUpdateOne is the builder for updating a single Day entity.
type DayUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DayMutation
}

// SetDate sets the "date" field.
func (duo *DayUpdateOne) SetDate(t time.Time) *DayUpdateOne {
	duo.mutation.SetDate(t)
	return duo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (duo *DayUpdateOne) SetNillableDate(t *time.Time) *DayUpdateOne {
	if t != nil {
		duo.SetDate(*t)
	}
	return duo
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (duo *DayUpdateOne) AddServiceIDs(ids ...int) *DayUpdateOne {
	duo.mutation.AddServiceIDs(ids...)
	return duo
}

// AddServices adds the "services" edges to the Service entity.
func (duo *DayUpdateOne) AddServices(s ...*Service) *DayUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.AddServiceIDs(ids...)
}

// Mutation returns the DayMutation object of the builder.
func (duo *DayUpdateOne) Mutation() *DayMutation {
	return duo.mutation
}

// ClearServices clears all "services" edges to the Service entity.
func (duo *DayUpdateOne) ClearServices() *DayUpdateOne {
	duo.mutation.ClearServices()
	return duo
}

// RemoveServiceIDs removes the "services" edge to Service entities by IDs.
func (duo *DayUpdateOne) RemoveServiceIDs(ids ...int) *DayUpdateOne {
	duo.mutation.RemoveServiceIDs(ids...)
	return duo
}

// RemoveServices removes "services" edges to Service entities.
func (duo *DayUpdateOne) RemoveServices(s ...*Service) *DayUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.RemoveServiceIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DayUpdateOne) Select(field string, fields ...string) *DayUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Day entity.
func (duo *DayUpdateOne) Save(ctx context.Context) (*Day, error) {
	var (
		err  error
		node *Day
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DayUpdateOne) SaveX(ctx context.Context) *Day {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DayUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DayUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DayUpdateOne) sqlSave(ctx context.Context) (_node *Day, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   day.Table,
			Columns: day.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: day.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Day.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, day.FieldID)
		for _, f := range fields {
			if !day.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != day.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: day.FieldDate,
		})
	}
	if duo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   day.ServicesTable,
			Columns: []string{day.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedServicesIDs(); len(nodes) > 0 && !duo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   day.ServicesTable,
			Columns: []string{day.ServicesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   day.ServicesTable,
			Columns: []string{day.ServicesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Day{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{day.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
