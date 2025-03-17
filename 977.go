package main

import "math"

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, len(nums))
	for i := right; i > -1; i-- {
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
			res[i] = nums[right] * nums[right]
			right--
		} else {
			res[i] = nums[left] * nums[left]
			left++
		}
	}
	return res
}
