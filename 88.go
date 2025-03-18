package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for i > -1 || j > -1 {
		if i == -1 {
			nums1[i+j+1] = nums2[j]
			j--
		} else if j == -1 {
			nums1[i+j+1] = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			nums1[i+j+1] = nums2[j]
			j--
		} else {
			nums1[i+j+1] = nums1[i]
			i--
		}
	}
}
