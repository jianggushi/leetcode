package solution

import "math"

type MinStack struct {
	data []int
	n    int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0, 8),
		n:    0,
		min:  math.MaxInt64,
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	this.n++
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if this.n == 0 {
		panic("stack is empty")
	}
	x := this.data[this.n-1]
	// pop 的栈顶元素为最小值，更新最小值
	if x == this.min {
		this.min = math.MaxInt64
		for i := 0; i < this.n-1; i++ {
			if this.data[i] < this.min {
				this.min = this.data[i]
			}
		}
	}
	this.data = this.data[:this.n-1]
	this.n--
}

func (this *MinStack) Top() int {
	if this.n == 0 {
		panic("stack is empty")
	}
	return this.data[this.n-1]
}

func (this *MinStack) GetMin() int {
	if this.n == 0 {
		panic("stack is empty")
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
