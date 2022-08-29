package sorting

// MergeSort ordena un arreglo de enteros
// empleando el algoritmo de ordenamiento merge sort
func MergeSort(arr []int) (sorted []int) {
	sorted = mergeSortAux(arr, 0, len(arr)-1)
	return sorted
}

func mergeSortAux(arr []int, left, right int) []int {
	if left == right {
		return []int{arr[left]}
	}

	middle := (left + right) / 2

	leftHalf := mergeSortAux(arr, left, middle)
	rightHalf := mergeSortAux(arr, middle+1, right)

	sorted := mixHalves(leftHalf, rightHalf)
	return sorted
}

func mixHalves(leftHalf, rightHalf []int) (sorted []int) {
	left := 0
	right := 0

	for left < len(leftHalf) && right < len(rightHalf) {
		if leftHalf[left] <= rightHalf[right] {
			sorted = append(sorted, leftHalf[left])
			left++
		} else if rightHalf[right] <= leftHalf[left] {
			sorted = append(sorted, rightHalf[right])
			right++
		}
	}

	for left < len(leftHalf) {
		sorted = append(sorted, leftHalf[left])
		left++
	}

	for right < len(rightHalf) {
		sorted = append(sorted, rightHalf[right])
		right++
	}

	return sorted
}
