package main

// func productExceptSelf(nums []int) []int {

// 	prodArr := make([]int, len(nums))

// 	prod := 1

// 	for _, val := range nums {
// 		prod *= val
// 	}

// 	for i, val := range nums {
// 		prodArr[i] = prod / val
// 	}

// 	return prodArr

// }

func productExceptSelf(nums []int) []int {

	prodArr := make([]int, len(nums))

	l, r := 1, 1

	for i := 0; i < len(nums); i++ {
		prodArr[i] = l
		l *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		prodArr[i] *= r
		r *= nums[i]

	}

	return prodArr
}
