package sqlo

import (
	"context"
)

// In ...
func (s *SQLO) In(ctx context.Context, dest Hydratable, qry string, args ...interface{}) error {
	newQry, newArgs, err := in(qry, args...)
	if err != nil {
		return err
	}

	return s.Select(ctx, dest, newQry, newArgs...)
}
