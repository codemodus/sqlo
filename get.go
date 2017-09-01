package sqlo

import (
	"context"
	"fmt"
)

// Get ...
func (s *SQLO) Get(ctx context.Context, scope string, dest Hydratable, qry string, args ...interface{}) error {
	if err := s.db.GetContext(ctx, dest, qry, args...); err != nil {
		return fmt.Errorf("get (%s): %s", scope, err)
	}

	if !dest.IsHydrated() {
		return fmt.Errorf("not hydrated")
	}

	return nil
}
