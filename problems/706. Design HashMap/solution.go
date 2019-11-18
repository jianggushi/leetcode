package solution

type MyHashMap struct {
	e [][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		e: make([][]int, 10000),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	mkey := key % 10000
	if this.e[mkey] != nil {
		n := len(this.e[mkey])
		for i := 0; i < n; i += 2 {
			if this.e[mkey][i] == key {
				this.e[mkey][i+1] = value
				return
			}
		}
	}
	this.e[mkey] = append(this.e[mkey], []int{key, value}...)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	mkey := key % 10000
	if this.e[mkey] != nil {
		n := len(this.e[mkey])
		for i := 0; i < n; i += 2 {
			if this.e[mkey][i] == key {
				return this.e[mkey][i+1]
			}
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	mkey := key % 10000
	if this.e[mkey] != nil {
		n := len(this.e[mkey])
		for i := 0; i < n; i += 2 {
			if this.e[mkey][i] == key {
				this.e[mkey] = append(this.e[mkey][:i], this.e[mkey][i+2:]...)
				return
			}
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
