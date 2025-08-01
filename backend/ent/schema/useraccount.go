package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserAccount holds the schema definition for the UserAccount entity.
type UserAccount struct {
	ent.Schema
}

// Fields of the UserAccount.
func (UserAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_number").Unique().
			SchemaType(map[string]string{
				"postgres": "bigint", // Tạo kiểu `bigint` trong PostgreSQL
			}),
		field.Float("balance").Default(0.0),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserAccount.
func (UserAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("account").Unique().Required(),
	}
}
