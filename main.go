package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"text/tabwriter"
	"time"

	"./heapSort"
	"./mergeSort"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, "Algorithm\tMax CPUs\tInputSize\tSequential Processing Time\tParallel Processing Time\tEquality Check\t")

	inputInt, _ := strconv.Atoi(os.Args[1])

	var unsortedHeapSortArraySequential []int
	var unsortedHeapSortArrayParallel []int

	var unsortedMergeSortArraySequential []int
	var unsortedMergeSortArrayParallel []int

	for i := 0; i < inputInt; i++ {

		randInt := rand.Intn(inputInt)

		unsortedHeapSortArraySequential = append(unsortedHeapSortArraySequential, randInt)
		unsortedHeapSortArrayParallel = append(unsortedHeapSortArrayParallel, randInt)
		unsortedMergeSortArraySequential = append(unsortedMergeSortArraySequential, randInt)
		unsortedMergeSortArrayParallel = append(unsortedMergeSortArrayParallel, randInt)
	}

	startHeapSortSequentialTime := time.Now()
	heapSort.SortSequential(unsortedHeapSortArraySequential)
	totalHeapSortSequentialTime := time.Since(startHeapSortSequentialTime)

	startHeapSortParallelTime := time.Now()
	heapSort.SortParallel(unsortedHeapSortArrayParallel, 0, len(unsortedHeapSortArrayParallel)-1)
	totalHeapSortParallelTime := time.Since(startHeapSortParallelTime)

	// Equality check to confirm sorting order remains same
	equalityCheckHeapSort := true

	for i, e := range unsortedHeapSortArraySequential {
		if e != unsortedHeapSortArrayParallel[i] {
			equalityCheckHeapSort = false
			break
		}
	}

	heapSortRow := fmt.Sprintf("Heap Sort\t%d\t%d\t%v\t%v\t%t\t", runtime.GOMAXPROCS(0), inputInt, totalHeapSortSequentialTime, totalHeapSortParallelTime, equalityCheckHeapSort)
	fmt.Fprintln(w, heapSortRow)

	startMergeSortSequentialTime := time.Now()
	mergeSort.SortSequential(unsortedMergeSortArraySequential, 0, (len(unsortedMergeSortArraySequential) - 1))
	totalMergeSortSequentialTime := time.Since(startMergeSortSequentialTime)

	startMergeSortParallelTime := time.Now()
	unsortedMergeSortArrayParallel = mergeSort.GoSortParallel(unsortedMergeSortArrayParallel)
	totalMergeSortParallelTime := time.Since(startMergeSortParallelTime)

	// Equality check to confirm sorting order remains same
	equalityCheckMergeSort := true

	for i, e := range unsortedMergeSortArraySequential {
		if e != unsortedMergeSortArrayParallel[i] {
			equalityCheckMergeSort = false
			break
		}
	}

	mergeSortRow := fmt.Sprintf("Merge Sort\t%d\t%d\t%v\t%v\t%t\t", runtime.GOMAXPROCS(0), inputInt, totalMergeSortSequentialTime, totalMergeSortParallelTime, equalityCheckMergeSort)
	fmt.Fprintln(w, mergeSortRow)

	w.Flush()
}
