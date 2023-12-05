package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrderValidation struct {
	ent.Schema
}

func (OrderValidation) Fields() []ent.Field {
	return []ent.Field{
		field.String("jession_id").
			NotEmpty(),
		field.String("captcha_image").
			NotEmpty(),
		field.String("cookies").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Order.
func (OrderValidation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			Required().
			Unique(),
	}
}
