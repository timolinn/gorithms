package main

import (
	"fmt"

	"github.com/timolinn/gorithms/sorting"
)

func main() {
	list := []int{2, 1, 12, 5, 7, 3, 14, 15, 18}
	fmt.Println(sorting.MergeSort(list))
	fmt.Println(sorting.QuickSort(list))
}
