package heap_demo

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &minHeap{}
	heap.Init(h)
}
