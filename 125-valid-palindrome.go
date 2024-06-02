package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	formattedStr := strings.ToLower(s)
	formattedStr = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(formattedStr, "")
	sir := []rune(formattedStr)

	for i, j := 0, len(sir)-1; i < j; i, j = i+1, j-1 {
		sir[i], sir[j] = sir[j], sir[i]
	}

	return formattedStr == string(sir)
}
