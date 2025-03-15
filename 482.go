package main

import "strings"

func licenseKeyFormatting(s string, k int) string {
	var revBuilder strings.Builder
	partCount := 0
	for i := len(s) - 1; i > -1; i-- {
		ch := s[i]
		if ch == '-' {
			continue
		}
		if partCount == k {
			revBuilder.WriteRune('-')
			partCount = 0
		}
		if ch > 96 {
			ch -= 32
		}
		revBuilder.WriteByte(ch)
		partCount++
	}
	reversed := revBuilder.String()
	var resBuilder strings.Builder
	for i := len(reversed) - 1; i > -1; i-- {
		resBuilder.WriteByte(reversed[i])
	}

	return resBuilder.String()
}
