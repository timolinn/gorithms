package sorting

// BubbleSortInt sorts list in ascending order
func BubbleSortInt(list []int) {
	for j := 0; j < len(list); j++ {
		swapped := false
		for i := 0; i < len(list)-1-j; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
