package sorting

func QuickSort(arr []int) (sorted []int) {
	sorted = append(sorted, arr...)
	quickSortAux(sorted, 0, len(arr)-1)
	return sorted
}

func quickSortAux(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortAux(arr, low, pi-1)
		quickSortAux(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) (idx int) {
	pivot := arr[high]
	idx = (low - 1)

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			idx++
			arr[idx], arr[j] = arr[j], arr[idx]
		}
	}
	arr[idx+1], arr[high] = arr[high], arr[idx+1]
	return idx + 1
}
