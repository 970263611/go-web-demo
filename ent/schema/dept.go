package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Dept holds the schema definition for the Dept entity.
type Dept struct {
	ent.Schema
}

// Fields of the Dept.
func (Dept) Fields() []ent.Field {
	return []ent.Field{
		field.String("deptId").NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("parentId").Optional(),
		field.String("ext").Optional(),
	}
}

// Edges of the Dept.
func (Dept) Edges() []ent.Edge {
	return nil
}
