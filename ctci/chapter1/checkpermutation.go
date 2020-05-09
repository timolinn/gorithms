package chapter1

// CheckPermutation checks if s2 is a permutation of s1
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Map := make(map[string]int)
	for _, c := range s1 {
		char := string(c)
		if _, ok := s1Map[char]; !ok {
			s1Map[char] = 1
		} else {
			s1Map[char]++
		}
	}

	for _, c := range s2 {
		if val, ok := s1Map[string(c)]; ok {
			if val < 1 {
				return false
			}
			s1Map[string(c)]--
		} else {
			return false
		}
	}
	return true
}
