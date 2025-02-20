// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "use_id", Type: field.TypeUint64, Unique: true},
		{Name: "use_type", Type: field.TypeInt64, Default: 2},
		{Name: "use_username", Type: field.TypeString, Unique: true},
		{Name: "use_pwd", Type: field.TypeString},
		{Name: "use_email", Type: field.TypeString, Unique: true},
		{Name: "use_reset_token", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "use_reset_token_expiry", Type: field.TypeTime, Nullable: true},
		{Name: "use_access_token", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "use_access_token_expiry", Type: field.TypeTime, Nullable: true},
		{Name: "use_created_at", Type: field.TypeTime, Nullable: true},
		{Name: "use_updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "use_deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
