package twopointers

import "testing"

func TestSumOfThree(t *testing.T) {
	testCase := []struct {
		in     []int
		target int
		want   bool
	}{
		{[]int{1, -1, 0}, -1, false},
		{[]int{3, 7, 1, 2, 8, 4, 5}, 10, true},
	}
	for _, tc := range testCase {
		if got := findSumOfThree(tc.in, tc.target); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
