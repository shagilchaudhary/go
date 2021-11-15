package mergeSort

func mergeParallel(leftArr []int, rightArr []int) (result []int) {

	n1 := len(leftArr)
	n2 := len(rightArr)

	// Creating Temp Array
	result = make([]int, (n1 + n2))

	// Initial indexes of first and second subarrays
	i, j := 0, 0

	for k := 0; k < cap(result); k++ {
		if i >= len(leftArr) {
			result[k] = rightArr[j]
			j++
		} else if j >= len(rightArr) {
			result[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			result[k] = leftArr[i]
			i++
		} else {
			result[k] = rightArr[j]
			j++
		}
	}

	return
}

func sortParallel(arr []int, result chan []int) {

	if len(arr) == 1 {
		result <- arr
		return
	}

	// var wg sync.WaitGroup
	leftChan := make(chan []int)
	rightChan := make(chan []int)

	m := len(arr) / 2

	// wg.Add(2)
	go sortParallel(arr[:m], leftChan)
	go sortParallel(arr[m:], rightChan)

	// wg.Wait()
	leftArrSorted := <-leftChan
	rightArrSorted := <-rightChan

	close(leftChan)
	close(rightChan)

	// wg.Add(1)
	finalChan := make(chan []int)
	result <- mergeParallel(leftArrSorted, rightArrSorted)
	close(finalChan)

	// wg.Wait()
	return
}

func GoSortParallel(arr []int) []int {
	result := make(chan []int)
	defer close(result)
	go sortParallel(arr, result)
	return <-result
}
