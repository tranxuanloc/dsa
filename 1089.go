package main

func duplicateZeros(arr []int) {
	n := len(arr) - 1
	count := 0
	for i := 0; i <= n-count; i++ {
		if arr[i] == 0 {
			if i == n-count {
				arr[n] = 0
				n--
			} else {
				count++
			}
		}
	}
	for i := n - count; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+count] = 0
			count--
			arr[i+count] = 0
		} else {
			arr[i+count] = arr[i]
		}
	}
}
