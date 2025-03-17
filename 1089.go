package main

func duplicateZeros(arr []int) {
	n := len(arr)
	j := 0
	queue := make([]int, 0, n)
	zero := false
	for i := 0; i < n; i++ {
		if zero {
			queue = append(queue, arr[i])
			arr[i] = 0
			zero = false
			continue
		}
		if j < len(queue) {
			queue = append(queue, arr[i])
			arr[i] = queue[j]
			j++
		}
		if arr[i] == 0 {
			zero = true
		}
	}
}
