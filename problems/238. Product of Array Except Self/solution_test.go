package solution

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}
	ans := productExceptSelf(nums)
	if !reflect.DeepEqual(ans, expected) {
		t.Errorf("expected: %v, productExceptSelf: %v", expected, ans)
	}
}
