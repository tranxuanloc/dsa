package main

func mySqrt(x int) int {
	for i := 1; i <= x; i++ {
		r := i * i
		if r == x {
			return i
		}
		if r > x {
			return i - 1
		}
	}
	return 0
}
