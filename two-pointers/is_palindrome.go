package twopointers

func isPalindrome(inputString string) bool {
	left, right := 0, len(inputString)-1
	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}
	return true
}
