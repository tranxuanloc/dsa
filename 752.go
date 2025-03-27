package main

func openLock(deadends []string, target string) int {
	cd := map[string]*struct{}{}
	start := "0000"
	cache := map[string]int{
		start: 0,
	}
	if target == start {
		return 0
	}
	for _, deadend := range deadends {
		if deadend == start {
			return -1
		}
		cd[deadend] = &struct{}{}
	}
	queue := make([]string, 0)
	queue = append(queue, "0000")
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		height := cache[cur]
		for i, n := range cur {
			nextState := [4]byte{cur[0], cur[1], cur[2], cur[3]}
			nextState[i] = byte(rotateForward(n))
			state1 := string(nextState[:])
			if state1 == target {
				return height + 1
			}
			if _, ok := cache[state1]; !ok && cd[state1] == nil {
				queue = append(queue, state1)
				cache[state1] = height + 1
			}
			nextState[i] = byte(rotateBackward(n))
			state2 := string(nextState[:])
			if state2 == target {
				return height + 1
			}
			if _, ok := cache[state2]; !ok && cd[state2] == nil {
				queue = append(queue, state2)
				cache[state2] = height + 1
			}
		}
	}
	return -1
}

func rotateBackward(n rune) rune {
	if n == '0' {
		return '9'
	}
	return n - 1
}

func rotateForward(n rune) rune {
	if n == '9' {
		return '0'
	}
	return n + 1
}
