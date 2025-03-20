package main

func sortArrayByParity(nums []int) []int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 0 {
			if i != l {
				nums[l], nums[i] = nums[i], nums[l]
			}
			l++
		}
	}
	return nums
}
