package single

import (
	"context"
	"fmt"
	"time"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	time.Sleep(time.Second * 1)
	return Result(fmt.Sprintf("result for %q", query)), nil
}
