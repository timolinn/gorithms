package main

import (
	"fmt"

	"github.com/timolinn/gorithms/trees"
)

var left *trees.TreeNode
var right *trees.TreeNode

func main() {
	lnkdlst := trees.LinkedList{
		Head: &trees.Node{
			Value: 3,
		},
	}
	five := trees.Node{
		Value: 5,
		Next: &trees.Node{
			Value: 7,
		},
	}
	lnkdlst.Add(&five)
	// fmt.Println(lnkdlst.Head.Next.Next)
	// fmt.Println(lnkdlst.Search(7).Value)

	bst := trees.BinaryTree{}
	bst.Insert(&trees.TreeNode{5, nil, nil})
	bst.Insert(&trees.TreeNode{3, nil, nil})
	bst.Insert(&trees.TreeNode{7, nil, nil})
	bst.Insert(&trees.TreeNode{2, nil, nil})
	bst.Insert(&trees.TreeNode{1, nil, nil})
	bst.Insert(&trees.TreeNode{10, nil, nil})
	bst.Insert(&trees.TreeNode{17, nil, nil})
	bst.Insert(&trees.TreeNode{3, nil, nil})
	bst.Insert(&trees.TreeNode{7, nil, nil})
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
	x := make([][]byte, 3)
	// x[0] = []byte{1, 2}
	// x[1] = []byte{1, 2, 3, 4}
	// fmt.Println(x[0])
	// fmt.Println(x[1])
	fmt.Println(x)
}
