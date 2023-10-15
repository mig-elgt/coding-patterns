package twopointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		in   string
		want bool
	}{
		{"kayak", true},
		{"hello", false},
		{"RACEACAR", false},
		{"A", true},
		{"ABCDABCD", false},
	}
	for _, tc := range testCases {
		if got := isPalindrome(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
