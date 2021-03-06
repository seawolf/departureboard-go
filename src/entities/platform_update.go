// Code generated by entc, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/callingpoint"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/platform"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/predicate"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/station"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlatformUpdate is the builder for updating Platform entities.
type PlatformUpdate struct {
	config
	hooks    []Hook
	mutation *PlatformMutation
}

// Where appends a list predicates to the PlatformUpdate builder.
func (pu *PlatformUpdate) Where(ps ...predicate.Platform) *PlatformUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PlatformUpdate) SetName(s string) *PlatformUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PlatformUpdate) SetNillableName(s *string) *PlatformUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetStationID sets the "station" edge to the Station entity by ID.
func (pu *PlatformUpdate) SetStationID(id int) *PlatformUpdate {
	pu.mutation.SetStationID(id)
	return pu
}

// SetNillableStationID sets the "station" edge to the Station entity by ID if the given value is not nil.
func (pu *PlatformUpdate) SetNillableStationID(id *int) *PlatformUpdate {
	if id != nil {
		pu = pu.SetStationID(*id)
	}
	return pu
}

// SetStation sets the "station" edge to the Station entity.
func (pu *PlatformUpdate) SetStation(s *Station) *PlatformUpdate {
	return pu.SetStationID(s.ID)
}

// AddCallingPointIDs adds the "calling_points" edge to the CallingPoint entity by IDs.
func (pu *PlatformUpdate) AddCallingPointIDs(ids ...int) *PlatformUpdate {
	pu.mutation.AddCallingPointIDs(ids...)
	return pu
}

// AddCallingPoints adds the "calling_points" edges to the CallingPoint entity.
func (pu *PlatformUpdate) AddCallingPoints(c ...*CallingPoint) *PlatformUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCallingPointIDs(ids...)
}

// Mutation returns the PlatformMutation object of the builder.
func (pu *PlatformUpdate) Mutation() *PlatformMutation {
	return pu.mutation
}

// ClearStation clears the "station" edge to the Station entity.
func (pu *PlatformUpdate) ClearStation() *PlatformUpdate {
	pu.mutation.ClearStation()
	return pu
}

// ClearCallingPoints clears all "calling_points" edges to the CallingPoint entity.
func (pu *PlatformUpdate) ClearCallingPoints() *PlatformUpdate {
	pu.mutation.ClearCallingPoints()
	return pu
}

// RemoveCallingPointIDs removes the "calling_points" edge to CallingPoint entities by IDs.
func (pu *PlatformUpdate) RemoveCallingPointIDs(ids ...int) *PlatformUpdate {
	pu.mutation.RemoveCallingPointIDs(ids...)
	return pu
}

// RemoveCallingPoints removes "calling_points" edges to CallingPoint entities.
func (pu *PlatformUpdate) RemoveCallingPoints(c ...*CallingPoint) *PlatformUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCallingPointIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlatformUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformMutation)
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
				return 0, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
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
func (pu *PlatformUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlatformUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlatformUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlatformUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   platform.Table,
			Columns: platform.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: platform.FieldID,
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
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: platform.FieldName,
		})
	}
	if pu.mutation.StationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   platform.StationTable,
			Columns: []string{platform.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: station.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   platform.StationTable,
			Columns: []string{platform.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: station.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CallingPointsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.CallingPointsTable,
			Columns: []string{platform.CallingPointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: callingpoint.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCallingPointsIDs(); len(nodes) > 0 && !pu.mutation.CallingPointsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.CallingPointsTable,
			Columns: []string{platform.CallingPointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: callingpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CallingPointsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.CallingPointsTable,
			Columns: []string{platform.CallingPointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: callingpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platform.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PlatformUpdateOne is the builder for updating a single Platform entity.
type PlatformUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlatformMutation
}

// SetName sets the "name" field.
func (puo *PlatformUpdateOne) SetName(s string) *PlatformUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableName(s *string) *PlatformUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetStationID sets the "station" edge to the Station entity by ID.
func (puo *PlatformUpdateOne) SetStationID(id int) *PlatformUpdateOne {
	puo.mutation.SetStationID(id)
	return puo
}

// SetNillableStationID sets the "station" edge to the Station entity by ID if the given value is not nil.
func (puo *PlatformUpdateOne) SetNillableStationID(id *int) *PlatformUpdateOne {
	if id != nil {
		puo = puo.SetStationID(*id)
	}
	return puo
}

// SetStation sets the "station" edge to the Station entity.
func (puo *PlatformUpdateOne) SetStation(s *Station) *PlatformUpdateOne {
	return puo.SetStationID(s.ID)
}

// AddCallingPointIDs adds the "calling_points" edge to the CallingPoint entity by IDs.
func (puo *PlatformUpdateOne) AddCallingPointIDs(ids ...int) *PlatformUpdateOne {
	puo.mutation.AddCallingPointIDs(ids...)
	return puo
}

// AddCallingPoints adds the "calling_points" edges to the CallingPoint entity.
func (puo *PlatformUpdateOne) AddCallingPoints(c ...*CallingPoint) *PlatformUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCallingPointIDs(ids...)
}

// Mutation returns the PlatformMutation object of the builder.
func (puo *PlatformUpdateOne) Mutation() *PlatformMutation {
	return puo.mutation
}

// ClearStation clears the "station" edge to the Station entity.
func (puo *PlatformUpdateOne) ClearStation() *PlatformUpdateOne {
	puo.mutation.ClearStation()
	return puo
}

// ClearCallingPoints clears all "calling_points" edges to the CallingPoint entity.
func (puo *PlatformUpdateOne) ClearCallingPoints() *PlatformUpdateOne {
	puo.mutation.ClearCallingPoints()
	return puo
}

// RemoveCallingPointIDs removes the "calling_points" edge to CallingPoint entities by IDs.
func (puo *PlatformUpdateOne) RemoveCallingPointIDs(ids ...int) *PlatformUpdateOne {
	puo.mutation.RemoveCallingPointIDs(ids...)
	return puo
}

// RemoveCallingPoints removes "calling_points" edges to CallingPoint entities.
func (puo *PlatformUpdateOne) RemoveCallingPoints(c ...*CallingPoint) *PlatformUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCallingPointIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlatformUpdateOne) Select(field string, fields ...string) *PlatformUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Platform entity.
func (puo *PlatformUpdateOne) Save(ctx context.Context) (*Platform, error) {
	var (
		err  error
		node *Platform
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlatformMutation)
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
				return nil, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
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
func (puo *PlatformUpdateOne) SaveX(ctx context.Context) *Platform {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlatformUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlatformUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlatformUpdateOne) sqlSave(ctx context.Context) (_node *Platform, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   platform.Table,
			Columns: platform.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: platform.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Platform.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, platform.FieldID)
		for _, f := range fields {
			if !platform.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != platform.FieldID {
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
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: platform.FieldName,
		})
	}
	if puo.mutation.StationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   platform.StationTable,
			Columns: []string{platform.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: station.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   platform.StationTable,
			Columns: []string{platform.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: station.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CallingPointsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.CallingPointsTable,
			Columns: []string{platform.CallingPointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: callingpoint.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCallingPointsIDs(); len(nodes) > 0 && !puo.mutation.CallingPointsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.CallingPointsTable,
			Columns: []string{platform.CallingPointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: callingpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CallingPointsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platform.CallingPointsTable,
			Columns: []string{platform.CallingPointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: callingpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Platform{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platform.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
