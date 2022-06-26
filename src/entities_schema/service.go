package entities_schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			Default("00000"),
		field.String("headcode").
			Default("0Z00"),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("toc", TOC.Type).
			Ref("services").
			Unique(),
		edge.From("day", Day.Type).
			Ref("services").
			Unique(),

		edge.To("calling_points", CallingPoint.Type),
	}
}
