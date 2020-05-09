package chapter1

import (
	"fmt"
)

// Compress compresses s
func Compress(s string) string {
	// aabccccbaa => a2b1c4b1a2
	// abbccdd => ab2c2d2 return abbccdd because output is not less than input
	count := 0
	res := fmt.Sprintf("%s%d", string(s[0]), count)
	for i := 0; i < len(s); i++ {
		lastIndex := len(res) - 2
		lastC := res[lastIndex]
		if lastC == s[i] {
			count++
			res = res[:lastIndex]
		} else {
			count = 1
		}
		c := fmt.Sprintf("%s%d", string(s[i]), count)
		res += c
	}

	if len(res) > len(s) {
		return s
	}
	return res
}
