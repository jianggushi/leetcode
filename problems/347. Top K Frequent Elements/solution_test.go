package solution

import (
	"reflect"
	"testing"
)

func Test_topKFrequent_1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}
	ans := topKFrequent(nums, k)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, topKFrequent: %v", expected, ans)
	}
}

func Test_topKFrequent_2(t *testing.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}
	ans := topKFrequent(nums, k)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, topKFrequent: %v", expected, ans)
	}
}
