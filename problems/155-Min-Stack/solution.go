package solution

import "math"

type MinStack struct {
	data    []int
	minData []int
	n       int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:    []int{},
		minData: []int{math.MaxInt64},
		n:       0,
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	this.minData = append(this.minData, min(x, this.minData[this.n]))
	this.n++
}

func (this *MinStack) Pop() {
	if this.n == 0 {
		panic("stack is empty")
	}
	this.data = this.data[:this.n-1]
	this.minData = this.minData[:this.n]
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
	return this.minData[this.n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
