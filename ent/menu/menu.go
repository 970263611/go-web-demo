// Code generated by ent, DO NOT EDIT.

package menu

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNameId holds the string denoting the nameid field in the database.
	FieldNameId = "name_id"
	// FieldParentId holds the string denoting the parentid field in the database.
	FieldParentId = "parent_id"
	// Table holds the table name of the menu in the database.
	Table = "menus"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNameId,
	FieldParentId,
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