package two_sum

import (
	"testing"
)

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		results []int
		input   []int
		target  int
	}{
		{[]int{1, 2}, []int{3, 5, 4, 8, 1}, 9},
		{[]int{1, 4}, []int{3, 5, 4, 8, 1}, 6},
		{[]int{0, 2}, []int{3, 5, 4, 8, 1}, 7},
		{nil, []int{3, 5, 4, 8, 1}, 88},
	}

	for _, test := range tests {
		if got := TwoSum(test.input, test.target); !isEqual(got, test.results) {
			t.Errorf("TwoSum(%v,%v) got:%v want:%v", test.input, test.target, got, test.results)
		}
	}

}
