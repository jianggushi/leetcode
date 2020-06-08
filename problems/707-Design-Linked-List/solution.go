package solution

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	i := 0
	for p := this.next; p != nil; p = p.next {
		if i == index {
			return p.val
		}
		i++
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkedList{
		val:  val,
		next: this.next,
	}
	this.next = node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &MyLinkedList{
		val: val,
	}
	p := this
	for ; p.next != nil; p = p.next {
	}
	p.next = node
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	node := &MyLinkedList{
		val: val,
	}
	p1, p := this, this.next
	i := 0
	for p != nil {
		if i == index {
			break
		}
		p1 = p
		p = p.next
		i++
	}
	if i == index {
		node.next = p1.next
		p1.next = node
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	p1, p := this, this.next
	i := 0
	for p != nil {
		if i == index {
			p1.next = p.next
		}
		p1 = p
		p = p.next
		i++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
