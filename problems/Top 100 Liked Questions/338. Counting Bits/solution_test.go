package solution

import (
	"reflect"
	"testing"
)

func Test_countBits_1(t *testing.T) {
	num := 2
	expected := []int{0, 1, 1}
	result := countBits(num)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, countBits: %v", expected, result)
	}
}

func Test_countBits_2(t *testing.T) {
	num := 5
	expected := []int{0, 1, 1, 2, 1, 2}
	result := countBits(num)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, countBits: %v", expected, result)
	}
}
