package sqlo

import (
	"context"
)

// PlainQry ...
type PlainQry struct {
	qry  string
	args []interface{}
}

// NewPlainQry ...
func (*SQLO) NewPlainQry(query string, args ...interface{}) *PlainQry {
	return &PlainQry{query, args}
}

// Send ...
func (qry *PlainQry) Send(ctx context.Context, q Queryable) error {
	_, err := q.ExecContext(ctx, qry.qry, qry.args...)
	return err
}

// NamedQry ...
type NamedQry struct {
	qry string
	arg interface{}
}

// NewNamedQry ...
func (*SQLO) NewNamedQry(query string, arg interface{}) *NamedQry {
	return &NamedQry{query, arg}
}

// Send ...
func (qry *NamedQry) Send(ctx context.Context, q Queryable) error {
	_, err := q.NamedExecContext(ctx, qry.qry, qry.arg)
	return err
}

// GetQry ...
type GetQry struct {
	dest interface{}
	qry  string
	args []interface{}
}

// NewGetQry ...
func (*SQLO) NewGetQry(dest interface{}, query string, args ...interface{}) *GetQry {
	return &GetQry{dest, query, args}
}

// Send ...
func (qry *GetQry) Send(ctx context.Context, q Queryable) error {
	return q.GetContext(ctx, qry.dest, qry.qry, qry.args...)
}

// SelectQry ...
type SelectQry struct {
	dest interface{}
	qry  string
	args []interface{}
}

// NewSelectQry ...
func (*SQLO) NewSelectQry(dest interface{}, query string, args ...interface{}) *SelectQry {
	return &SelectQry{dest, query, args}
}

// Send ...
func (qry *SelectQry) Send(ctx context.Context, q Queryable) error {
	return q.SelectContext(ctx, qry.dest, qry.qry, qry.args...)
}
