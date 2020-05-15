package main

import (
	"fmt"

	"github.com/timolinn/gorithms/datastructures/trees"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	graph := trees.NewGraph(trees.Directed)
	v0 := graph.AddNode(1)
	v1 := graph.AddNode(2)
	v2 := graph.AddNode(3)
	v3 := graph.AddNode(4)

	graph.AddEdge(v0.ID, v1.ID)
	graph.AddEdge(v0.ID, v3.ID)
	graph.AddEdge(v1.ID, v0.ID)
	graph.AddEdge(v1.ID, v2.ID)
	graph.AddEdge(v2.ID, v1.ID)
	graph.AddEdge(v1.ID, v3.ID)
	graph.AddEdge(v3.ID, v0.ID)
	graph.AddEdge(v3.ID, v1.ID)

	// fmt.Println(graph.HasPathDFS(v3.ID, v0.ID))
	// fmt.Println(graph.HasPathDFS(v2.ID, v1.ID))
	// fmt.Println(graph.HasPathDFS(v3.ID, v1.ID))
	fmt.Println(graph.HasPathBFS(v3.ID, v2.ID))
	fmt.Println("===========================================")
	fmt.Println(graph.HasPathDFS(v1.ID, v0.ID))
	fmt.Println(graph.HasPathDFS(v0.ID, v1.ID))
	fmt.Println(graph.HasPathDFS(v0.ID, v3.ID))
	fmt.Println(graph.HasPathDFS(v1.ID, v3.ID))
	fmt.Println(graph.HasPathDFS(v1.ID, v2.ID))
	fmt.Println(graph.HasPathDFS(v1.ID, v0.ID))
	fmt.Println(graph.HasPathDFS(v3.ID, v0.ID))
	fmt.Println(graph.HasPathDFS(v3.ID, v1.ID))
	// fmt.Println(graph.HasPathBFS(v4.ID, v1.ID))
	fmt.Println("===========================================")
	fmt.Println(graph.Contains(6))
}

/*
func main() {
	bst := trees.BinaryTree{}
	bst.Insert(trees.NewBSTreeNode(4))
	bst.Insert(trees.NewBSTreeNode(3))
	bst.Insert(trees.NewBSTreeNode(7))
	bst.Insert(trees.NewBSTreeNode(1))
	bst.Insert(trees.NewBSTreeNode(0))
	bst.Insert(trees.NewBSTreeNode(9))
	fmt.Println(trees.BSTBreadthFirstSearch(&bst))
}
*/
/*
func main() {
	queue := datastructures.Queue{}
	queue.Enqueue(7)
	queue.Enqueue(70)
	queue.Enqueue(71)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Peek())
}
*/
/*
func main() {
	stack := datastructures.Stack{}
	stack.Push(4)
	stack.Push(41)
	stack.Push(42)
	stack.Push(43)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	fmt.Println(stack)
}*/

/*
func main() {
	l := datastructures.NewLinkedList()
	l.Add(3)
	l.Add(5)
	l.Add(0)
	l.Add(7)
	l.Add(8)
	l.Add(5)
	l.Add(9)
	l.Print()
	fmt.Println("===============================")
	l.RemoveDuplicates()
	l.Print()
}
*/

/*
package main

import (
	"fmt"

	"github.com/timolinn/gorithms/datastructures"

	"github.com/timolinn/gorithms/datastructures/trees"
)

var left *trees.BSTreeNode
var right *trees.BSTreeNode

func main() {
	lnkdlst := datastructures.NewLinkedList()
	lnkdlst.Add(3)
	lnkdlst.Add(5)
	lnkdlst.Add(7)
	lnkdlst.Add(8)
	fmt.Println(lnkdlst.Head.Next.Next)
	fmt.Println(lnkdlst.Tail)

	bst := trees.BinaryTree{}
	bst.Insert(&trees.BSTreeNode{5, nil, nil})
	bst.Insert(&trees.BSTreeNode{3, nil, nil})
	bst.Insert(&trees.BSTreeNode{7, nil, nil})
	bst.Insert(&trees.BSTreeNode{2, nil, nil})
	bst.Insert(&trees.BSTreeNode{1, nil, nil})
	bst.Insert(&trees.BSTreeNode{10, nil, nil})
	bst.Insert(&trees.BSTreeNode{17, nil, nil})
	bst.Insert(&trees.BSTreeNode{3, nil, nil})
	bst.Insert(&trees.BSTreeNode{7, nil, nil})
	fmt.Println(bst.InOrderTraversal())
	fmt.Println(bst.Root.Left)
	fmt.Println(bst.Root.Right)
	res := make(chan int)
	var items []int
	go bst.InOrderTraversalC(res)
	for i := range res {
		items = append(items, i)
	}
	fmt.Println(bst.InOrderTraversal())
	fmt.Println(bst.PreOrderTraversal())
	fmt.Println(bst.PostOrderTraversal())
	fmt.Println(trees.Compare(&bst, &bst))
	fmt.Println(bst.Compare(&bst))
	bst.Remove(&trees.BSTreeNode{3, nil, nil})
	fmt.Println(bst.InOrderTraversal())
	bst.Remove(&trees.BSTreeNode{7, nil, nil})
	fmt.Println(bst.InOrderTraversal())
	fmt.Println(bst.Search(3))
	ht := datastructures.NewHashTable()
	ht = ht.Put("1", 1)
	ht = ht.Put("two", 2)
	ht = ht.Put("1", 5)
	fmt.Println(ht.Get("1"))
	fmt.Println(ht.Get("two"))

	var sb strings.Builder
	for _, c := range []string{"hello", "world", "out", "there"} {
		sb.WriteString(" ")
		sb.WriteString(c)
	}
	var sb datastructures.StringBuilder
	for _, c := range []string{"hello", "world", "out", "there"} {
		sb.WriteString(" ")
		sb.WriteString(c)
	}
	fmt.Println(sb.String())
}*/
