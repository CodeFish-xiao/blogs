package lru

import "container/list"

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element
}
type kv struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		this.list.MoveToFront(ele)
		return ele.Value.(kv).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.cache[key]; ok {
		if e.Prev() != nil {
			this.list.MoveToFront(e)
		}
		e.Value = kv{key, value}
	} else {
		if this.list.Len() >= this.capacity {
			e = this.list.Back()
			this.list.Remove(e)
			delete(this.cache, e.Value.(kv).key)
		}
		e = this.list.PushFront(kv{key, value})
		this.cache[key] = e
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
