package main

import (
	"fmt"

	"github.com/Gerardo-Adalberto-Calderon/algo-geeks-go/sorting"
)

func main() {
	arr := []int{8, 13, 1, -2, 0, 5, 9}

	sorted := sorting.SelectionSort(arr)

	fmt.Println(sorted)
}
