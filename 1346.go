package main

func checkIfExist(arr []int) bool {
	hm := map[int]struct{}{}
	for _, num := range arr {
		if _, exists := hm[num*2]; exists {
			return true
		}
		if num%2 == 0 {
			if _, exists := hm[num/2]; exists {
				return true
			}
		}
		hm[num] = struct{}{}
	}
	return false
}
