package main

import (
	"fmt"

	"github.com/teivah/bitvector"
)

var bv bitvector.Len8

func main() {
	fmt.Println('z' - 'a')
	fmt.Println(0 & (1 << ('z' - 'a')))
	fmt.Println(int([]byte("129")[0]))
}
