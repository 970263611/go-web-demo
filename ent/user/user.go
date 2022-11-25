// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserCode holds the string denoting the usercode field in the database.
	FieldUserCode = "user_code"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldDefaultPassword holds the string denoting the defaultpassword field in the database.
	FieldDefaultPassword = "default_password"
	// FieldIsAdmin holds the string denoting the isadmin field in the database.
	FieldIsAdmin = "is_admin"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldLoginTime holds the string denoting the logintime field in the database.
	FieldLoginTime = "login_time"
	// FieldAuthList holds the string denoting the authlist field in the database.
	FieldAuthList = "auth_list"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserCode,
	FieldUsername,
	FieldPassword,
	FieldDefaultPassword,
	FieldIsAdmin,
	FieldCreateTime,
	FieldLoginTime,
	FieldAuthList,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UserCodeValidator is a validator for the "userCode" field. It is called by the builders before save.
	UserCodeValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultPasswordValidator is a validator for the "defaultPassword" field. It is called by the builders before save.
	DefaultPasswordValidator func(string) error
	// DefaultCreateTime holds the default value on creation for the "createTime" field.
	DefaultCreateTime int64
)

// IsAdmin defines the type for the "isAdmin" enum field.
type IsAdmin string

// IsAdmin1 is the default value of the IsAdmin enum.
const DefaultIsAdmin = IsAdmin1

// IsAdmin values.
const (
	IsAdmin0 IsAdmin = "0"
	IsAdmin1 IsAdmin = "1"
)

func (ia IsAdmin) String() string {
	return string(ia)
}

// IsAdminValidator is a validator for the "isAdmin" field enum values. It is called by the builders before save.
func IsAdminValidator(ia IsAdmin) error {
	switch ia {
	case IsAdmin0, IsAdmin1:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for isAdmin field: %q", ia)
	}
}