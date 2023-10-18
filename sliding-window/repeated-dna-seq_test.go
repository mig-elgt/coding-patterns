package slidingwindow

import (
	"reflect"
	"testing"
)

func TestFindRepeatedSequenses(t *testing.T) {
	testCases := []struct {
		in   string
		k    int
		want Set
	}{
		{
			"ATATATATATATATAT",
			6,
			Set{
				hashMap: map[interface{}]bool{
					"ATATAT": true,
					"TATATA": true,
				},
			},
		},
		{
			"GGGGGGGGGGGGGGGGGGGGGGGGG",
			9,
			Set{
				hashMap: map[interface{}]bool{
					"GGGGGGGGG": true,
				},
			},
		},
	}
	for _, tc := range testCases {
		if got := findRepeatedSequences(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
