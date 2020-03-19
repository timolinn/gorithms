package gorithms

import (
	"strconv"
)

// DecToBase converts decimal integer to specified base
func DecToBase(dec, base int) string {
	var res string
	var base16 = map[string]string{
		"10": "A",
		"11": "B",
		"12": "C",
		"13": "D",
		"14": "E",
		"15": "F",
	}

	for dec > 0 {
		rem := strconv.Itoa(dec % base)
		if v, _ := strconv.Atoi(rem); v > 9 {
			rem = base16[rem]
		}
		dec = dec / base
		res = rem + res
	}
	return res
}

// DecToBaseAlternate converts decimal integer to specified base
func DecToBaseAlternate(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		dec = dec / base
		res = string(charset[rem]) + res
	}
	return res
}
