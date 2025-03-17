package main

import (
	"math"
	"testing"
)

func BenchmarkCalculateDigitsLog10(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		_ = int(math.Floor(math.Log10(float64(i))))
	}
}
func BenchmarkCalculateDigitsDivision10(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		_ = calDigits(i)
	}
}

func calDigits(num int) int {
	digits := 0
	for num != 0 {
		num /= 10
		digits++
	}
	return digits
}
