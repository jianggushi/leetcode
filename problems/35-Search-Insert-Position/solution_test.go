package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert_01(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	want := 2
	got := searchInsert(nums, target)
	assert.Equal(t, want, got)
}

func Test_searchInsert_02(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	want := 1
	got := searchInsert(nums, target)
	assert.Equal(t, want, got)
}

func Test_searchInsert_03(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	want := 4
	got := searchInsert(nums, target)
	assert.Equal(t, want, got)
}

func Test_searchInsert_04(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 0
	want := 0
	got := searchInsert(nums, target)
	assert.Equal(t, want, got)
}

func Test_searchInsert_05(t *testing.T) {
	nums := []int{1}
	target := 0
	want := 0
	got := searchInsert(nums, target)
	assert.Equal(t, want, got)
}
