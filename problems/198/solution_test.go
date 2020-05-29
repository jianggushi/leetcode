package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rob_01(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4
	ans := rob(nums)
	assert.Equal(t, expected, ans)
}

func Test_rob_02(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	expected := 12
	ans := rob(nums)
	assert.Equal(t, expected, ans)
}

func Test_rob_03(t *testing.T) {
	nums := []int{}
	expected := 0
	ans := rob(nums)
	assert.Equal(t, expected, ans)
}

func Test_rob_04(t *testing.T) {
	nums := []int{1}
	expected := 1
	ans := rob(nums)
	assert.Equal(t, expected, ans)
}

func Test_rob_05(t *testing.T) {
	nums := []int{1, 2}
	expected := 2
	ans := rob(nums)
	assert.Equal(t, expected, ans)
}

func Test_rob_06(t *testing.T) {
	nums := []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}
	expected := 4173
	ans := rob(nums)
	assert.Equal(t, expected, ans)
}

func Test_rob_07(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	expected := 12
	ans := rob(nums)
	assert.Equal(t, expected, ans)
}
