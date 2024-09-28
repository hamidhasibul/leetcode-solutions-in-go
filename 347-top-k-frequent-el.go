package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {

	freqMap := map[int]int{}

	for _, val := range nums {
		freqMap[val]++
	}

	keys := make([]int, 0, len(freqMap))

	for key := range freqMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	keys = keys[:k]

	return keys

}
