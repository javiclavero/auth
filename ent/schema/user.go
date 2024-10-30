// ent/schema/user.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("use_id").
			Unique().
			StorageKey("use_id").
			StructTag(`json:"use_id"`),

		field.Int64("use_type").
			Default(2).
			StorageKey("use_type"),

		field.String("use_username").
			Unique().
			NotEmpty().
			StorageKey("use_username").
			StructTag(`json:"username"`),

		field.String("use_pwd").
			NotEmpty().
			Sensitive().
			StorageKey("use_pwd"),

		field.String("use_email").
			Unique().
			NotEmpty().
			StorageKey("use_email").
			StructTag(`json:"email"`),

		field.String("use_reset_token").
			MaxLen(255).
			Optional().
			Nillable().
			StorageKey("use_reset_token"),

		field.Time("use_reset_token_expiry").
			Optional().
			Nillable().
			StorageKey("use_reset_token_expiry"),

		field.String("use_access_token").
			MaxLen(255).
			Optional().
			Nillable().
			StorageKey("use_access_token"),

		field.Time("use_access_token_expiry").
			Optional().
			Nillable().
			StorageKey("use_access_token_expiry"),

		field.Time("use_created_at").
			Default(time.Now).
			Optional().
			Nillable().
			StorageKey("use_created_at"),

		field.Time("use_updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Optional().
			Nillable().
			StorageKey("use_updated_at"),

		field.Time("use_deleted_at").
			Optional().
			Nillable().
			StorageKey("use_deleted_at"),
	}
}

// PrimaryKeyFields establece `use_id` como la clave primaria.
func (User) PrimaryKeyFields() []string {
	return []string{"use_id"}
}
