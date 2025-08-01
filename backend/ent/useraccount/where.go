// Code generated by ent, DO NOT EDIT.

package useraccount

import (
	"time"

	"comb.com/banking/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldLTE(FieldID, id))
}

// AccountNumber applies equality check predicate on the "account_number" field. It's identical to AccountNumberEQ.
func AccountNumber(v int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldEQ(FieldAccountNumber, v))
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldEQ(FieldBalance, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// AccountNumberEQ applies the EQ predicate on the "account_number" field.
func AccountNumberEQ(v int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldEQ(FieldAccountNumber, v))
}

// AccountNumberNEQ applies the NEQ predicate on the "account_number" field.
func AccountNumberNEQ(v int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldNEQ(FieldAccountNumber, v))
}

// AccountNumberIn applies the In predicate on the "account_number" field.
func AccountNumberIn(vs ...int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldIn(FieldAccountNumber, vs...))
}

// AccountNumberNotIn applies the NotIn predicate on the "account_number" field.
func AccountNumberNotIn(vs ...int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldNotIn(FieldAccountNumber, vs...))
}

// AccountNumberGT applies the GT predicate on the "account_number" field.
func AccountNumberGT(v int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldGT(FieldAccountNumber, v))
}

// AccountNumberGTE applies the GTE predicate on the "account_number" field.
func AccountNumberGTE(v int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldGTE(FieldAccountNumber, v))
}

// AccountNumberLT applies the LT predicate on the "account_number" field.
func AccountNumberLT(v int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldLT(FieldAccountNumber, v))
}

// AccountNumberLTE applies the LTE predicate on the "account_number" field.
func AccountNumberLTE(v int) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldLTE(FieldAccountNumber, v))
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldEQ(FieldBalance, v))
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldNEQ(FieldBalance, v))
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldIn(FieldBalance, vs...))
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldNotIn(FieldBalance, vs...))
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldGT(FieldBalance, v))
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldGTE(FieldBalance, v))
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldLT(FieldBalance, v))
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v float64) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldLTE(FieldBalance, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserAccount {
	return predicate.UserAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserAccount {
	return predicate.UserAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserAccount {
	return predicate.UserAccount(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserAccount) predicate.UserAccount {
	return predicate.UserAccount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserAccount) predicate.UserAccount {
	return predicate.UserAccount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserAccount) predicate.UserAccount {
	return predicate.UserAccount(sql.NotPredicates(p))
}
