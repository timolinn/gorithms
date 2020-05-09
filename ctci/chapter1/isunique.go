package chapter1

// IsUnique check whether a string has
// all unique characters
func IsUnique(s string) bool {
	cm := make(map[string]int)
	for _, c := range s {
		if _, ok := cm[string(c)]; !ok {
			cm[string(c)] = 1
		} else {
			return false
		}
	}
	return true
}

// IsUniqueWithBitVector check whether a string has
// all unique characters
// This assumes s is all lower case
func IsUniqueWithBitVector(s string) bool {
	bt := int64(0)
	for _, charCode := range s {
		// get te position of the ASCII char between 1 - 26
		val := charCode - 'A'
		if (bt & (1 << val)) > 0 {
			return false
		}
		bt |= (1 << val)
	}
	return true
}
