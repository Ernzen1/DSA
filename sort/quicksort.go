package quicksort


type Range struct {
	low  int
	high int
}

type Metricas struct {
	Comparisons int64
	Swaps       int64
}

type Elemento struct {
	Valor int
	Indice int
}

func findMedianOfThreeIndex(arr []int, low, high int, metrics *Metricas) int {
	mid := low + (high-low)/2
	l, m, h := arr[low], arr[mid], arr[high]

	metrics.Comparisons++
	if l < m { 
		metrics.Comparisons++
		if m < h { 
			return mid
		}
		metrics.Comparisons++
		if l < h { 
			return high
		}
		return low 
	} else { 
		metrics.Comparisons++
		if l < h {
			return low
		}
		metrics.Comparisons++
		if m < h { 
			return high
		}
		return mid 
	}
}

func partition(arr []int, low, high int, metrics *Metricas) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		metrics.Comparisons++ 
		if arr[j] <= pivot {
			i++
			metrics.Swaps++ 
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	metrics.Swaps++
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}


func Quicksort(array []int, metrics *Metricas) []int {
	if len(array) < 2 {
		return array
	}

	pivo := array[len(array)-1]
	resto := array[:len(array)-1]

	menores := make([]int, 0, len(resto)/2)
	maiores := make([]int, 0, len(resto)/2)

	for _, valor := range resto {
		metrics.Comparisons++ // Conta comparação
		if valor <= pivo {
			menores = append(menores, valor)
		} else {
			maiores = append(maiores, valor)
		}
	}

	menoresOrdenados := Quicksort(menores, metrics)
	maioresOrdenados := Quicksort(maiores, metrics)

	final := make([]int, 0, len(array))
	final = append(final, menoresOrdenados...)
	final = append(final, pivo)
	final = append(final, maioresOrdenados...)

	return final
}
