package sqlo

import (
	"context"
	"database/sql"
)

// Queryable ...
type Queryable interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// TxQueryable ...
type TxQueryable interface {
	Queryable
	Rollback() error
	Commit() error
}

// DB ...
type DB interface {
	SQLDB() *sql.DB
	Queryable
	BeginTxContext(ctx context.Context, opts *sql.TxOptions) (TxQueryable, error)
}

// Query ...
type Query interface {
	Send(ctx context.Context, q Queryable) error
}

// Hydratable ...
type Hydratable interface {
	IsHydrated() bool
}

// SQLO ...
type SQLO struct {
	db DB
}

// New ...
func New(db DB) (*SQLO, error) {
	return &SQLO{db: db}, nil
}
