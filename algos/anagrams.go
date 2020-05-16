package algos

import (
	"sort"
	"strings"
)

func Anagram(word1, word2 string) bool {
	// listen <anagram> silent == true
	// sweep <anagram> pewse
	if len(word1) != len(word2) {
		return false
	}

	word1Slice := strings.Split(word1, "")
	word2Slice := strings.Split(word2, "")
	sort.Strings(word1Slice)
	sort.Strings(word2Slice)
	return strings.Join(word1Slice, "") == strings.Join(word2Slice, "")
}

func Anagram2(w1, w2 string) bool {
	// listen <anagram> silent == true
	// sweep <anagram> pewse
	if len(w1) != len(w2) {
		return false
	}

	w1Map := make(map[string]int, len(w1))
	for _, c := range w1 {
		if _, ok := w1Map[string(c)]; !ok {
			w1Map[string(c)] = 1
			continue
		}
		w1Map[string(c)]++
	}

	for _, c := range w2 {
		if _, ok := w1Map[string(c)]; !ok || w1Map[string(c)] < 1 {
			return false
		}
		w1Map[string(c)]--
	}

	return true
}

// Reverse reverses a string
func Reverse(text string) string {
	textSlice := strings.Split(text, "")
	first := 0
	last := len(text) - 1
	for first < last {
		textSlice[first], textSlice[last] = textSlice[last], textSlice[first]
		last--
		first++
	}
	return strings.Join(textSlice, "")
}
