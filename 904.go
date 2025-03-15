package main

// 1,2,3,2,2
// i j
func totalFruit(fruits []int) int {
	var left, right, maxPicked int
	basket := map[int]int{}
	n := len(fruits)
	for ; right < n; right++ {
		if typeCount, exists := basket[fruits[right]]; exists || len(basket) < 2 {
			basket[fruits[right]] = typeCount + 1
			if count := right - left + 1; count > maxPicked {
				maxPicked = count
			}
			continue
		}
		basket[fruits[right]] = 1
		for len(basket) > 2 {
			typeCount := basket[fruits[left]] - 1
			basket[fruits[left]] = typeCount
			if typeCount == 0 {
				delete(basket, fruits[left])
			}
			left++
		}
	}
	return maxPicked
}
