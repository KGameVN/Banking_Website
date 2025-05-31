package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// LoginToken holds the schema definition for the LoginToken entity.
type LoginToken struct {
	ent.Schema
}

// Fields of the LoginToken.
func (LoginToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").NotEmpty(),       // JWT token
		field.Time("expiredtime").Optional(),   // Cho phép expiredtime là null, có thể thay thế bằng Default nếu cần
		field.Time("created_at").Default(time.Now), // Thời gian tạo token
	}
}

// Edges of the LoginToken.
func (LoginToken) Edges() []ent.Edge {
	return []ent.Edge{
		// Mối quan hệ với User, mỗi LoginToken thuộc về một User
		edge.To("user", User.Type).Unique().Required(),
	}
}
