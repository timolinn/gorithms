package sorting

import (
	"math"
)

// MergeSort sorts a list of ints using mergesort algorithm
func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	mid := math.Floor(float64(len(list) / 2))
	left := list[:int(mid)]
	right := list[int(mid):]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(l1, l2 []int) []int {
	var result []int
	for len(l1) > 0 && len(l2) > 0 {
		if l1[0] < l2[0] {
			result = append(result, l1[0])
			l1 = l1[1:]
		} else {
			result = append(result, l2[0])
			l2 = l2[1:]
		}
	}
	return append(result, append(l1, l2...)...)
}
