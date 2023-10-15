package twopointers

import "sort"

func findSumOfThree(nums []int, target int) bool {
	if len(nums) < 3 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i <= len(nums)-3; i++ {
		if twoSum(i, i+1, len(nums)-1, target, nums) {
			return true
		}
	}
	return false
}

func twoSum(left, mid, right, target int, nums []int) bool {
	for mid < right {
		sum := nums[left] + nums[mid] + nums[right]
		if sum == target {
			return true
		}
		if sum > target {
			right--
		} else {
			mid++
		}
	}
	return false
}
