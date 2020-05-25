package concurrency

import (
	"fmt"
	"sync"
)

type WebSiteChecker func(string) bool

var wg sync.WaitGroup

// Checker does nothing serious
func Checker() int {
	waitCount := 5
	wg.Add(waitCount)
	for i := 0; i < waitCount; i++ {
		go func(index int) {
			fmt.Println(index)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return 0
}
