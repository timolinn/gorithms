package main

import (
	"fmt"

	"github.com/timolinn/gorithms/datastructures"

	"github.com/timolinn/gorithms/datastructures/trees"
)

var left *trees.TreeNode
var right *trees.TreeNode

func main() {
	// lnkdlst := datastructures.NewLinkedList()
	// lnkdlst.Add(3)
	// lnkdlst.Add(5)
	// lnkdlst.Add(7)
	// lnkdlst.Add(8)
	// fmt.Println(lnkdlst.Head.Next.Next)
	// fmt.Println(lnkdlst.Tail)

	// bst := trees.BinaryTree{}
	// bst.Insert(&trees.TreeNode{5, nil, nil})
	// bst.Insert(&trees.TreeNode{3, nil, nil})
	// bst.Insert(&trees.TreeNode{7, nil, nil})
	// bst.Insert(&trees.TreeNode{2, nil, nil})
	// bst.Insert(&trees.TreeNode{1, nil, nil})
	// bst.Insert(&trees.TreeNode{10, nil, nil})
	// bst.Insert(&trees.TreeNode{17, nil, nil})
	// bst.Insert(&trees.TreeNode{3, nil, nil})
	// bst.Insert(&trees.TreeNode{7, nil, nil})
	// fmt.Println(bst.InOrderTraversal())
	// fmt.Println(bst.Root.Left)
	// fmt.Println(bst.Root.Right)
	// res := make(chan int)
	// var items []int
	// go bst.InOrderTraversalC(res)
	// for i := range res {
	// 	items = append(items, i)
	// }
	// fmt.Println(bst.InOrderTraversal())
	// fmt.Println(bst.PreOrderTraversal())
	// fmt.Println(bst.PostOrderTraversal())
	// fmt.Println(trees.Compare(&bst, &bst))
	// fmt.Println(bst.Compare(&bst))
	// bst.Remove(&trees.TreeNode{3, nil, nil})
	// fmt.Println(bst.InOrderTraversal())
	// bst.Remove(&trees.TreeNode{7, nil, nil})
	// fmt.Println(bst.InOrderTraversal())
	// fmt.Println(bst.Search(3))
	// ht := datastructures.NewHashTable()
	// ht = ht.Put("1", 1)
	// ht = ht.Put("two", 2)
	// ht = ht.Put("1", 5)
	// fmt.Println(ht.Get("1"))
	// fmt.Println(ht.Get("two"))

	// var sb strings.Builder
	// for _, c := range []string{"hello", "world", "out", "there"} {
	// 	sb.WriteString(" ")
	// 	sb.WriteString(c)
	// }
	var sb datastructures.StringBuilder
	for _, c := range []string{"hello", "world", "out", "there"} {
		sb.WriteString(" ")
		sb.WriteString(c)
	}
	fmt.Println(sb.String())
}
