package main

func containsDuplicate(nums []int) bool {

	// O(n^2)

	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] == nums[j] {
	// 			return true
	// 		}
	// 	}
	// }

	// O(n)

	existentialMap := map[int]bool{}

	for _, val := range nums {
		if existentialMap[val] {
			return true
		} else {
			existentialMap[val] = true
		}

	}

	return false
}
