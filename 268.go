package main

func missingNumber(nums []int) int {
	n := len(nums)
	result := nums[0]
	for i := 1; i < n; i++ {
		result ^= nums[i] ^ i
	}
	return result ^ n
}
