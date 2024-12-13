package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Todo holds the schema definition for the Todo entity.
type Routine struct {
	ent.Schema
}

// Fields of the Todo.
func (Routine) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.Text("description").
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Todo.
func (Routine) Edges() []ent.Edge {
	return nil
}

func (Routine) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("completed"),
		index.Fields("created_at"),
	}
}
