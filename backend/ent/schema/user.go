package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("email"),
		field.String("password"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_id", UserAccount.Type).Unique(), // 1-1
		edge.To("profile", UserProfile.Type).Unique(),                        // 1-1
		edge.To("tokens", Token.Type).StorageKey(edge.Column("user_tokens")), // 1-n
	}
}
