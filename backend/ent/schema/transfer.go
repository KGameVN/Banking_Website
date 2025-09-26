package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Transfer struct {
	ent.Schema
}

func (Transfer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount"),
	}
}

func (Transfer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("from_account", UserAccount.Type).
			Ref("outgoing_transfers").
			Unique(),
		edge.From("to_account", UserAccount.Type).
			Ref("incoming_transfers").
			Unique(),
	}
}
