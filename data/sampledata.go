package data

import (
	"math/rand"
	"runtime"
	"slices"
	"sync"
	"time"
)

func SliceOrdenado(nvalor int) []int {

	size := nvalor + 1
	data := make([]int, size)
	
	numGoroutines := runtime.NumCPU()

	if numGoroutines == 0 {
		numGoroutines = 1
	}

	chunkSize := (size + numGoroutines - 1) / numGoroutines

	var wg sync.WaitGroup

	for i := range numGoroutines {
		wg.Add(1)

		startIdx := i * chunkSize
		endIdx := startIdx + chunkSize

		if endIdx > size {
			endIdx = size
		}

		if startIdx >= endIdx {
			wg.Done()
			continue
		}

		go func(e, s int) {
			defer wg.Done()

			for i := e - 1; i >= s; i-- {
				data[i] = nvalor - i
			}


		}(endIdx, startIdx)
		
	}
	wg.Wait()
	return data
}

func randomSlice(size, maxValueExclusive int) []int {
	data := make([]int, 100000)

	numGoroutines := runtime.NumCPU()
	
	if numGoroutines == 0 {
		numGoroutines = 1
	}

	chunkSize := (size + numGoroutines - 1) / numGoroutines

	var wg sync.WaitGroup
	baseSeed := time.Now().UnixNano()

		for i := range numGoroutines {
		wg.Add(1)

		startIdx := i * chunkSize
		endIdx := startIdx + chunkSize

		if endIdx > size {
			endIdx = size
		}

		if startIdx >= endIdx {
			wg.Done()
			continue
		}

		go func(sIdx, eIdx, id int) {
			defer wg.Done()

			source := rand.NewSource(baseSeed + int64(id))
			randomGenerator := rand.New(source)
			for i := sIdx; i < eIdx; i++ {
				data[i] = randomGenerator.Intn(maxValueExclusive)
			}
		}(startIdx, endIdx, i)		
	}
	wg.Wait()

	slices.Sort(data)
	slices.Reverse(data)
	
	return data
}
func GenerateData() []int {

	const maxValue = 100000000
	const size = 100000

	data := randomSlice(size, maxValue)
	return data
}

func DescendingOrderedData() []int {
	const size = 100000

	return SliceOrdenado(size)


}
