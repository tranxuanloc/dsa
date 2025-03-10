package main

func removeElement(nums []int, val int) int {
	nextPos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[nextPos] = nums[i]
			nextPos++
		}
	}
	return nextPos
}
