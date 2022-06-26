package entities_schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TOC holds the schema definition for the TOC entity.
type TOC struct {
	ent.Schema
}

// Fields of the TOC.
func (TOC) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("???"),
	}
}

// Edges of the TOC.
func (TOC) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("services", Service.Type),
	}
}
