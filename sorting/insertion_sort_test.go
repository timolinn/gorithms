package sorting

import (
	"testing"

	"github.com/timolinn/gorithms/sorting/sorttest"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, InsertionSortInt)
}

func BenchmarkInsertionSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, InsertionSortInt)
}

// func TestInsertionSortString(t *testing.T) {
// 	sorttest.TestString(t, InsertionSortString)
// }

// func TestInsertionSortInterface(t *testing.T) {
// 	sorttest.TestInterface(t, InsertionSort)
// }
