package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	i := 0
	j := n - 1
	var pivot int
	for i <= j {
		pivot = (i + j) / 2
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] > target {
			j = pivot - 1
		} else {
			i = pivot + 1
		}
	}
	if i > j {
		return i
	}
	return j
}
