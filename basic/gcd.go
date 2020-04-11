package gorithms

// GCD returns the highest number that can divide a and b
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
