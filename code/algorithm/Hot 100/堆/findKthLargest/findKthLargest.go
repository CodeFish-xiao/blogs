package findKthLargest

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, Data{Key: nums[i], Val: nums[i]})
	}
	res := 0
	for i := 0; i < k; i++ {
		res = heap.Pop(h).(Data).Val
	}
	return res
}

type Data struct {
	Key int
	Val int
}

type MaxHeap []Data

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Val > h[j].Val
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Data))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
