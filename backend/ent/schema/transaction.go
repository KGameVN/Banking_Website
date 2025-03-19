// ent/schema/transaction.go
package schema

import (
	"time"

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
		// Thời gian giao dịch
		field.Time("TransactionTime").
			Default(time.Now), // Mặc định thời gian hiện tại

		// ID tài khoản người gửi
		field.Int("From").
			SchemaType(map[string]string{
				"postgres": "bigint", // Tạo kiểu `bigint` trong PostgreSQL
			}),

		// ID tài khoản người nhận
		field.Int("To").
			SchemaType(map[string]string{
				"postgres": "bigint", // Tạo kiểu `bigint` trong PostgreSQL
			}),

		// Số tiền giao dịch
		field.Int("Amount").
			SchemaType(map[string]string{
				"postgres": "bigint", // Tạo kiểu `bigint` trong PostgreSQL
			}),

		// Thời gian tạo giao dịch
		field.Time("created_at").
			Default(time.Now), // Mặc định thời gian hiện tại

		// Thời gian cập nhật giao dịch
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now), // Mặc định thời gian hiện tại và sẽ tự động cập nhật khi có thay đổi
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return nil // Không có edges cho bảng này
}
