package main

import (
	"github.com/timolinn/gorithms/bitmanipulation"
)

func main() {
	println(bitmanipulation.GetBit(16, 3))
	println(bitmanipulation.SetBit(16, 3))
	println(bitmanipulation.ClearBit(24, 3))
	println(bitmanipulation.UpdateBit(24, 3, bitmanipulation.Falsy))
	println(bitmanipulation.UpdateBit(16, 3, bitmanipulation.Truthy))
	println(bitmanipulation.HasBit(24, 4))
	println(bitmanipulation.HasBit2(24, 4))
}
