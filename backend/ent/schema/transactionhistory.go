package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TransactionHistory lưu lịch sử nạp/rút
type TransactionHistory struct {
	ent.Schema
}

func (TransactionHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "transaction_history"}, // khớp với DB thật
	}
}

// Fields định nghĩa các cột
func (TransactionHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			NotEmpty().
			Comment("deposit | withdraw"),
		field.Int("account_number_id").Unique(),
		field.Int64("amount"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges định nghĩa quan hệ
func (TransactionHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_accounts", UserAccount.Type).
			Ref("account_number_id").
			Field("account_number_id").
			Unique().
			Required(),
	}
}
