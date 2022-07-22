package single

import (
	"context"
	"fmt"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	return Result(fmt.Sprintf("result for %q", query)), nil
}
