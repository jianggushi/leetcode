package solution

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures_1(t *testing.T) {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	result := dailyTemperatures(T)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, dailyTemperatures: %v", expected, result)
	}
}

func Test_dailyTemperatures_2(t *testing.T) {
	T := []int{76, 73}
	expected := []int{0, 0}
	result := dailyTemperatures(T)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, dailyTemperatures: %v", expected, result)
	}
}
