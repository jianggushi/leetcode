package solution

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers_1(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{5, 6}

	result := findDisappearedNumbers(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("excepted: %v, findDisappearedNumbers: %v", expected, result)
	}
}
