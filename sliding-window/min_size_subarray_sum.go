package slidingwindow

import (
	"fmt"
	"math"
)

func MinimumSizeSubarraySum(nums []int, target int) int {
	var left, right, currSum int
	minSize := math.MaxInt
	for right < len(nums) {
		currSum += nums[right]
		for left <= right && currSum >= target {
			size := (right - left) + 1
			if size < minSize {
				minSize = size
			}
			currSum -= nums[left]
			left++
			fmt.Println(currSum, left, right)
		}
		right++
	}
	if minSize != math.MaxInt {
		return minSize
	}
	return 0
}
