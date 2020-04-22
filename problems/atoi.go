package problems

import (
	"fmt"
)

// Atoi is the implementation of C's atoi function
func Atoi(num string) int {
	res := 0
	sign := 1

	if num[0] == '-' {
		sign = -1
		num = num[1:]
	}
	for i := 0; i < len(num); i++ {
		fmt.Println('0', num[i])
		res = res*10 + int(num[i]-'0')
	}
	return sign * res
}
