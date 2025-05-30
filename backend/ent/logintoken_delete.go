// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"comb.com/banking/ent/logintoken"
	"comb.com/banking/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginTokenDelete is the builder for deleting a LoginToken entity.
type LoginTokenDelete struct {
	config
	hooks    []Hook
	mutation *LoginTokenMutation
}

// Where appends a list predicates to the LoginTokenDelete builder.
func (ltd *LoginTokenDelete) Where(ps ...predicate.LoginToken) *LoginTokenDelete {
	ltd.mutation.Where(ps...)
	return ltd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ltd *LoginTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ltd.sqlExec, ltd.mutation, ltd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ltd *LoginTokenDelete) ExecX(ctx context.Context) int {
	n, err := ltd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ltd *LoginTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(logintoken.Table, sqlgraph.NewFieldSpec(logintoken.FieldID, field.TypeInt))
	if ps := ltd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ltd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ltd.mutation.done = true
	return affected, err
}

// LoginTokenDeleteOne is the builder for deleting a single LoginToken entity.
type LoginTokenDeleteOne struct {
	ltd *LoginTokenDelete
}

// Where appends a list predicates to the LoginTokenDelete builder.
func (ltdo *LoginTokenDeleteOne) Where(ps ...predicate.LoginToken) *LoginTokenDeleteOne {
	ltdo.ltd.mutation.Where(ps...)
	return ltdo
}

// Exec executes the deletion query.
func (ltdo *LoginTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := ltdo.ltd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{logintoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ltdo *LoginTokenDeleteOne) ExecX(ctx context.Context) {
	if err := ltdo.Exec(ctx); err != nil {
		panic(err)
	}
}
