package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/pkg/models"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_time"),
		field.Time("end_time"),
		field.String("departure_station").
			NotEmpty(),
		field.String("arrival_station").
			NotEmpty(),
		field.String("id_number").
			NotEmpty(),
		field.String("phone_number").
			NotEmpty(),
		field.String("email").
			NotEmpty(),
		field.String("status").
			NotEmpty().
			Default(models.OrderPending),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique(),
		edge.From("validation", OrderValidation.Type).
			Ref("order"),
	}
}
