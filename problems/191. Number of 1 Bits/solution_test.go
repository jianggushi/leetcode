package solution

import "testing"

func Test_hammingWeight_1(t *testing.T) {
	num := uint32(0x0b)
	expected := 3
	result := hammingWeight(num)
	if result != expected {
		t.Errorf("expected: %v, hammingWeight: %v", expected, result)
	}
}
