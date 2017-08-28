package sqlo

import (
	"context"
	"fmt"
)

// Select ...
func (s *SQLO) Select(ctx context.Context, dest Hydratable, qry string, args ...interface{}) error {
	if err := s.db.SelectContext(ctx, dest, qry, args...); err != nil {
		return err
	}

	if !dest.IsHydrated() {
		return fmt.Errorf("not hydrated")
	}

	return nil
}
