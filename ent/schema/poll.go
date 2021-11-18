package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Poll holds the schema definition for the Poll entity.
type Poll struct {
	ent.Schema
}

// Fields of the Poll.
func (Poll) Fields() []ent.Field {
	return []ent.Field{
		// field.Int("id").
		// 	Default(-1),
		field.Int("ownerid").
			Default(-1),
		field.String("topic").
			Default("unknown"),
		field.JSON("options", []string{}).
			Optional(),
		field.Time("createtime").
			Default(time.Now),
		field.Time("modifytime").
			Default(time.Now),
	}
}

// Edges of the Poll.
func (Poll) Edges() []ent.Edge {
	return nil
}
