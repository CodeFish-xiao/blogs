package lru

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	mu     sync.Mutex
	maxLen int
	// list
	l *list.List
	// map：同时是O(1)判断器
	m map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxLen: capacity,
		l:      list.New(),
		m:      make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()
	if ele, ok := this.m[key]; ok {
		this.l.MoveToFront(ele)
		kv := ele.Value.(int)
		return kv
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if ele, ok := this.m[key]; ok {
		this.l.MoveToFront(ele)
		ele.Value = value
	} else {
		if this.l.Len() >= this.maxLen {
			back := this.l.Back()
			delete(this.m, back.Value.(int))
			this.l.Remove(back)
		}
		el := this.l.PushFront(value)
		this.m[key] = el
	}
}
