package main

func removeDuplicates(nums []int) int {
	nextPos := 1
	for _, num := range nums {
		if num != nums[nextPos-1] {
			nums[nextPos] = num
			nextPos++
		}
	}
	return nextPos
}
