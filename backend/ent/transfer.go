// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"comb.com/banking/ent/transfer"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Transfer is the model entity for the Transfer schema.
type Transfer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TransactionTime holds the value of the "TransactionTime" field.
	TransactionTime time.Time `json:"TransactionTime,omitempty"`
	// From holds the value of the "From" field.
	From int `json:"From,omitempty"`
	// To holds the value of the "To" field.
	To int `json:"To,omitempty"`
	// Amount holds the value of the "Amount" field.
	Amount int `json:"Amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transfer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transfer.FieldID, transfer.FieldFrom, transfer.FieldTo, transfer.FieldAmount:
			values[i] = new(sql.NullInt64)
		case transfer.FieldTransactionTime, transfer.FieldCreatedAt, transfer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transfer fields.
func (t *Transfer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transfer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transfer.FieldTransactionTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field TransactionTime", values[i])
			} else if value.Valid {
				t.TransactionTime = value.Time
			}
		case transfer.FieldFrom:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field From", values[i])
			} else if value.Valid {
				t.From = int(value.Int64)
			}
		case transfer.FieldTo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field To", values[i])
			} else if value.Valid {
				t.To = int(value.Int64)
			}
		case transfer.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Amount", values[i])
			} else if value.Valid {
				t.Amount = int(value.Int64)
			}
		case transfer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case transfer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Transfer.
// This includes values selected through modifiers, order, etc.
func (t *Transfer) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Transfer.
// Note that you need to call Transfer.Unwrap() before calling this method if this Transfer
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transfer) Update() *TransferUpdateOne {
	return NewTransferClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Transfer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transfer) Unwrap() *Transfer {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transfer is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transfer) String() string {
	var builder strings.Builder
	builder.WriteString("Transfer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("TransactionTime=")
	builder.WriteString(t.TransactionTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("From=")
	builder.WriteString(fmt.Sprintf("%v", t.From))
	builder.WriteString(", ")
	builder.WriteString("To=")
	builder.WriteString(fmt.Sprintf("%v", t.To))
	builder.WriteString(", ")
	builder.WriteString("Amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Transfers is a parsable slice of Transfer.
type Transfers []*Transfer
