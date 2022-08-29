package sorting

func BubbleSort(arr []int) []int {
	var sorted []int
	sorted = append(sorted, arr...)
	length := len(arr)
	i := 0
	aux := 0

	for i < length-1 {
		j := 0
		for j < length-1 {

			if sorted[j] > sorted[j+1] {
				aux = sorted[j]
				sorted[j] = sorted[j+1]
				sorted[j+1] = aux
			}

			j++
		}
		i++
	}

	return sorted
}
