package _go

import "sort"

func topKFrequents(nums []int, k int) []int {
	// initialize the array to count the frequency of elements
	numsCounted := make(map[int]int)
	for _, num := range nums {
		numsCounted[num]++
	}
	// create an array only with the keys
	keys := make([]int, 0, len(numsCounted))
	for key := range numsCounted {
		keys = append(keys, key)
	}
	// sort the array
	sort.Slice(keys, func(i, j int) bool {
		return numsCounted[keys[i]] > numsCounted[keys[j]]
	})
	//slice the array from start to end
	return keys[:k]
}
