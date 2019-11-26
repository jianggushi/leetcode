package solution

import (
	"reflect"
	"testing"
)

func Test_merge_1(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	ans := merge(intervals)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, merge: %v", expected, ans)
	}
}

func Test_merge_2(t *testing.T) {
	intervals := [][]int{{1, 4}, {4, 5}}
	expected := [][]int{{1, 5}}
	ans := merge(intervals)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, merge: %v", expected, ans)
	}
}

func Test_merge_3(t *testing.T) {
	intervals := [][]int{{1, 4}, {0, 4}}
	expected := [][]int{{0, 4}}
	ans := merge(intervals)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, merge: %v", expected, ans)
	}
}
