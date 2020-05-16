package algos

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		want  bool
		word1 string
		word2 string
	}{
		{true, "listen", "silent"},
		{true, "pewes", "sweep"},
		{false, "one", "two"},
		{true, "join", "njio"},
		{false, "ofmekf", "okaopkd"},
		{false, "ofren", "often"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v is anagram of %v", tc.word1, tc.word2), func(t *testing.T) {
			got := Anagram(tc.word1, tc.word2)
			if got != tc.want {
				t.Fatalf("Anagram(%v, %v), want=%v got=%v", tc.word1, tc.word2, tc.want, got)
			}
		})
	}
}

func TestAnagram2(t *testing.T) {
	tests := []struct {
		want  bool
		word1 string
		word2 string
	}{
		{true, "listen", "silent"},
		{true, "pewes", "sweep"},
		{false, "one", "two"},
		{true, "join", "njio"},
		{false, "ofmekf", "okaopkd"},
		{false, "ofren", "often"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v is anagram of %v", tc.word1, tc.word2), func(t *testing.T) {
			got := Anagram2(tc.word1, tc.word2)
			if got != tc.want {
				t.Fatalf("Anagram2(%v, %v), want=%v got=%v", tc.word1, tc.word2, tc.want, got)
			}
		})
	}
}
