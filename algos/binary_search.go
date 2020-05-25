package algos

// BinarySearch searches n in list using
// binary search algorithm
// returns the index of the element if found
// returns -1 otherwise
func BinarySearch(list []int, n int) int {
	// [1, 2, 3, 4, 5, 6, 7] 1 => 3
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == n {
			return mid
		}

		if n > list[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BinarySearchRecursive(list []int, n, low, high int) int {
	mid := (low + high) / 2
	midValue := list[mid]
	if midValue == n {
		return mid
	} else if midValue < n {
		return BinarySearchRecursive(list, n, mid+1, high)
	} else {
		return BinarySearchRecursive(list, n, low, mid-1)
	}
}
