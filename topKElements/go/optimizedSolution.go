package _go

import (
	"container/heap"
)

// Pair is a simple key-value pair representing the number and its frequency.
type Pair struct {
	number    int
	frequency int
}

// A MinHeap implements heap.Interface and holds Pairs.
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num]++
	}

	// Initialize a min heap.
	h := &MinHeap{}
	heap.Init(h)

	// Keep only the top k elements in the heap.
	for num, freq := range numCount {
		heap.Push(h, Pair{num, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Extract the numbers from the heap.
	topK := make([]int, k)
	for i := range topK {
		topK[i] = heap.Pop(h).(Pair).number
	}
	return topK
}
