package gorithms

// Sum return the sum of all values in the list
func Sum(list []int) int {
	result := 0
	for _, n := range list {
		result += n
	}
	return result
}

// RecursiveSum recursive solution to Sum
func RecursiveSum(list []int) int {
	if len(list) < 1 {
		return 0
	}
	return list[0] + Sum(list[1:])
}
