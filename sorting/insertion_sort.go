package sorting

func InsertionSort(arr []int) (sorted []int) {
	sorted = append(sorted, arr...)
	length := len(arr)
	i := 0
	aux := 0

	for i < length-1 {
		j := i + 1
		for j < length {
			if sorted[i] > sorted[j] {
				aux = sorted[i]
				sorted[i] = sorted[j]
				sorted[j] = aux
			}
			j++
		}
		i++
	}

	return sorted
}
