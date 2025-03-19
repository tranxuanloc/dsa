package main

func validMountainArray(arr []int) bool {
	n := len(arr) - 1
	if n < 2 {
		return false
	}
	i := 0
	for i < n && arr[i] < arr[i+1] {
		i++
	}
	if i == 0 || i == n {
		return false
	}
	for ; i < n; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}
	return true
}
