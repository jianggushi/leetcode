package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestRectangleArea_01(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	expected := 10
	ans := largestRectangleArea(heights)
	assert.Equal(t, expected, ans)
}

func Test_largestRectangleArea_02(t *testing.T) {
	heights := []int{2}
	expected := 2
	ans := largestRectangleArea(heights)
	assert.Equal(t, expected, ans)
}

func Test_largestRectangleArea_03(t *testing.T) {
	heights := []int{2, 1}
	expected := 2
	ans := largestRectangleArea(heights)
	assert.Equal(t, expected, ans)
}

func Test_largestRectangleArea_04(t *testing.T) {
	heights := []int{1, 1}
	expected := 2
	ans := largestRectangleArea(heights)
	assert.Equal(t, expected, ans)
}

func Test_largestRectangleArea_05(t *testing.T) {
	heights := []int{2, 1, 2}
	expected := 3
	ans := largestRectangleArea(heights)
	assert.Equal(t, expected, ans)
}
