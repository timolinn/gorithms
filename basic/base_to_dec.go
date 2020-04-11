package gorithms

import (
	"math"
)

// BaseToDec converts value to base 10
func BaseToDec(value string, base int) int {
	var result int
	for i := 0; i < len(value); i++ {
		intVal := getIndexValue(string(value[i]))
		result += intVal * int(math.Pow(float64(base), float64(len(value)-(i+1))))
	}
	return result
}

func getIndexValue(val string) int {
	const charset = "0123456789ABCDEF"
	for i, char := range charset {
		if string(char) == val {
			return i
		}
	}
	return 0
}
