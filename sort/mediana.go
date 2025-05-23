package quicksort

func Mquicksort(array []int, low, high int, metrics *Metricas) {
	if low < high {
		if high-low+1 >= 3 {
			medianaIdx := findMedianOfThreeIndex(array, low, high, metrics)
			metrics.Swaps++ 
			array[medianaIdx], array[high] = array[high], array[medianaIdx]
		}

		
		pi := partition(array, low, high, metrics)

		
		Mquicksort(array, low, pi-1, metrics)
		Mquicksort(array, pi+1, high, metrics)
	}
}
