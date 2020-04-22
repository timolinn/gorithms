package main

import (
	"fmt"

	"github.com/timolinn/gorithms/sorting"
)

func main() {
	list := []int{2, 1, 3, 5, 4, 7, 6, 30, 12, 20, 21, 11, 29, 30, 50, 44, 43, 42, 98}
	sorting.BubbleSortInt(list)
	fmt.Println(list)
}
