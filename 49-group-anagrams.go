package main

import "sort"

func groupAnagrams(strs []string) [][]string {

	countMap := map[string][]string{}
	tempArr := [][]string{}

	for _, val := range strs {
		runes := []rune(val)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedStr := string(runes)

		countMap[sortedStr] = append(countMap[sortedStr], val)
	}

	for _, v := range countMap {
		tempArr = append(tempArr, v)
	}

	return tempArr

}
