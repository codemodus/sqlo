package sqlo

import (
	"context"
	"database/sql"
)

// SQLO ...
type SQLO struct {
	*sql.DB
}

// New ...
func New(db *sql.DB) *SQLO {
	return &SQLO{db}
}

// Scope ...
func (db *SQLO) Scope() Scope {
	return Scope{SQLO: db}
}

// Begin ...
func (db *SQLO) Begin() Tx {
	t := Tx{errs: &errs{}}
	t.Tx, t.tx = db.DB.Begin()
	return t
}

// BeginTx ...
func (db *SQLO) BeginTx(ctx context.Context, opts *sql.TxOptions) Tx {
	t := Tx{errs: &errs{}}
	t.Tx, t.tx = db.DB.BeginTx(ctx, opts)
	return t
}

//

// Scope ...
type Scope struct {
	*SQLO
	*errs
}

// Prepare ...
func (scx *Scope) Prepare(query string) Stmt {
	s := Stmt{errs: scx.errs}
	if s.hasError() {
		return s
	}
	s.Stmt, s.st = scx.DB.Prepare(query)
	return s
}

// Begin ...
func (scx *Scope) Begin() Tx {
	t := Tx{errs: scx.errs}
	if t.hasError() {
		return t
	}
	t.Tx, t.tx = scx.DB.Begin()
	return t
}

// BeginTx ...
func (scx *Scope) BeginTx(ctx context.Context, opts *sql.TxOptions) Tx {
	t := Tx{errs: scx.errs}
	if t.hasError() {
		return t
	}
	t.Tx, t.tx = scx.DB.BeginTx(ctx, opts)
	return t
}

//

type errs struct {
	tx, st, ro, sc error
}

func (es *errs) hasError() bool {
	return es != nil &&
		(es.tx != nil || es.st != nil || es.ro != nil || es.sc != nil)
}

func (es *errs) Err() error {
	if es == nil {
		return nil
	}
	switch {
	case es.tx != nil:
		return es.tx
	case es.st != nil:
		return es.st
	case es.ro != nil:
		return es.ro
	case es.sc != nil:
		return es.sc
	default:
		return nil
	}
}

//

// Tx ...
type Tx struct {
	Tx *sql.Tx
	*errs
}

// Rollback  ...
func (t *Tx) Rollback() error {
	return t.Tx.Rollback()
}

// Commit ...
func (t *Tx) Commit() error {
	if t.hasError() {
		return t.Err()
	}
	t.tx = t.Tx.Commit()
	return t.tx
}

// Prepare ...
func (t *Tx) Prepare(query string) Stmt {
	s := Stmt{errs: t.errs}
	if s.hasError() {
		return s
	}
	s.Stmt, s.st = t.Tx.Prepare(query)
	return s
}

//

// Stmt ...
type Stmt struct {
	Stmt *sql.Stmt
	*errs
}

// Close ...
func (s *Stmt) Close() error {
	return s.Stmt.Close()
}

// ExecContext ...
func (s *Stmt) ExecContext(ctx context.Context, args ...interface{}) Result {
	r := oResult{errs: s.errs}
	if r.hasError() {
		return r
	}
	r.Result, r.ro = s.Stmt.ExecContext(ctx, args...)
	return r
}

// QueryContext ...
func (s *Stmt) QueryContext(ctx context.Context, args ...interface{}) Rows {
	rs := Rows{errs: s.errs}
	if rs.hasError() {
		return rs
	}
	rs.Rows, rs.ro = s.Stmt.QueryContext(ctx, args...)
	return rs
}

// QueryRowContext ...
func (s *Stmt) QueryRowContext(ctx context.Context) Row {
	r := Row{errs: s.errs}
	if r.hasError() {
		return r
	}
	r.Row = s.Stmt.QueryRowContext(ctx)
	return r
}

//

// Result ...
type Result interface {
	LastInsertId() int64
	RowsAffected() int64
}

type oResult struct {
	sql.Result
	*errs
}

func (r oResult) LastInsertId() int64 {
	if r.hasError() {
		return 0
	}
	n, err := r.Result.LastInsertId()
	r.ro = err
	return n

}

func (r oResult) RowsAffected() int64 {
	if r.hasError() {
		return 0
	}
	n, err := r.Result.RowsAffected()
	r.ro = err
	return n
}

//

// Row ...
type Row struct {
	Row *sql.Row
	*errs
}

// Scan ...
func (r *Row) Scan(dest ...interface{}) {
	if r.hasError() {
		return
	}
	r.sc = r.Row.Scan(dest...)
}

//

// Rows ...
type Rows struct {
	Rows *sql.Rows
	*errs
}

// Close ...
func (rs *Rows) Close() error {
	return rs.Rows.Close()
}

// Next ...
func (rs *Rows) Next() bool {
	if rs.hasError() {
		return false
	}
	return rs.Rows.Next()
}

// Scan ...
func (rs *Rows) Scan(dest ...interface{}) {
	if rs.hasError() {
		return
	}
	rs.sc = rs.Rows.Scan(dest...)
}
