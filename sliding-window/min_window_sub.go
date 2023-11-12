package slidingwindow

import "math"

func minWindow(str1 string, str2 string) string {
	minWindow := math.MaxInt
	var res string
	var start, end int
	var i int
	for end < len(str1) {
		if i == len(str2)-1 {
			k := end
			for i >= 0 {
				if str1[k] == str2[i] {
					i--
				}
				k--
			}
			i = 0
			start = k + 1
			if minWindow > (end - start) {
				minWindow = end - start
				res = str1[start:end]
			}
			end = start + 1
			continue
		}
		if str1[end] == str2[i] {
			i++
		}
		end++
	}
	return res
}
