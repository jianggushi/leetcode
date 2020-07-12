package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	assert.Equal(t, -3, stack.GetMin())
	stack.Pop()
	assert.Equal(t, 0, stack.Top())
	assert.Equal(t, -2, stack.GetMin())
}
