package slidingwindow

import "testing"

func TestMinWindow(t *testing.T) {
	testCases := []struct {
		str1, str2 string
		want       string
	}{
		{"abcdebdde", "bde", "bcde"},
	}
	for _, tc := range testCases {
		if got := minWindow(tc.str1, tc.str2); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
