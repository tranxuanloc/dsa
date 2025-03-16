package main

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	count := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 1 {
			count++
		} else {
			res = max(res, count)
			count = 0
		}
	}
	return max(res, count)
}
