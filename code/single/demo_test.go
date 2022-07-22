package single

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"testing"
)

func TestSingleFlightDemo(t *testing.T) {
	var g singleflight.Group
	var wg sync.WaitGroup
	const n = 5
	key := "something url(db)"
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			v, _, shared := g.Do(key, func() (interface{}, error) {
				ret, err := find(context.Background(), key)
				return ret, err
			})
			fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
		}(i)
	}
	wg.Wait()
}
