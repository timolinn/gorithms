package main

import (
	"github.com/timolinn/gorithms/algos"
	"github.com/timolinn/gorithms/bitmanipulation"
)

func main() {
	println(algos.Anagram("listen", "silent"))
	println(algos.Anagram2("pwese", "sweep"))
	println(algos.Palidrome("radar"))
	println(algos.Palidrome("sweep"))
	println(algos.Atoi("123"))
	println(algos.Atoi("-123"))
	println("============ BIT MANIPULATION =================")
	println(bitmanipulation.GetBit(16, 3))
	println(bitmanipulation.SetBit(16, 3))
	println(bitmanipulation.ClearBit(24, 3))
	println(bitmanipulation.UpdateBit(24, 3, bitmanipulation.Falsy))
	println(bitmanipulation.UpdateBit(16, 3, bitmanipulation.Truthy))
}
