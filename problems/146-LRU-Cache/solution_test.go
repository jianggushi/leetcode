package solution

import "testing"

func Test_LRUCache_01(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	ans := cache.Get(1) // 返回  1
	if ans != 1 {
		t.Error(ans)
	}
	cache.Put(3, 3)    // 该操作会使得密钥 2 作废
	ans = cache.Get(2) // 返回 -1 (未找到)
	if ans != -1 {
		t.Error(ans)
	}
	cache.Put(4, 4)    // 该操作会使得密钥 1 作废
	ans = cache.Get(1) // 返回 -1 (未找到)
	if ans != -1 {
		t.Error(ans)
	}
	ans = cache.Get(3) // 返回  3
	if ans != 3 {
		t.Error(ans)
	}
	ans = cache.Get(4) // 返回  4
	if ans != 4 {
		t.Error(ans)
	}
}

func Test_LRUCache_02(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	ans := cache.Get(2) // 返回  1
	if ans != 1 {
		t.Error(ans)
	}
}
