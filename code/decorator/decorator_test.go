package decorator

import "testing"

func TestDecorator(t *testing.T) {
	hello := decorator(Hello)
	hello("Hello")
}
