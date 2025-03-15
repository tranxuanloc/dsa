package main

func totalFruit(fruits []int) int {
	baskets := make([]int, 0)
	fruitTypes := make([]int, 0)
	n := len(fruits)
	currentFruit := fruits[0]
	currentBasket := 0
	for j := 0; j < n; j++ {
		if fruits[j] == currentFruit {
			currentBasket++
		} else {
			baskets = append(baskets, currentBasket)
			fruitTypes = append(fruitTypes, currentFruit)
			currentBasket = 1
			currentFruit = fruits[j]
		}
	}
	baskets = append(baskets, currentBasket)
	fruitTypes = append(fruitTypes, currentFruit)
	// fmt.Println(baskets)
	// fmt.Println(fruitTypes)
	n = len(fruitTypes)
	max := 0
	for k := 0; k < n; k++ {
		bag := map[int]int{}
		for i := k; i < n; i++ {
			if count, exists := bag[fruitTypes[i]]; exists {
				bag[fruitTypes[i]] = count + baskets[i]
			} else if len(bag) < 2 {
				bag[fruitTypes[i]] = baskets[i]
			} else {
				break
			}
		}
		total := 0
		for _, v := range bag {
			total += v
		}
		if total > max {
			max = total
		}
	}
	return max
}
