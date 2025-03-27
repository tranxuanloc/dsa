package main

import "math"

func numSquares(n int) int {
	num := int(math.Sqrt(float64(n)))
	squareNums := make([]int, 0, num)
	for i := 1; i <= num; i++ {
		squareNums = append(squareNums, i*i)
	}
	visited := map[int]int{}
	visited[n] = 0
	queue := make([]int, 0)
	queue = append(queue, n)
	height := 0
	newSum := 0
	for len(queue) > 0 {
		sum := queue[0]
		queue = queue[1:]
		height = visited[sum]
	loop:
		for _, num := range squareNums {
			newSum = sum - num
			if newSum == 0 {
				return height + 1
			}
			if newSum < 0 {
				break loop
			}
			if _, ok := visited[newSum]; ok {
				continue loop
			}
			visited[newSum] = height + 1
			queue = append(queue, newSum)
		}
	}
	return 0
}
