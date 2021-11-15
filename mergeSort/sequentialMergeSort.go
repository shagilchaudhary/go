package mergeSort

func mergeSequential(arr []int, l int, m int, r int) {

	// Find sizes of two subarrays to be merged
	n1 := m - l + 1
	n2 := r - m

	/* Create temp arrays */
	L := make([]int, n1)
	R := make([]int, n2)

	/*Copy data to temp arrays*/
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	/* Merge the temp arrays */

	// Initial indexes of first and second subarrays
	i, j := 0, 0

	// Initial index of merged subarray array
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	/* Copy remaining elements of L[] if any */
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	/* Copy remaining elements of R[] if any */
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func SortSequential(arr []int, l int, r int) {
	if l < r {
		// Find the middle point
		m := l + (r-l)/2

		// Sort first and second halves
		SortSequential(arr, l, m)
		SortSequential(arr, m+1, r)

		// Merge the sorted halves
		mergeSequential(arr, l, m, r)
	}
}
