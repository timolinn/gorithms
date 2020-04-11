package gorithms

// NumInList will return true if num is in list
func NumInList(list []int, num int) bool {
	if list == nil {
		return false
	}
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}
