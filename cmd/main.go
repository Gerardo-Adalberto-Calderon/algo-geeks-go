package main

import (
	"fmt"

	"github.com/Gerardo-Adalberto-Calderon/algo-geeks-go/sorting"
)

func main() {
	arr := []int{8, 13, 1, -2, 0, 5, 9, 294, 24, 68, 23}

	sorted := sorting.MergeSort(arr)

	fmt.Println(sorted)
}
