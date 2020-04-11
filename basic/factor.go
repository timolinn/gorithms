package gorithms

// Factor returns a list of factors of number
func Factor(primes []int, number int) []int {
	var result []int
	for i := 0; i < len(primes); i++ {
		element := primes[i]
		if number%element == 0 {
			result = append(result, element)
			number = number / element

			// repeat with the same prime
			i--
		}

	}
	if number > 1 {
		result = append(result, number)
	}
	return result
}
