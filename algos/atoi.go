package algos

// Atoi converts string num to interger value
func Atoi(strnum string) int {
	// "123" => 123
	// "-123" => -123
	sign := 1
	if strnum[0] == '-' {
		sign = -1
		strnum = strnum[1:]
	}

	result := 0
	for _, num := range strnum {
		real := num - '0'
		result = result*10 + int(real)
	}
	return result * sign
}
