package lru

import "testing"

func TestLru(t *testing.T) {
	//输入
	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	cache := Constructor(2)
	cache.Put(1, 0)
	cache.Put(2, 2)
	if cache.Get(1) != 0 {
		t.Fatal("get 1 error")
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Fatal("get 2 error")
	}
	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Fatal("get 1 error")
	}
	if cache.Get(3) != 3 {
		t.Fatal("get 3 error")
	}
	if cache.Get(4) != 4 {
		t.Fatal("get 4 error")
	}

}
