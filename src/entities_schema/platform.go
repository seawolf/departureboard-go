package entities_schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Platform holds the schema definition for the Platform entity.
type Platform struct {
	ent.Schema
}

// Fields of the Platform.
func (Platform) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("-"),
	}
}

// Edges of the Platform.
func (Platform) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("station", Station.Type).
			Ref("platforms").
			Unique(),
	}
}
