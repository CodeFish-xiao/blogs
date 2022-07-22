package single

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"testing"
	"time"
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
func TestSingleFlightDemoDoChan(t *testing.T) {
	var g singleflight.Group
	var wg sync.WaitGroup
	const n = 5
	key := "something url(db)"
	for i := 0; i < n; i++ {
		wg.Add(1)
		i2 := i
		go func(j int) {
			defer wg.Done()
			ch := g.DoChan(key, func() (interface{}, error) {
				ret, err := find(context.Background(), key)
				return ret, err
			})
			// Create our timeout
			timeout := time.After(time.Duration(500*i2) * time.Millisecond)

			var ret singleflight.Result
			select {
			case <-timeout: // Timeout elapsed
				fmt.Println("Timeout")
				return
			case ret = <-ch: // Received result from channel
				fmt.Printf("index: %d, val: %v, shared: %v\n", j, ret.Val, ret.Shared)
			}
		}(i)
	}
	wg.Wait()

}
