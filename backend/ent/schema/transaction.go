package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Transaction struct {
	ent.Schema
}

func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount"),
	}
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", UserAccount.Type).
			Ref("transactions").
			Unique(),
	}
}
