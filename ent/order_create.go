// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/ordervalidation"
	"github.com/mikestefanello/pagoda/ent/user"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetStartTime sets the "start_time" field.
func (oc *OrderCreate) SetStartTime(t time.Time) *OrderCreate {
	oc.mutation.SetStartTime(t)
	return oc
}

// SetEndTime sets the "end_time" field.
func (oc *OrderCreate) SetEndTime(t time.Time) *OrderCreate {
	oc.mutation.SetEndTime(t)
	return oc
}

// SetDepartureStation sets the "departure_station" field.
func (oc *OrderCreate) SetDepartureStation(s string) *OrderCreate {
	oc.mutation.SetDepartureStation(s)
	return oc
}

// SetArrivalStation sets the "arrival_station" field.
func (oc *OrderCreate) SetArrivalStation(s string) *OrderCreate {
	oc.mutation.SetArrivalStation(s)
	return oc
}

// SetIDNumber sets the "id_number" field.
func (oc *OrderCreate) SetIDNumber(s string) *OrderCreate {
	oc.mutation.SetIDNumber(s)
	return oc
}

// SetPhoneNumber sets the "phone_number" field.
func (oc *OrderCreate) SetPhoneNumber(s string) *OrderCreate {
	oc.mutation.SetPhoneNumber(s)
	return oc
}

// SetEmail sets the "email" field.
func (oc *OrderCreate) SetEmail(s string) *OrderCreate {
	oc.mutation.SetEmail(s)
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(s string) *OrderCreate {
	oc.mutation.SetStatus(s)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrderCreate) SetNillableStatus(s *string) *OrderCreate {
	if s != nil {
		oc.SetStatus(*s)
	}
	return oc
}

// SetErrorMessage sets the "error_message" field.
func (oc *OrderCreate) SetErrorMessage(s string) *OrderCreate {
	oc.mutation.SetErrorMessage(s)
	return oc
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (oc *OrderCreate) SetNillableErrorMessage(s *string) *OrderCreate {
	if s != nil {
		oc.SetErrorMessage(*s)
	}
	return oc
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (oc *OrderCreate) SetUserID(id int) *OrderCreate {
	oc.mutation.SetUserID(id)
	return oc
}

// SetUser sets the "user" edge to the User entity.
func (oc *OrderCreate) SetUser(u *User) *OrderCreate {
	return oc.SetUserID(u.ID)
}

// AddValidationIDs adds the "validation" edge to the OrderValidation entity by IDs.
func (oc *OrderCreate) AddValidationIDs(ids ...int) *OrderCreate {
	oc.mutation.AddValidationIDs(ids...)
	return oc
}

// AddValidation adds the "validation" edges to the OrderValidation entity.
func (oc *OrderCreate) AddValidation(o ...*OrderValidation) *OrderCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddValidationIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.Status(); !ok {
		v := order.DefaultStatus
		oc.mutation.SetStatus(v)
	}
	if _, ok := oc.mutation.ErrorMessage(); !ok {
		v := order.DefaultErrorMessage
		oc.mutation.SetErrorMessage(v)
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := order.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "Order.start_time"`)}
	}
	if _, ok := oc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "Order.end_time"`)}
	}
	if _, ok := oc.mutation.DepartureStation(); !ok {
		return &ValidationError{Name: "departure_station", err: errors.New(`ent: missing required field "Order.departure_station"`)}
	}
	if v, ok := oc.mutation.DepartureStation(); ok {
		if err := order.DepartureStationValidator(v); err != nil {
			return &ValidationError{Name: "departure_station", err: fmt.Errorf(`ent: validator failed for field "Order.departure_station": %w`, err)}
		}
	}
	if _, ok := oc.mutation.ArrivalStation(); !ok {
		return &ValidationError{Name: "arrival_station", err: errors.New(`ent: missing required field "Order.arrival_station"`)}
	}
	if v, ok := oc.mutation.ArrivalStation(); ok {
		if err := order.ArrivalStationValidator(v); err != nil {
			return &ValidationError{Name: "arrival_station", err: fmt.Errorf(`ent: validator failed for field "Order.arrival_station": %w`, err)}
		}
	}
	if _, ok := oc.mutation.IDNumber(); !ok {
		return &ValidationError{Name: "id_number", err: errors.New(`ent: missing required field "Order.id_number"`)}
	}
	if v, ok := oc.mutation.IDNumber(); ok {
		if err := order.IDNumberValidator(v); err != nil {
			return &ValidationError{Name: "id_number", err: fmt.Errorf(`ent: validator failed for field "Order.id_number": %w`, err)}
		}
	}
	if _, ok := oc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`ent: missing required field "Order.phone_number"`)}
	}
	if v, ok := oc.mutation.PhoneNumber(); ok {
		if err := order.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Order.phone_number": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Order.email"`)}
	}
	if v, ok := oc.mutation.Email(); ok {
		if err := order.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Order.email": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	if _, ok := oc.mutation.ErrorMessage(); !ok {
		return &ValidationError{Name: "error_message", err: errors.New(`ent: missing required field "Order.error_message"`)}
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
	}
	if _, ok := oc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Order.user"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.StartTime(); ok {
		_spec.SetField(order.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := oc.mutation.EndTime(); ok {
		_spec.SetField(order.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := oc.mutation.DepartureStation(); ok {
		_spec.SetField(order.FieldDepartureStation, field.TypeString, value)
		_node.DepartureStation = value
	}
	if value, ok := oc.mutation.ArrivalStation(); ok {
		_spec.SetField(order.FieldArrivalStation, field.TypeString, value)
		_node.ArrivalStation = value
	}
	if value, ok := oc.mutation.IDNumber(); ok {
		_spec.SetField(order.FieldIDNumber, field.TypeString, value)
		_node.IDNumber = value
	}
	if value, ok := oc.mutation.PhoneNumber(); ok {
		_spec.SetField(order.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := oc.mutation.Email(); ok {
		_spec.SetField(order.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.ErrorMessage(); ok {
		_spec.SetField(order.FieldErrorMessage, field.TypeString, value)
		_node.ErrorMessage = value
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := oc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ValidationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   order.ValidationTable,
			Columns: []string{order.ValidationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ordervalidation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	err      error
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
