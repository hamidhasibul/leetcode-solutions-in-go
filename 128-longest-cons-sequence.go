package main

import (
	"math"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxCount, count := 0, 1

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i]-nums[i-1] == 1 {
				count++
			} else {
				maxCount = int(math.Max(float64(maxCount), float64(count)))
				count = 1
			}
		}
	}

	return int(math.Max(float64(maxCount), float64(count)))
}
