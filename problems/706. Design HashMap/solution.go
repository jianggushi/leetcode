package solution

// MyHashMap .
type MyHashMap struct {
	e [][]int
}

// Constructor Initialize your data structure here
func Constructor() MyHashMap {
	return MyHashMap{
		e: make([][]int, 10000),
	}
}

// Put value will always be non-negative
func (h *MyHashMap) Put(key int, value int) {
	mkey := key % 10000
	if h.e[mkey] != nil {
		n := len(h.e[mkey])
		for i := 0; i < n; i += 2 {
			if h.e[mkey][i] == key {
				h.e[mkey][i+1] = value
				return
			}
		}
	}
	h.e[mkey] = append(h.e[mkey], []int{key, value}...)
}

// Get Returns the value to which the specified key is mapped, or -1 if h map contains no mapping for the key
func (h *MyHashMap) Get(key int) int {
	mkey := key % 10000
	if h.e[mkey] != nil {
		n := len(h.e[mkey])
		for i := 0; i < n; i += 2 {
			if h.e[mkey][i] == key {
				return h.e[mkey][i+1]
			}
		}
	}
	return -1
}

// Remove Removes the mapping of the specified value key if h map contains a mapping for the key
func (h *MyHashMap) Remove(key int) {
	mkey := key % 10000
	if h.e[mkey] != nil {
		n := len(h.e[mkey])
		for i := 0; i < n; i += 2 {
			if h.e[mkey][i] == key {
				h.e[mkey] = append(h.e[mkey][:i], h.e[mkey][i+2:]...)
				return
			}
		}
	}
}
