package solution

type LRUCache struct {
	cap  int
	len  int
	node *Node
}

type Node struct {
	key   int
	value int
	next  *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	var prev, curr *Node
	curr = this.node
	for curr != nil {
		if curr.key == key {
			// find it
			if prev != nil {
				prev.next = curr.next
				curr.next = this.node
				this.node = curr
			}
			return curr.value
		}
		prev = curr
		curr = curr.next
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	var prev, curr *Node
	curr = this.node
	for curr != nil {
		if curr.key == key {
			// find it
			curr.value = value
			if prev != nil {
				prev.next = curr.next
				curr.next = this.node
				this.node = curr
			}
			return
		}
		prev = curr
		curr = curr.next
	}
	// not find
	if this.len >= this.cap {
		// remove the last node
		prev, curr = nil, this.node
		for curr != nil && curr.next != nil {
			prev = curr
			curr = curr.next
		}
		if prev != nil {
			prev.next = nil
		}
	}
	// insert
	this.node = &Node{
		key:   key,
		value: value,
		next:  this.node,
	}
	this.len++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
