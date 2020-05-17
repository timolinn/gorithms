package sorting

func QuickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	var result []int
	pivot := list[len(list)-1]
	list = list[:len(list)-1]
	left, right := make([]int, 0), make([]int, 0)
	for _, v := range list {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	result = append(result, QuickSort(left)...)
	result = append(result, pivot)
	return append(result, QuickSort(right)...)
}
