package heapSort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// SortSequential is a Sequential Heap Sort Algorithm implementation in Golang
func SortSequential(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		temp := arr[0]
		arr[0] = arr[i]
		arr[i] = temp

		// Heapify root element
		heapify(arr, i, 0)
	}
}

// heapify
func heapify(arr []int, n int, i int) {
	// Find largest among root, left child and right child

	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// Swap and continue heapifying if root is not largest
	if largest != i {
		swap := arr[i]
		arr[i] = arr[largest]
		arr[largest] = swap
		heapify(arr, n, largest)
	}
}
