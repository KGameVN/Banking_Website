package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserProfile struct {
	ent.Schema
}

func (UserProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstname"),
		field.String("lastname"),
		field.String("address"),
		field.String("gender"),
		field.String("birthday"),
	}
}

func (UserProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("profile").
			Unique(),
	}
}
