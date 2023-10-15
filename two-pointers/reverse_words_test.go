package twopointers

import "testing"

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			"Reverse the words in a sentence", "sentence a in words the Reverse",
		},
		{
			"You are amazing", "amazing are You",
		},
		{
			"Hello    World", "World Hello",
		},
	}
	for _, tc := range testCases {
		if got := reverseWords(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
