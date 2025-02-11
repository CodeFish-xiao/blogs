package lru

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3)
	obj.Get(2)
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	obj.Get(4)
}
