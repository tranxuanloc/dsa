package main

import "math"

func findNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if int(math.Floor(math.Log10(float64(nums[i]))))&1 == 1 {
			count++
		}
	}
	return count
}
