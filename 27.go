package main

func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return right
}
