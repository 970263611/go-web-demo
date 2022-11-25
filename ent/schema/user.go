package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
	Menu
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("userCode").Unique().NotEmpty().MaxLen(10),
		field.String("username").NotEmpty().MaxLen(10),
		field.String("password").NotEmpty().MaxLen(50).Optional(),
		field.String("defaultPassword").NotEmpty().MaxLen(50),
		field.Enum("isAdmin").Values("0", "1").Default("1").Optional(),
		field.Int64("createTime").Default(time.Now().Unix()),
		field.Int64("loginTime").Optional(),
		field.JSON("authList", map[string]any{}).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
