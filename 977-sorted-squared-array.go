package main

import "math"

func sortedSquares(nums []int) []int {
	squaredArr := make([]int, len(nums))
	left := 0
	right := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		leftVal := nums[left]
		rightVal := nums[right]

		if math.Abs(float64(leftVal)) > math.Abs(float64(rightVal)) {
			squaredArr[i] = leftVal * leftVal
			left++
		} else {
			squaredArr[i] = rightVal * rightVal
			right--
		}
	}

	return squaredArr
}
