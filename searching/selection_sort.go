package searching

// SelectionSort toma un arreglo de enteros y regresa una copia ordenada
// siguiende el algoritmo selection sort
func SelectionSort(arr []int) []int {
	arrLength := len(arr)
	sorted := make([]int, arrLength)
	current := 0

	i := 0

	for i < arrLength {
		j := i + 1
		minPos := i
		for j < arrLength {

			if arr[minPos] > arr[j] {
				minPos = j
			}

			j++
		}

		sorted[current] = arr[minPos]
		current++
		i++
	}

	return sorted
}
