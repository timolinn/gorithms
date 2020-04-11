package gorithms

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Printf("Fizz Buzz")
		} else if i%5 == 0 {
			fmt.Printf("Buzz")
		} else if i%3 == 0 {
			fmt.Printf("Fizz")
		} else {
			fmt.Printf("%d", i)
		}
		if i != n {
			fmt.Printf(", ")
		}
	}
	fmt.Println()
}
