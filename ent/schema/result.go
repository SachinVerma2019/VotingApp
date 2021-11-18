package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Result holds the schema definition for the Result entity.
type Result struct {
	ent.Schema
}

// Fields of the Result.
func (Result) Fields() []ent.Field {
	return []ent.Field{
		field.Int("userid").
			Default(-1),
		field.Int("pollid").
			Default(-1),
		field.String("option").
			Default("unknown"),
		field.Time("createtime").
			Default(time.Now),
		field.Time("modifytime").
			Default(time.Now),
	}
}

// Edges of the Result.
func (Result) Edges() []ent.Edge {
	return nil
}
