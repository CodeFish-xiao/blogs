package demo

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var list []int
	for i := 0; i < 5; i++ {
		list = append(list, i)
	}
	fmt.Println(list)
	var sbList []*int
	for i := 0; i < 5; i++ {
		sbList = append(sbList, &i)
	}
	fmt.Println(sbList)
}
