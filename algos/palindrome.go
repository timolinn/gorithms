package algos

// Palidrome returns true if a word is a palindrome
//
// eg. radar <reverse> radar means radar is a palindrome
func Palidrome(word string) bool {
	reverse := Reverse(word)
	return reverse == word
}
