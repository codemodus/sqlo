package sqlo

import (
	"context"
	"database/sql"
	"fmt"
)

// Exec ...
func (s *SQLO) Exec(ctx context.Context, qs ...Query) error {
	if len(qs) == 1 {
		return exec(ctx, s.db, qs[0])
	}

	tx, err := s.db.BeginTxContext(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	for k := range qs {
		if err = exec(ctx, tx, qs[k]); err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return rerr
			}

			return err
		}
	}

	return tx.Commit()
}

func exec(ctx context.Context, q Queryable, qry Query) error {
	if qry == nil {
		return nil
	}

	if err := qry.Send(ctx, q); err != nil {
		return fmt.Errorf("send (%s): %s", qry.Scope(), err)
	}

	return nil
}
