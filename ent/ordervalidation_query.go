// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/ordervalidation"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// OrderValidationQuery is the builder for querying OrderValidation entities.
type OrderValidationQuery struct {
	config
	ctx        *QueryContext
	order      []ordervalidation.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderValidation
	withOrder  *OrderQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderValidationQuery builder.
func (ovq *OrderValidationQuery) Where(ps ...predicate.OrderValidation) *OrderValidationQuery {
	ovq.predicates = append(ovq.predicates, ps...)
	return ovq
}

// Limit the number of records to be returned by this query.
func (ovq *OrderValidationQuery) Limit(limit int) *OrderValidationQuery {
	ovq.ctx.Limit = &limit
	return ovq
}

// Offset to start from.
func (ovq *OrderValidationQuery) Offset(offset int) *OrderValidationQuery {
	ovq.ctx.Offset = &offset
	return ovq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ovq *OrderValidationQuery) Unique(unique bool) *OrderValidationQuery {
	ovq.ctx.Unique = &unique
	return ovq
}

// Order specifies how the records should be ordered.
func (ovq *OrderValidationQuery) Order(o ...ordervalidation.OrderOption) *OrderValidationQuery {
	ovq.order = append(ovq.order, o...)
	return ovq
}

// QueryOrder chains the current query on the "order" edge.
func (ovq *OrderValidationQuery) QueryOrder() *OrderQuery {
	query := (&OrderClient{config: ovq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ovq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ovq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ordervalidation.Table, ordervalidation.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ordervalidation.OrderTable, ordervalidation.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(ovq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderValidation entity from the query.
// Returns a *NotFoundError when no OrderValidation was found.
func (ovq *OrderValidationQuery) First(ctx context.Context) (*OrderValidation, error) {
	nodes, err := ovq.Limit(1).All(setContextOp(ctx, ovq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ordervalidation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ovq *OrderValidationQuery) FirstX(ctx context.Context) *OrderValidation {
	node, err := ovq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderValidation ID from the query.
// Returns a *NotFoundError when no OrderValidation ID was found.
func (ovq *OrderValidationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ovq.Limit(1).IDs(setContextOp(ctx, ovq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ordervalidation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ovq *OrderValidationQuery) FirstIDX(ctx context.Context) int {
	id, err := ovq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderValidation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderValidation entity is found.
// Returns a *NotFoundError when no OrderValidation entities are found.
func (ovq *OrderValidationQuery) Only(ctx context.Context) (*OrderValidation, error) {
	nodes, err := ovq.Limit(2).All(setContextOp(ctx, ovq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ordervalidation.Label}
	default:
		return nil, &NotSingularError{ordervalidation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ovq *OrderValidationQuery) OnlyX(ctx context.Context) *OrderValidation {
	node, err := ovq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderValidation ID in the query.
// Returns a *NotSingularError when more than one OrderValidation ID is found.
// Returns a *NotFoundError when no entities are found.
func (ovq *OrderValidationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ovq.Limit(2).IDs(setContextOp(ctx, ovq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ordervalidation.Label}
	default:
		err = &NotSingularError{ordervalidation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ovq *OrderValidationQuery) OnlyIDX(ctx context.Context) int {
	id, err := ovq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderValidations.
func (ovq *OrderValidationQuery) All(ctx context.Context) ([]*OrderValidation, error) {
	ctx = setContextOp(ctx, ovq.ctx, "All")
	if err := ovq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderValidation, *OrderValidationQuery]()
	return withInterceptors[[]*OrderValidation](ctx, ovq, qr, ovq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ovq *OrderValidationQuery) AllX(ctx context.Context) []*OrderValidation {
	nodes, err := ovq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderValidation IDs.
func (ovq *OrderValidationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ovq.ctx.Unique == nil && ovq.path != nil {
		ovq.Unique(true)
	}
	ctx = setContextOp(ctx, ovq.ctx, "IDs")
	if err = ovq.Select(ordervalidation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ovq *OrderValidationQuery) IDsX(ctx context.Context) []int {
	ids, err := ovq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ovq *OrderValidationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ovq.ctx, "Count")
	if err := ovq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ovq, querierCount[*OrderValidationQuery](), ovq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ovq *OrderValidationQuery) CountX(ctx context.Context) int {
	count, err := ovq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ovq *OrderValidationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ovq.ctx, "Exist")
	switch _, err := ovq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ovq *OrderValidationQuery) ExistX(ctx context.Context) bool {
	exist, err := ovq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderValidationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ovq *OrderValidationQuery) Clone() *OrderValidationQuery {
	if ovq == nil {
		return nil
	}
	return &OrderValidationQuery{
		config:     ovq.config,
		ctx:        ovq.ctx.Clone(),
		order:      append([]ordervalidation.OrderOption{}, ovq.order...),
		inters:     append([]Interceptor{}, ovq.inters...),
		predicates: append([]predicate.OrderValidation{}, ovq.predicates...),
		withOrder:  ovq.withOrder.Clone(),
		// clone intermediate query.
		sql:  ovq.sql.Clone(),
		path: ovq.path,
	}
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (ovq *OrderValidationQuery) WithOrder(opts ...func(*OrderQuery)) *OrderValidationQuery {
	query := (&OrderClient{config: ovq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ovq.withOrder = query
	return ovq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		JessionID string `json:"jession_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderValidation.Query().
//		GroupBy(ordervalidation.FieldJessionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ovq *OrderValidationQuery) GroupBy(field string, fields ...string) *OrderValidationGroupBy {
	ovq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderValidationGroupBy{build: ovq}
	grbuild.flds = &ovq.ctx.Fields
	grbuild.label = ordervalidation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		JessionID string `json:"jession_id,omitempty"`
//	}
//
//	client.OrderValidation.Query().
//		Select(ordervalidation.FieldJessionID).
//		Scan(ctx, &v)
func (ovq *OrderValidationQuery) Select(fields ...string) *OrderValidationSelect {
	ovq.ctx.Fields = append(ovq.ctx.Fields, fields...)
	sbuild := &OrderValidationSelect{OrderValidationQuery: ovq}
	sbuild.label = ordervalidation.Label
	sbuild.flds, sbuild.scan = &ovq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderValidationSelect configured with the given aggregations.
func (ovq *OrderValidationQuery) Aggregate(fns ...AggregateFunc) *OrderValidationSelect {
	return ovq.Select().Aggregate(fns...)
}

func (ovq *OrderValidationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ovq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ovq); err != nil {
				return err
			}
		}
	}
	for _, f := range ovq.ctx.Fields {
		if !ordervalidation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ovq.path != nil {
		prev, err := ovq.path(ctx)
		if err != nil {
			return err
		}
		ovq.sql = prev
	}
	return nil
}

func (ovq *OrderValidationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderValidation, error) {
	var (
		nodes       = []*OrderValidation{}
		withFKs     = ovq.withFKs
		_spec       = ovq.querySpec()
		loadedTypes = [1]bool{
			ovq.withOrder != nil,
		}
	)
	if ovq.withOrder != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ordervalidation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderValidation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderValidation{config: ovq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ovq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ovq.withOrder; query != nil {
		if err := ovq.loadOrder(ctx, query, nodes, nil,
			func(n *OrderValidation, e *Order) { n.Edges.Order = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ovq *OrderValidationQuery) loadOrder(ctx context.Context, query *OrderQuery, nodes []*OrderValidation, init func(*OrderValidation), assign func(*OrderValidation, *Order)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OrderValidation)
	for i := range nodes {
		if nodes[i].order_validation_order == nil {
			continue
		}
		fk := *nodes[i].order_validation_order
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(order.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_validation_order" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ovq *OrderValidationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ovq.querySpec()
	_spec.Node.Columns = ovq.ctx.Fields
	if len(ovq.ctx.Fields) > 0 {
		_spec.Unique = ovq.ctx.Unique != nil && *ovq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ovq.driver, _spec)
}

func (ovq *OrderValidationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ordervalidation.Table, ordervalidation.Columns, sqlgraph.NewFieldSpec(ordervalidation.FieldID, field.TypeInt))
	_spec.From = ovq.sql
	if unique := ovq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ovq.path != nil {
		_spec.Unique = true
	}
	if fields := ovq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordervalidation.FieldID)
		for i := range fields {
			if fields[i] != ordervalidation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ovq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ovq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ovq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ovq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ovq *OrderValidationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ovq.driver.Dialect())
	t1 := builder.Table(ordervalidation.Table)
	columns := ovq.ctx.Fields
	if len(columns) == 0 {
		columns = ordervalidation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ovq.sql != nil {
		selector = ovq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ovq.ctx.Unique != nil && *ovq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ovq.predicates {
		p(selector)
	}
	for _, p := range ovq.order {
		p(selector)
	}
	if offset := ovq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ovq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderValidationGroupBy is the group-by builder for OrderValidation entities.
type OrderValidationGroupBy struct {
	selector
	build *OrderValidationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ovgb *OrderValidationGroupBy) Aggregate(fns ...AggregateFunc) *OrderValidationGroupBy {
	ovgb.fns = append(ovgb.fns, fns...)
	return ovgb
}

// Scan applies the selector query and scans the result into the given value.
func (ovgb *OrderValidationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ovgb.build.ctx, "GroupBy")
	if err := ovgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderValidationQuery, *OrderValidationGroupBy](ctx, ovgb.build, ovgb, ovgb.build.inters, v)
}

func (ovgb *OrderValidationGroupBy) sqlScan(ctx context.Context, root *OrderValidationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ovgb.fns))
	for _, fn := range ovgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ovgb.flds)+len(ovgb.fns))
		for _, f := range *ovgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ovgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ovgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderValidationSelect is the builder for selecting fields of OrderValidation entities.
type OrderValidationSelect struct {
	*OrderValidationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ovs *OrderValidationSelect) Aggregate(fns ...AggregateFunc) *OrderValidationSelect {
	ovs.fns = append(ovs.fns, fns...)
	return ovs
}

// Scan applies the selector query and scans the result into the given value.
func (ovs *OrderValidationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ovs.ctx, "Select")
	if err := ovs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderValidationQuery, *OrderValidationSelect](ctx, ovs.OrderValidationQuery, ovs, ovs.inters, v)
}

func (ovs *OrderValidationSelect) sqlScan(ctx context.Context, root *OrderValidationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ovs.fns))
	for _, fn := range ovs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ovs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ovs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
