package main

// Linear Time & O(n) Space

func majorityElement(nums []int) int {

	frequency := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		frequency[nums[i]]++
	}

	for k, v := range frequency {
		if len(nums)/2 < v {
			return k
		}
	}

	return -1
}

// Boyer Moore O(n) & O(1)

// func majorityElement(nums []int) int {

// }
