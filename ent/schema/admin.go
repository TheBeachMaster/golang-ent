package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Album holds the schema definition for the Album entity.
type Admin struct {
	ent.Schema
}

// Fields of the Album.
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

// Edges of the Album.
func (Admin) Edges() []ent.Edge {
	return nil
}
