package entities_schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "time"
)

// Day holds the schema definition for the Day entity.
type Day struct {
    ent.Schema
}

// Fields of the Day.
func (Day) Fields() []ent.Field {
    return []ent.Field{
        field.Time("date").
            Default(time.Time{}),
    }
}

// Edges of the Day.
func (Day) Edges() []ent.Edge {
    return nil
}
