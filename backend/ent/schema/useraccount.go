package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserAccount struct {
	ent.Schema
}

func (UserAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_accounts"}, // khớp với DB thật
	}
}

func (UserAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("account_number"),
		field.Int64("balance"),
	}
}

func (UserAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("accounts").
			Unique(),
		edge.To("transactions", Transaction.Type),
		edge.To("outgoing_transfers", Transfer.Type),
		edge.To("incoming_transfers", Transfer.Type),
	}
}
