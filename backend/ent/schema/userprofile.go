package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type UserProfile struct {
	ent.Schema
}

// Fields of the User.
func (UserProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstname").Unique().NotEmpty(),
		field.String("lastname").Unique().NotEmpty(),
		field.String("cmnd").Unique().NotEmpty(),
		field.String("address").NotEmpty(),
		field.Bool("gender").Default(true),
		field.String("birthday").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}
