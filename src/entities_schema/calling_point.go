package entities_schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CallingPoint holds the schema definition for the CallingPoint entity.
type CallingPoint struct {
	ent.Schema
}

// Fields of the CallingPoint.
func (CallingPoint) Fields() []ent.Field {
	return []ent.Field{
		field.Time("arrival_time").
			Default(time.Time{}),
		field.Time("departure_time").
			Default(time.Time{}),
	}
}

// Edges of the CallingPoint.
func (CallingPoint) Edges() []ent.Edge {
	return nil
}
