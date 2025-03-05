package heap_demo

type minHeap []struct {
	key int
	val int
}

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool {
	// 如果是最大堆则h[i].val > h[j].val
	return h[i].val < h[j].val
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(struct {
		key int
		val int
	}))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
