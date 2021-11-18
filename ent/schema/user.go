package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// // Annotations of the User.
// func (User) Annotations() []schema.Annotation {
// 	return []schema.Annotation{
// 		entsql.Annotation{Table: "User"},
// 	}
// }

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
		field.String("email").
			Default("unknown"),
		field.String("password").
			Default("xxxxx"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
