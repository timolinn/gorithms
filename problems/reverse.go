package problems

func ReverseInts(s []int) []int {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		last--
		first++
	}
	return s
}

func ReverseStrings(s []string) []string {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}
