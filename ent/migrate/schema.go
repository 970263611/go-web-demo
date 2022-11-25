// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MenusColumns holds the columns for the "menus" table.
	MenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "name_id", Type: field.TypeString},
		{Name: "parent_id", Type: field.TypeInt64},
	}
	// MenusTable holds the schema information for the "menus" table.
	MenusTable = &schema.Table{
		Name:       "menus",
		Columns:    MenusColumns,
		PrimaryKey: []*schema.Column{MenusColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_code", Type: field.TypeString, Unique: true, Size: 10},
		{Name: "username", Type: field.TypeString, Size: 10},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "default_password", Type: field.TypeString, Size: 50},
		{Name: "is_admin", Type: field.TypeEnum, Nullable: true, Enums: []string{"0", "1"}, Default: "1"},
		{Name: "create_time", Type: field.TypeInt64, Default: 1669283392},
		{Name: "login_time", Type: field.TypeInt64, Nullable: true},
		{Name: "auth_list", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MenusTable,
		UsersTable,
	}
)

func init() {
}
