package trees

// Node in a LinkedList data structure
type Node struct {
	Value int
	Next  *Node
}

// LinkedList data structure
type LinkedList struct {
	Head *Node
	Tail *Node
}

// Add adds a node to the tail
func (l *LinkedList) Add(n *Node) {
	if l.Head == nil {
		l.Head = n
		return
	}

	for {
		if l.Head.Next == nil {
			l.Head.Next = n
			return
		}
		l.Head.Next = l.Head.Next.Next
	}
}

// Search searches the linkedlist
func (l *LinkedList) Search(value int) *Node {
	if l.Head.Value == value {
		return l.Head
	}

	lv := l.Head
	for {
		if lv.Next.Value == value {
			return lv.Next
		}
		if lv.Next.Next == nil {
			break
		}
		lv.Next = lv.Next.Next
	}
	return nil
}

// import (
// 	"fmt"

// 	"github.com/timolinn/gorithms/trees"
// )

// func main() {
// 	lnkdlst := trees.LinkedList{
// 		Head: &trees.Node{
// 			Value: 3,
// 		},
// 	}
// 	five := trees.Node{
// 		Value: 5,
// 		Next: &trees.Node{
// 			Value: 7,
// 		},
// 	}
// 	lnkdlst.Add(&five)
// 	fmt.Println(lnkdlst.Head.Next.Next)
// 	fmt.Println(lnkdlst.Search(7).Value)
// }
