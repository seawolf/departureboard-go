package entities_schema

import "entgo.io/ent"

// CallingPoint holds the schema definition for the CallingPoint entity.
type CallingPoint struct {
    ent.Schema
}

// Fields of the CallingPoint.
func (CallingPoint) Fields() []ent.Field {
    return nil
}

// Edges of the CallingPoint.
func (CallingPoint) Edges() []ent.Edge {
    return nil
}
