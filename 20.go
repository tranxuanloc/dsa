package main

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	bracket := map[rune]rune{
		41:  40,
		93:  91,
		125: 123,
	}
	n := -1
	for _, r := range s {
		if c, ok := bracket[r]; ok {
			if n < 0 {
				return false
			}
			top := stack[n]
			if top != c {
				return false
			}
			stack = stack[:n]
			n--
		} else {
			stack = append(stack, r)
			n++
		}
	}
	return n < 0
}
