package entities_schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
	return []ent.Edge{
		edge.From("platform", Platform.Type).
			Ref("calling_points").
			Unique(),

		edge.From("service", Service.Type).
			Ref("calling_points").
			Unique(),
	}
}
