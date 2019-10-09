package solution

type MinStack struct {
	elems []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		elems: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.elems) == 0 {
		this.min = x
	}
	if x < this.min {
		this.min = x
	}
	this.elems = append(this.elems, x)
}

func (this *MinStack) Pop() {
	n := len(this.elems)
	if n == 0 {
		panic("stack empty")
	}
	x := this.elems[n-1]
	// pop 的栈顶元素为最小值，更新最小值
	if x == this.min {
		this.min = this.elems[0]
		for i := 1; i < len(this.elems)-1; i++ {
			if this.elems[i] < this.min {
				this.min = this.elems[i]
			}
		}
	}
	this.elems = this.elems[:n-1]
}

func (this *MinStack) Top() int {
	n := len(this.elems)
	if n == 0 {
		panic("stack empty")
	}
	return this.elems[n-1]
}

func (this *MinStack) GetMin() int {
	n := len(this.elems)
	if n == 0 {
		panic("stack empty")
	}
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
