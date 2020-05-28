package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate_01(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := true
	ans := containsDuplicate(nums)
	assert.Equal(t, expected, ans)
}

func Test_containsDuplicate_02(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := false
	ans := containsDuplicate(nums)
	assert.Equal(t, expected, ans)
}

func Test_containsDuplicate_03(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	expected := true
	ans := containsDuplicate(nums)
	assert.Equal(t, expected, ans)
}

func Test_containsDuplicate_04(t *testing.T) {
	nums := []int{}
	expected := false
	ans := containsDuplicate(nums)
	assert.Equal(t, expected, ans)
}

func Test_containsDuplicate_05(t *testing.T) {
	nums := []int{1}
	expected := false
	ans := containsDuplicate(nums)
	assert.Equal(t, expected, ans)
}
