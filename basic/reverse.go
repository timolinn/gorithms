package gorithms

// Reverse reverses a word
func Reverse(word string) string {
	var reversed string
	for _, c := range word {
		reversed = string(c) + reversed
	}
	return reversed
}
