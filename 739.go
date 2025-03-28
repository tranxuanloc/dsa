package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := make([]int, 0, n)
	t := 0
	stack = append(stack, t)
	for i := 1; i < n; i++ {
	loop:
		for j := t; j > -1; j-- {
			// fmt.Println(i, j, t)
			if temperatures[i] <= temperatures[stack[j]] {
				break loop
			}
			answer[stack[j]] = i - stack[j]
			stack = stack[:j]
			t--
			// fmt.Println(stack, i, j, t)
		}
		stack = append(stack, i)
		t++
		// fmt.Println(stack, i, t)
	}
	return answer
}
