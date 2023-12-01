// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldDepartureStation holds the string denoting the departure_station field in the database.
	FieldDepartureStation = "departure_station"
	// FieldArrivalStation holds the string denoting the arrival_station field in the database.
	FieldArrivalStation = "arrival_station"
	// FieldIDNumber holds the string denoting the id_number field in the database.
	FieldIDNumber = "id_number"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "orders"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "order_user"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldStartTime,
	FieldEndTime,
	FieldDepartureStation,
	FieldArrivalStation,
	FieldIDNumber,
	FieldPhoneNumber,
	FieldEmail,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DepartureStationValidator is a validator for the "departure_station" field. It is called by the builders before save.
	DepartureStationValidator func(string) error
	// ArrivalStationValidator is a validator for the "arrival_station" field. It is called by the builders before save.
	ArrivalStationValidator func(string) error
	// IDNumberValidator is a validator for the "id_number" field. It is called by the builders before save.
	IDNumberValidator func(string) error
	// PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Order queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByDepartureStation orders the results by the departure_station field.
func ByDepartureStation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartureStation, opts...).ToFunc()
}

// ByArrivalStation orders the results by the arrival_station field.
func ByArrivalStation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArrivalStation, opts...).ToFunc()
}

// ByIDNumber orders the results by the id_number field.
func ByIDNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIDNumber, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
