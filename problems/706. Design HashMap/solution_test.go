package solution

import "testing"

func Test_MyHashMap1(t *testing.T) {
	hashMap := Constructor()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	ans := hashMap.Get(1) // returns 1
	if ans != 1 {
		t.Error(ans)
	}
	ans = hashMap.Get(3) // returns -1 (not found)
	if ans != -1 {
		t.Error(ans)
	}
	hashMap.Put(2, 1)    // update the existing value
	ans = hashMap.Get(2) // returns 1
	if ans != 1 {
		t.Error(ans)
	}
	hashMap.Remove(2)    // remove the mapping for 2
	ans = hashMap.Get(2) // returns -1 (not found)
	if ans != -1 {
		t.Error(ans)
	}
}
