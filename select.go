package sqlo

import (
	"context"
	"fmt"
)

// Select ...
func (s *SQLO) Select(ctx context.Context, scope string, dest Hydratable, qry string, args ...interface{}) error {
	if err := s.db.SelectContext(ctx, dest, qry, args...); err != nil {
		return fmt.Errorf("select (%s): %s", scope, err)
	}

	if !dest.IsHydrated() {
		return fmt.Errorf("not hydrated")
	}

	return nil
}
