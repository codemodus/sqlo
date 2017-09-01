package sqlo

import (
	"context"
)

// PlainQry ...
type PlainQry struct {
	scp  string
	qry  string
	args []interface{}
}

// NewPlainQry ...
func (*SQLO) NewPlainQry(scope, query string, args ...interface{}) *PlainQry {
	return &PlainQry{scope, query, args}
}

// Send ...
func (qry *PlainQry) Send(ctx context.Context, q Queryable) error {
	_, err := q.ExecContext(ctx, qry.qry, qry.args...)
	return err
}

// Scope ...
func (qry *PlainQry) Scope() string {
	return qry.scp
}

// NamedQry ...
type NamedQry struct {
	scp string
	qry string
	arg interface{}
}

// NewNamedQry ...
func (*SQLO) NewNamedQry(scope, query string, arg interface{}) *NamedQry {
	return &NamedQry{scope, query, arg}
}

// Send ...
func (qry *NamedQry) Send(ctx context.Context, q Queryable) error {
	_, err := q.NamedExecContext(ctx, qry.qry, qry.arg)
	return err
}

// Scope ...
func (qry *NamedQry) Scope() string {
	return qry.scp
}

// GetQry ...
type GetQry struct {
	scp  string
	dest interface{}
	qry  string
	args []interface{}
}

// NewGetQry ...
func (*SQLO) NewGetQry(scope string, dest interface{}, query string, args ...interface{}) *GetQry {
	return &GetQry{scope, dest, query, args}
}

// Send ...
func (qry *GetQry) Send(ctx context.Context, q Queryable) error {
	return q.GetContext(ctx, qry.dest, qry.qry, qry.args...)
}

// Scope ...
func (qry *GetQry) Scope() string {
	return qry.scp
}

// SelectQry ...
type SelectQry struct {
	scp  string
	dest interface{}
	qry  string
	args []interface{}
}

// NewSelectQry ...
func (*SQLO) NewSelectQry(scope string, dest interface{}, query string, args ...interface{}) *SelectQry {
	return &SelectQry{scope, dest, query, args}
}

// Send ...
func (qry *SelectQry) Send(ctx context.Context, q Queryable) error {
	return q.SelectContext(ctx, qry.dest, qry.qry, qry.args...)
}

// Scope ...
func (qry *SelectQry) Scope() string {
	return qry.scp
}
