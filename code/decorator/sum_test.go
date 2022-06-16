package decorator

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)

	fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}
