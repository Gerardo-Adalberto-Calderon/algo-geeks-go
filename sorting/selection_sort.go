package sorting

// SelectionSort toma un arreglo de enteros y regresa una copia ordenada
// siguiende el algoritmo selection sort
func SelectionSort(arr []int) []int {
	var sorted []int
	arrLength := len(arr)
	i := 0
	sorted = append(sorted, arr...)
	aux := 0

	for i < arrLength {
		j := i + 1
		minPos := i
		for j < arrLength {

			if sorted[minPos] > sorted[j] {
				minPos = j
			}

			j++
		}

		aux = sorted[i]
		sorted[i] = sorted[minPos]
		sorted[minPos] = aux
		i++
	}

	return sorted
}
