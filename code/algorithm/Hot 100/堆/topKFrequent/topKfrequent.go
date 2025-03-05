package topKFrequent

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	//
	for _, num := range nums {
		occurrences[num]++
	}
	h := &MinHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, Data{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).(Data).Key
	}
	return ret
}

type Data struct {
	Key int
	Val int
}
type MinHeap []Data

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Data))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
