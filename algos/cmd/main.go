package main

import (
	"fmt"

	"github.com/timolinn/gorithms/algos"
	"github.com/timolinn/gorithms/bitmanipulation"
)

func main() {
	fmt.Println(algos.Anagram("listen", "silent"))
	fmt.Println(algos.Anagram2("pwese", "sweep"))
	fmt.Println(algos.Palidrome("radar"))
	fmt.Println(algos.Palidrome("sweep"))
	fmt.Println(algos.Atoi("123"))
	fmt.Println(algos.Atoi("-123"))
	fmt.Println("============ BIT MANIPULATION =================")
	fmt.Println(bitmanipulation.GetBit(16, 3))
	fmt.Println(bitmanipulation.SetBit(16, 3))
	fmt.Println(bitmanipulation.ClearBit(24, 3))
	fmt.Println(bitmanipulation.UpdateBit(24, 3, bitmanipulation.Falsy))
	fmt.Println(bitmanipulation.UpdateBit(16, 3, bitmanipulation.Truthy))
	fmt.Println("=========== Search =================")
	fmt.Println(algos.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7}, 7))
	fmt.Println(algos.BinarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7}, 7, 0, len([]int{1, 2, 3, 4, 5, 6, 7})))
	fmt.Println("====================================")
	fmt.Println(algos.Toss(algos.NewCoin()))
}
