package main

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if _, ok := hashmap[target-nums[i]]; ok {
			return []int{hashmap[target-nums[i]], i}
		}

		hashmap[nums[i]] = i
	}

	return []int{}
}
