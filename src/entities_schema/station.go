package entities_schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// Station holds the schema definition for the Station entity.
type Station struct {
    ent.Schema
}

// Fields of the Station.
func (Station) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").
            Default("???"),
        field.String("crs").
            Default("ZZZ"),
    }
}

// Edges of the Station.
func (Station) Edges() []ent.Edge {
    return nil
}
