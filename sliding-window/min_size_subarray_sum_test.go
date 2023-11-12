package slidingwindow

import "testing"

func TestMinimumSizeSubarraySum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 7, 1, 8}, 9, 2},
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
	}
	for _, tc := range testCases {
		if got := MinimumSizeSubarraySum(tc.nums, tc.target); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}

}
