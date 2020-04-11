package gorithms

// FindTwoThatSum returns the index of two numbers that
// add up to sum
func FindTwoThatSum(numbers []int, sum int) (int, int) {
	for i, num := range numbers {
		for j := i; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if (num + numbers[j]) == sum {
				return i, j
			}
		}
	}
	return -1, -1
}

// FindTwoThatSum2 returns the index of two numbers that
// add up to sum
func FindTwoThatSum2(numbers []int, sum int) (int, int) {
	if len(numbers) < 1 {
		return -1, -1
	}

	m := make(map[int]int)
	for i, num := range numbers {
		if j, ok := m[num]; ok {
			return j, i
		}
		m[sum-num] = i
	}
	return -1, -1
}
