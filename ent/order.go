// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/user"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StartTime holds the value of the "start_time" field.
	StartTime time.Time `json:"start_time,omitempty"`
	// EndTime holds the value of the "end_time" field.
	EndTime time.Time `json:"end_time,omitempty"`
	// DepartureStation holds the value of the "departure_station" field.
	DepartureStation string `json:"departure_station,omitempty"`
	// ArrivalStation holds the value of the "arrival_station" field.
	ArrivalStation string `json:"arrival_station,omitempty"`
	// IDNumber holds the value of the "id_number" field.
	IDNumber string `json:"id_number,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// ErrorMessage holds the value of the "error_message" field.
	ErrorMessage string `json:"error_message,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges        OrderEdges `json:"edges"`
	order_user   *int
	selectValues sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Validation holds the value of the validation edge.
	Validation []*OrderValidation `json:"validation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ValidationOrErr returns the Validation value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) ValidationOrErr() ([]*OrderValidation, error) {
	if e.loadedTypes[1] {
		return e.Validation, nil
	}
	return nil, &NotLoadedError{edge: "validation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			values[i] = new(sql.NullInt64)
		case order.FieldDepartureStation, order.FieldArrivalStation, order.FieldIDNumber, order.FieldPhoneNumber, order.FieldEmail, order.FieldStatus, order.FieldErrorMessage:
			values[i] = new(sql.NullString)
		case order.FieldStartTime, order.FieldEndTime, order.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case order.ForeignKeys[0]: // order_user
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				o.StartTime = value.Time
			}
		case order.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				o.EndTime = value.Time
			}
		case order.FieldDepartureStation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field departure_station", values[i])
			} else if value.Valid {
				o.DepartureStation = value.String
			}
		case order.FieldArrivalStation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field arrival_station", values[i])
			} else if value.Valid {
				o.ArrivalStation = value.String
			}
		case order.FieldIDNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id_number", values[i])
			} else if value.Valid {
				o.IDNumber = value.String
			}
		case order.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				o.PhoneNumber = value.String
			}
		case order.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				o.Email = value.String
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = value.String
			}
		case order.FieldErrorMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_message", values[i])
			} else if value.Valid {
				o.ErrorMessage = value.String
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_user", value)
			} else if value.Valid {
				o.order_user = new(int)
				*o.order_user = int(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Order entity.
func (o *Order) QueryUser() *UserQuery {
	return NewOrderClient(o.config).QueryUser(o)
}

// QueryValidation queries the "validation" edge of the Order entity.
func (o *Order) QueryValidation() *OrderValidationQuery {
	return NewOrderClient(o.config).QueryValidation(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("start_time=")
	builder.WriteString(o.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(o.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("departure_station=")
	builder.WriteString(o.DepartureStation)
	builder.WriteString(", ")
	builder.WriteString("arrival_station=")
	builder.WriteString(o.ArrivalStation)
	builder.WriteString(", ")
	builder.WriteString("id_number=")
	builder.WriteString(o.IDNumber)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(o.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(o.Email)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(o.Status)
	builder.WriteString(", ")
	builder.WriteString("error_message=")
	builder.WriteString(o.ErrorMessage)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
