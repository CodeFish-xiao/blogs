package fuck_defer

import (
	"testing"
)

func TestName(t *testing.T) {
	e := &Entity{}
	e.WithOptions(
		FieldOne("field one's value"),
	).Do()
}
