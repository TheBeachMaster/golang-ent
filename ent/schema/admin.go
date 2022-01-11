package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("email_address").NotEmpty(),
		field.String("password").NotEmpty(),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}
