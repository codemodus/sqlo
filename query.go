package sqlo

import (
	"context"
)

// PlainQry ...
type PlainQry struct {
	name string
	qry  string
	args []interface{}
}

// NewPlainQry ...
func (*SQLO) NewPlainQry(name, query string, args ...interface{}) *PlainQry {
	return &PlainQry{name, query, args}
}

// Query ...
func (qry *PlainQry) Query(ctx context.Context, q Queryable) error {
	_, err := q.ExecContext(ctx, qry.qry, qry.args...)
	return err
}

// NamedQry ...
type NamedQry struct {
	name string
	qry  string
	arg  interface{}
}

// NewNamedQry ...
func (*SQLO) NewNamedQry(name, query string, arg interface{}) *NamedQry {
	return &NamedQry{name, query, arg}
}

// Query ...
func (qry *NamedQry) Query(ctx context.Context, q Queryable) error {
	_, err := q.NamedExecContext(ctx, qry.qry, qry.arg)
	return err
}

// GetQry ...
type GetQry struct {
	name string
	dest interface{}
	qry  string
	args []interface{}
}

// NewGetQry ...
func (*SQLO) NewGetQry(name string, dest interface{}, query string, args ...interface{}) *GetQry {
	return &GetQry{name, dest, query, args}
}

// Query ...
func (qry *GetQry) Query(ctx context.Context, q Queryable) error {
	return q.GetContext(ctx, qry.dest, qry.qry, qry.args...)
}

// SelectQry ...
type SelectQry struct {
	name string
	dest interface{}
	qry  string
	args []interface{}
}

// NewSelectQry ...
func (*SQLO) NewSelectQry(name string, dest interface{}, query string, args ...interface{}) *SelectQry {
	return &SelectQry{name, dest, query, args}
}

// Query ...
func (qry *SelectQry) Query(ctx context.Context, q Queryable) error {
	return q.SelectContext(ctx, qry.dest, qry.qry, qry.args...)
}
