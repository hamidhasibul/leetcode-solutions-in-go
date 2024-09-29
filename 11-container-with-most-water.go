package main

import "math"

func maxArea(height []int) int {

	maxCap := 0

	left := 0
	right := len(height) - 1

	for left < right {
		area := int(math.Abs(float64(left)-float64(right))) * int(math.Min(float64(height[left]), float64(height[right])))

		maxCap = int(math.Max(float64(maxCap), float64(area)))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxCap
}
