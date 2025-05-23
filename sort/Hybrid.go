package quicksort

import (
)

const insertionSortThreshold = 20


func InsertionSort(arr []int, low, high int, metrics *Metricas) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		for {
			metrics.Comparisons++ 
			if j < low || arr[j] <= key {
				break
			}
			arr[j+1] = arr[j] 
			j = j - 1
		}
		arr[j+1] = key 
	}
}


func HybridIterativeQuickSort(arr []int, metrics *Metricas) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	stack := make([]Range, 0)
	stack = append(stack, Range{low: 0, high: n - 1})
	maxStackSize := 0

	for len(stack) > 0 {
		
		if len(stack) > maxStackSize {
			maxStackSize = len(stack)
		}

		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low, high := r.low, r.high

		if low >= high {
			continue
		}

		if high-low+1 < insertionSortThreshold {
			InsertionSort(arr, low, high, metrics)
			continue
		}

		medianIdx := findMedianOfThreeIndex(arr, low, high, metrics)
		metrics.Swaps++
		arr[medianIdx], arr[high] = arr[high], arr[medianIdx]

		pi := partition(arr, low, high, metrics)

		stack = append(stack, Range{low: pi + 1, high: high})
		stack = append(stack, Range{low: low, high: pi - 1})
	}
	return maxStackSize
}
