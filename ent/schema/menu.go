package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("nameId"),
		field.Int64("parentId"),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return nil
}
