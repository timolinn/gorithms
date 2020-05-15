package datastructures

import (
	"fmt"
)

// Node in a LinkedList data structure
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList data structure
type LinkedList struct {
	Head *Node
	Tail *Node
}

// NewLinkedList creates a new linkedlist
func NewLinkedList() *LinkedList {
	emptyNode := &Node{
		Value: nil,
		Next:  nil,
	}

	return &LinkedList{
		Head: emptyNode,
		Tail: emptyNode,
	}
}

// Add adds a node to the tail
func (l *LinkedList) Add(n interface{}) {
	node := &Node{n, nil}
	if l.Head.Value == nil {
		l.Head = node
		return
	}

	if l.Head.Next == nil {
		l.Head.Next = node
	}
	l.Tail.Next = node
	l.Tail = l.Tail.Next
}

// Search searches the linkedlist
func (l *LinkedList) Search(value interface{}) *Node {
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

// Print outputs all the values in the nodes
func (l *LinkedList) Print() {
	n := l.Head
	for n != nil {
		fmt.Println(n.Value)
		n = n.Next
	}
}

func (l *LinkedList) RemoveDuplicates() {
	cur := l.Head
	prev := l.Head
	values := make(map[int]bool)
	for cur != nil {
		if _, ok := values[cur.Value.(int)]; !ok {
			values[cur.Value.(int)] = true
		} else {
			// duplicate found
			prev.Next = cur.Next
			cur = prev.Next
			continue
		}
		prev = cur
		cur = cur.Next
	}
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
