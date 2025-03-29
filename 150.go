package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	t := -1
	funcs := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	answer := 0
	for _, token := range tokens {
		fun := funcs[token]
		if fun != nil {
			answer = fun(stack[t-1], stack[t])
			stack[t-1] = answer
			stack = stack[:t]
			t--
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
			t++
		}
	}
	return stack[0]
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}
