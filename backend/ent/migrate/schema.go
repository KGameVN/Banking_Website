// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LoginTokensColumns holds the columns for the "login_tokens" table.
	LoginTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString},
		{Name: "expiredtime", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// LoginTokensTable holds the schema information for the "login_tokens" table.
	LoginTokensTable = &schema.Table{
		Name:       "login_tokens",
		Columns:    LoginTokensColumns,
		PrimaryKey: []*schema.Column{LoginTokensColumns[0]},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "transaction_time", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
	}
	// TransfersColumns holds the columns for the "transfers" table.
	TransfersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "transaction_time", Type: field.TypeTime},
		{Name: "from", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "to", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "amount", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TransfersTable holds the schema information for the "transfers" table.
	TransfersTable = &schema.Table{
		Name:       "transfers",
		Columns:    TransfersColumns,
		PrimaryKey: []*schema.Column{TransfersColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "login_token_user", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_login_tokens_user",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{LoginTokensColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserAccountsColumns holds the columns for the "user_accounts" table.
	UserAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "account_number", Type: field.TypeInt, Unique: true, SchemaType: map[string]string{"postgres": "bigint"}},
		{Name: "balance", Type: field.TypeFloat64, Default: 0},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_account", Type: field.TypeInt, Unique: true},
	}
	// UserAccountsTable holds the schema information for the "user_accounts" table.
	UserAccountsTable = &schema.Table{
		Name:       "user_accounts",
		Columns:    UserAccountsColumns,
		PrimaryKey: []*schema.Column{UserAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_accounts_users_account",
				Columns:    []*schema.Column{UserAccountsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserProfilesColumns holds the columns for the "user_profiles" table.
	UserProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "firstname", Type: field.TypeString, Unique: true},
		{Name: "lastname", Type: field.TypeString, Unique: true},
		{Name: "cmnd", Type: field.TypeString, Unique: true},
		{Name: "address", Type: field.TypeString},
		{Name: "gender", Type: field.TypeBool, Default: true},
		{Name: "birthday", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UserProfilesTable holds the schema information for the "user_profiles" table.
	UserProfilesTable = &schema.Table{
		Name:       "user_profiles",
		Columns:    UserProfilesColumns,
		PrimaryKey: []*schema.Column{UserProfilesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LoginTokensTable,
		TransactionsTable,
		TransfersTable,
		UsersTable,
		UserAccountsTable,
		UserProfilesTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = LoginTokensTable
	UserAccountsTable.ForeignKeys[0].RefTable = UsersTable
}
