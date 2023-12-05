// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/ordervalidation"
)

// OrderValidation is the model entity for the OrderValidation schema.
type OrderValidation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// JessionID holds the value of the "jession_id" field.
	JessionID string `json:"jession_id,omitempty"`
	// CaptchaImage holds the value of the "captcha_image" field.
	CaptchaImage string `json:"captcha_image,omitempty"`
	// Cookies holds the value of the "cookies" field.
	Cookies string `json:"cookies,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderValidationQuery when eager-loading is set.
	Edges                  OrderValidationEdges `json:"edges"`
	order_validation_order *int
	selectValues           sql.SelectValues
}

// OrderValidationEdges holds the relations/edges for other nodes in the graph.
type OrderValidationEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderValidationEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Order == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderValidation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ordervalidation.FieldID:
			values[i] = new(sql.NullInt64)
		case ordervalidation.FieldJessionID, ordervalidation.FieldCaptchaImage, ordervalidation.FieldCookies:
			values[i] = new(sql.NullString)
		case ordervalidation.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case ordervalidation.ForeignKeys[0]: // order_validation_order
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderValidation fields.
func (ov *OrderValidation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ordervalidation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ov.ID = int(value.Int64)
		case ordervalidation.FieldJessionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jession_id", values[i])
			} else if value.Valid {
				ov.JessionID = value.String
			}
		case ordervalidation.FieldCaptchaImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field captcha_image", values[i])
			} else if value.Valid {
				ov.CaptchaImage = value.String
			}
		case ordervalidation.FieldCookies:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cookies", values[i])
			} else if value.Valid {
				ov.Cookies = value.String
			}
		case ordervalidation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ov.CreatedAt = value.Time
			}
		case ordervalidation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_validation_order", value)
			} else if value.Valid {
				ov.order_validation_order = new(int)
				*ov.order_validation_order = int(value.Int64)
			}
		default:
			ov.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderValidation.
// This includes values selected through modifiers, order, etc.
func (ov *OrderValidation) Value(name string) (ent.Value, error) {
	return ov.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the OrderValidation entity.
func (ov *OrderValidation) QueryOrder() *OrderQuery {
	return NewOrderValidationClient(ov.config).QueryOrder(ov)
}

// Update returns a builder for updating this OrderValidation.
// Note that you need to call OrderValidation.Unwrap() before calling this method if this OrderValidation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ov *OrderValidation) Update() *OrderValidationUpdateOne {
	return NewOrderValidationClient(ov.config).UpdateOne(ov)
}

// Unwrap unwraps the OrderValidation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ov *OrderValidation) Unwrap() *OrderValidation {
	_tx, ok := ov.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderValidation is not a transactional entity")
	}
	ov.config.driver = _tx.drv
	return ov
}

// String implements the fmt.Stringer.
func (ov *OrderValidation) String() string {
	var builder strings.Builder
	builder.WriteString("OrderValidation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ov.ID))
	builder.WriteString("jession_id=")
	builder.WriteString(ov.JessionID)
	builder.WriteString(", ")
	builder.WriteString("captcha_image=")
	builder.WriteString(ov.CaptchaImage)
	builder.WriteString(", ")
	builder.WriteString("cookies=")
	builder.WriteString(ov.Cookies)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ov.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// OrderValidations is a parsable slice of OrderValidation.
type OrderValidations []*OrderValidation