package solution

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue_1(t *testing.T) {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	expected := [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}
	result := reconstructQueue(people)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, reconstructQueue: %v", expected, result)
	}
}
