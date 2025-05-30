// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"comb.com/banking/ent/predicate"
	"comb.com/banking/ent/transaction"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTransactionTime sets the "TransactionTime" field.
func (tu *TransactionUpdate) SetTransactionTime(t time.Time) *TransactionUpdate {
	tu.mutation.SetTransactionTime(t)
	return tu
}

// SetNillableTransactionTime sets the "TransactionTime" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableTransactionTime(t *time.Time) *TransactionUpdate {
	if t != nil {
		tu.SetTransactionTime(*t)
	}
	return tu
}

// SetFrom sets the "From" field.
func (tu *TransactionUpdate) SetFrom(i int) *TransactionUpdate {
	tu.mutation.ResetFrom()
	tu.mutation.SetFrom(i)
	return tu
}

// SetNillableFrom sets the "From" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableFrom(i *int) *TransactionUpdate {
	if i != nil {
		tu.SetFrom(*i)
	}
	return tu
}

// AddFrom adds i to the "From" field.
func (tu *TransactionUpdate) AddFrom(i int) *TransactionUpdate {
	tu.mutation.AddFrom(i)
	return tu
}

// SetTo sets the "To" field.
func (tu *TransactionUpdate) SetTo(i int) *TransactionUpdate {
	tu.mutation.ResetTo()
	tu.mutation.SetTo(i)
	return tu
}

// SetNillableTo sets the "To" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableTo(i *int) *TransactionUpdate {
	if i != nil {
		tu.SetTo(*i)
	}
	return tu
}

// AddTo adds i to the "To" field.
func (tu *TransactionUpdate) AddTo(i int) *TransactionUpdate {
	tu.mutation.AddTo(i)
	return tu
}

// SetAmount sets the "Amount" field.
func (tu *TransactionUpdate) SetAmount(i int) *TransactionUpdate {
	tu.mutation.ResetAmount()
	tu.mutation.SetAmount(i)
	return tu
}

// SetNillableAmount sets the "Amount" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableAmount(i *int) *TransactionUpdate {
	if i != nil {
		tu.SetAmount(*i)
	}
	return tu
}

// AddAmount adds i to the "Amount" field.
func (tu *TransactionUpdate) AddAmount(i int) *TransactionUpdate {
	tu.mutation.AddAmount(i)
	return tu
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.TransactionTime(); ok {
		_spec.SetField(transaction.FieldTransactionTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.From(); ok {
		_spec.SetField(transaction.FieldFrom, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedFrom(); ok {
		_spec.AddField(transaction.FieldFrom, field.TypeInt, value)
	}
	if value, ok := tu.mutation.To(); ok {
		_spec.SetField(transaction.FieldTo, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedTo(); ok {
		_spec.AddField(transaction.FieldTo, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedAmount(); ok {
		_spec.AddField(transaction.FieldAmount, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionMutation
}

// SetTransactionTime sets the "TransactionTime" field.
func (tuo *TransactionUpdateOne) SetTransactionTime(t time.Time) *TransactionUpdateOne {
	tuo.mutation.SetTransactionTime(t)
	return tuo
}

// SetNillableTransactionTime sets the "TransactionTime" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableTransactionTime(t *time.Time) *TransactionUpdateOne {
	if t != nil {
		tuo.SetTransactionTime(*t)
	}
	return tuo
}

// SetFrom sets the "From" field.
func (tuo *TransactionUpdateOne) SetFrom(i int) *TransactionUpdateOne {
	tuo.mutation.ResetFrom()
	tuo.mutation.SetFrom(i)
	return tuo
}

// SetNillableFrom sets the "From" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableFrom(i *int) *TransactionUpdateOne {
	if i != nil {
		tuo.SetFrom(*i)
	}
	return tuo
}

// AddFrom adds i to the "From" field.
func (tuo *TransactionUpdateOne) AddFrom(i int) *TransactionUpdateOne {
	tuo.mutation.AddFrom(i)
	return tuo
}

// SetTo sets the "To" field.
func (tuo *TransactionUpdateOne) SetTo(i int) *TransactionUpdateOne {
	tuo.mutation.ResetTo()
	tuo.mutation.SetTo(i)
	return tuo
}

// SetNillableTo sets the "To" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableTo(i *int) *TransactionUpdateOne {
	if i != nil {
		tuo.SetTo(*i)
	}
	return tuo
}

// AddTo adds i to the "To" field.
func (tuo *TransactionUpdateOne) AddTo(i int) *TransactionUpdateOne {
	tuo.mutation.AddTo(i)
	return tuo
}

// SetAmount sets the "Amount" field.
func (tuo *TransactionUpdateOne) SetAmount(i int) *TransactionUpdateOne {
	tuo.mutation.ResetAmount()
	tuo.mutation.SetAmount(i)
	return tuo
}

// SetNillableAmount sets the "Amount" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableAmount(i *int) *TransactionUpdateOne {
	if i != nil {
		tuo.SetAmount(*i)
	}
	return tuo
}

// AddAmount adds i to the "Amount" field.
func (tuo *TransactionUpdateOne) AddAmount(i int) *TransactionUpdateOne {
	tuo.mutation.AddAmount(i)
	return tuo
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tuo *TransactionUpdateOne) Where(ps ...predicate.Transaction) *TransactionUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.TransactionTime(); ok {
		_spec.SetField(transaction.FieldTransactionTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.From(); ok {
		_spec.SetField(transaction.FieldFrom, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedFrom(); ok {
		_spec.AddField(transaction.FieldFrom, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.To(); ok {
		_spec.SetField(transaction.FieldTo, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedTo(); ok {
		_spec.AddField(transaction.FieldTo, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedAmount(); ok {
		_spec.AddField(transaction.FieldAmount, field.TypeInt, value)
	}
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
