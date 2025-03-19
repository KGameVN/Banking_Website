package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Time("TransactionTime"),
		field.Int("From").SchemaType(map[string]string{
			"postgres": "bigint",
		}),
		field.Int("To").SchemaType(map[string]string{
			"postgres": "bigint",
		}),
		field.Int("Amount").SchemaType(map[string]string{
			"postgres": "bigint",
		}),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return nil
}
