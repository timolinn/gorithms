package sorting

// InsertionSortInt sorts the items in a list using
// InsertionSort algorithm
func InsertionSortInt(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insert(sorted, item)
	}
	for i, val := range sorted {
		list[i] = val
	}
}

func insert(sortedList []int, item int) []int {
	for i := 0; i < len(sortedList); i++ {
		if item < sortedList[i] {
			return append(sortedList[:i], append([]int{item}, sortedList[i:]...)...)
		}
	}
	return append(sortedList, item)
}

// func BinarySearch
