package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	chMap := map[byte]int{}
	l, r := 0, 0
	length := 0

	for r < len(s) {
		if _, ok := chMap[s[r]]; ok && chMap[s[r]] >= l {
			l = chMap[s[r]] + 1
		}
		chMap[s[r]] = r

		length = int(math.Max(float64(length), float64(r-l+1)))

		r++
	}

	return length
}
