package main

import (
	"fmt"

	"github.com/timolinn/gorithms/ctci/chapter1"
)

func main() {
	fmt.Println(chapter1.IsUnique("nold"))
	fmt.Println(chapter1.IsUnique("teleri"))
	fmt.Println(chapter1.IsUniqueWithBitVector("noldNOLD"))
	fmt.Println(chapter1.IsUniqueWithBitVector("teleri"))
	fmt.Println(chapter1.CheckPermutation("teleri", "erilte"))
	fmt.Println(chapter1.Compress("aabccccbaa"))
	fmt.Println(chapter1.Compress("abbccdd"))
}
