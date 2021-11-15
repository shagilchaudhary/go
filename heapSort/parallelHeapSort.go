package heapSort

// Reference
// Mwaffaq A. Abu Al Hija , Arwa Zabian , Sami Qawasmeh , Omer H. Abu Al Haija A Heapify Based Parallel Sorting Algorithm (2008)
// Heap Sort in Golang

import (
	"math"
	"runtime"
	"sync"
	// "sort"
)

// min finds the minimum value in array
func min(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func partition(arr []int, magicNumber int, l int, r int, ch chan []int) {

	if len(arr) <= magicNumber {
		ch <- arr
		return
	}

	if l < r {
		// Find the middle point
		m := l + (r-l)/2

		arrLeft := arr[0 : m+1]
		arrRight := arr[m+1:]

		arrLeftMin := min(arrLeft)
		arrLeftMax := max(arrLeft)

		arrRightMin := min(arrRight)
		arrRightMax := max(arrRight)

		arrMidVal := int(math.Ceil(float64(arrLeftMin+arrLeftMax+arrRightMin+arrRightMax) / float64(4)))

		x := 0
		for {
			if arrLeft[x] >= arrMidVal {
				y := (len(arrRight) - 1)
				for y >= 0 {
					if arrRight[y] <= arrMidVal || arrRight[y] < arrLeft[x] {
						swap := arrRight[y]
						arrRight[y] = arrLeft[x]
						arrLeft[x] = swap
					}
					y = y - 1
				}
			}
			x = x + 1
			if x >= len(arrLeft) {
				break
			}
		}

		a := 0
		for {
			if arrRight[a] <= arrMidVal {
				b := (len(arrLeft) - 1)
				for b >= 0 {
					if arrLeft[b] >= arrMidVal || arrLeft[b] > arrRight[a] {
						swap := arrLeft[b]
						arrLeft[b] = arrRight[a]
						arrRight[a] = swap
					}
					b = b - 1
				}
			}
			a = a + 1
			if a >= len(arrRight) {
				break
			}
		}

		partition(arrLeft, magicNumber, l, m, ch)
		partition(arrRight, magicNumber, l, m, ch)
	}
}

func goPartition(arr []int, magicNumber int, l int, r int, ch chan []int) {
	partition(arr, magicNumber, l, r, ch)
	close(ch)
}

func SortParallel(arr []int, l int, r int) {

	// Magic Number 1000 :(
	// Need to develop formula for no of partitioned arrays based on input size
	// Close buffered channel as the buffer is huge and receiver will be communicated that all values have been received
	brokenArraysLen := 1000
	partitionedArraysChannel := make(chan []int, brokenArraysLen)

	// Magic Number 3 is for keeping minimum no of elements
	// Bad Coding Should be based on some criteria on input
	magicNumber := 3

	maxPartitionedArraySize := int(math.Ceil(float64(len(arr)) / float64(runtime.GOMAXPROCS(0))))

	if maxPartitionedArraySize > 3 {
		magicNumber = maxPartitionedArraySize
	}

	go goPartition(arr, magicNumber, l, r, partitionedArraysChannel)

	var wg sync.WaitGroup
	var sortedPartitionedArrays [][]int
	for i := 0; i < brokenArraysLen; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			partitionedArray := <-partitionedArraysChannel
			SortSequential(partitionedArray)
			sortedPartitionedArrays = append(sortedPartitionedArrays, partitionedArray)
		}(i)
	}

	wg.Wait()

	// Below code is required when the goPartition recursive would be
	// Even it requires tweaks when first element would be same

	// sort.Slice(sortedPartitionedArrays, func(i, j int) bool {
	//   if len(sortedPartitionedArrays[i]) == 0 && len(sortedPartitionedArrays[j]) == 0 {
	//     return false
	//   }
	//   if len(sortedPartitionedArrays[i]) == 0 || len(sortedPartitionedArrays[j]) == 0 {
	//     return len(sortedPartitionedArrays[i]) == 0
	//   }
	//   return sortedPartitionedArrays[i][0] < sortedPartitionedArrays[j][0]
	// })
}
