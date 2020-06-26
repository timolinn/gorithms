package trees

import (
	"github.com/timolinn/gorithms/datastructures"
)

// BSTreeNode represents one node in a tree
type BSTreeNode struct {
	Value int
	Left  *BSTreeNode
	Right *BSTreeNode
}

type BinaryTree struct {
	Root *BSTreeNode
}

var items []int

func NewBSTreeNode(value int) *BSTreeNode {
	return &BSTreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (bst *BinaryTree) Sum() int {
	if bst == nil {
		return 0
	}

	return bst.Root.Value + bst.Root.Left.Sum() + bst.Root.Right.Sum()
}

func (bst *BSTreeNode) Sum() int {
	if bst == nil {
		return 0
	}

	return bst.Value + bst.Left.Sum() + bst.Right.Sum()
}

func (bst *BinaryTree) Insert(e *BSTreeNode) *BSTreeNode {
	if bst.Root == nil {
		bst.Root = e
		return e
	}

	if bst.Root.Value >= e.Value {
		return bst.InsertLeft(e, bst.Root)
	}
	return bst.InsertRight(e, bst.Root)
}

// InsertLeft inserts a new node left of at
func (bst BinaryTree) InsertLeft(n, at *BSTreeNode) *BSTreeNode {
	if at.Left == nil {
		at.Left = n
		return at.Left
	}
	if at.Left.Value >= n.Value {
		return bst.InsertLeft(n, at.Left)
	}
	return bst.InsertRight(n, at.Left)
}

// InsertRight inserts a new node right of at
func (bst BinaryTree) InsertRight(n, at *BSTreeNode) *BSTreeNode {
	if at.Right == nil {
		at.Right = n
		return at.Right
	}
	if at.Right.Value >= n.Value {
		return bst.InsertLeft(n, at.Right)
	}
	return bst.InsertRight(n, at.Right)
}

// InOrderTraversal performs InOrder traversal
func (bst *BinaryTree) InOrderTraversal() []int {
	var items []int
	inOrderTreeTraversal(bst.Root, func(n int) {
		items = append(items, n)
	})
	return items
}

func (bst *BinaryTree) InOrderTraversalC(res chan int) {
	inOrderTreeTraversalCh(bst.Root, res)
	close(res)
}

func inOrderTreeTraversal(n *BSTreeNode, f func(int)) {
	if n == nil {
		return
	}
	f(n.Value)
	inOrderTreeTraversal(n.Left, f)
	inOrderTreeTraversal(n.Right, f)
}

func inOrderTreeTraversalCh(n *BSTreeNode, ch chan int) {
	if n == nil {
		return
	}
	ch <- n.Value
	inOrderTreeTraversalCh(n.Left, ch)
	inOrderTreeTraversalCh(n.Right, ch)
}

// PreOrderTraversal performs preorder traversal
func (bst *BinaryTree) PreOrderTraversal() []int {
	var items []int
	preOrderTraversal(bst.Root, func(i int) {
		items = append(items, i)
	})
	return items
}

func preOrderTraversal(n *BSTreeNode, f func(int)) {
	if n == nil {
		return
	}
	preOrderTraversal(n.Left, f)
	f(n.Value)
	preOrderTraversal(n.Right, f)
}

// PostOrderTraversal performs post order traversal
func (bst *BinaryTree) PostOrderTraversal() []int {
	var items []int
	postOrderTraversal(bst.Root, func(i int) {
		items = append(items, i)
	})
	return items
}

func postOrderTraversal(n *BSTreeNode, f func(int)) {
	if n == nil {
		return
	}
	postOrderTraversal(n.Left, f)
	postOrderTraversal(n.Right, f)
	f(n.Value)
}

func (bst *BinaryTree) Compare(b *BinaryTree) bool {
	if bst.Root == nil || b.Root == nil {
		return false
	}
	c1, c2 := make(chan int), make(chan int)
	go bst.InOrderTraversalC(c1)
	go b.InOrderTraversalC(c2)
	// go inOrderTreeTraversalCh(b1.Root, c1)
	// go inOrderTreeTraversalCh(b2.Root, c2)
	for v1 := range c1 {
		v2, ok := <-c2
		if !ok {
			return false
		}

		if v1 != v2 {
			return false
		}
	}
	return true
}

func Compare(b1, b2 *BinaryTree) bool {
	if b1.Root == nil || b2.Root == nil {
		return false
	}
	c1, c2 := make(chan int), make(chan int)
	go b1.InOrderTraversalC(c1)
	go b2.InOrderTraversalC(c2)
	// go inOrderTreeTraversalCh(b1.Root, c1)
	// go inOrderTreeTraversalCh(b2.Root, c2)
	for v1 := range c1 {
		v2, ok := <-c2
		if !ok {
			return false
		}

		if v1 != v2 {
			return false
		}
	}
	return true
}

// Remove removes n from bst
func (bst *BinaryTree) Remove(n *BSTreeNode) *BSTreeNode {
	return remove(bst.Root, n)
}

func remove(root, n *BSTreeNode) *BSTreeNode {
	if n == nil {
		return nil
	}

	if root.Value == n.Value {

	}

	if root.Value > n.Value {
		root.Left = remove(root.Left, n)
		return root
	}

	if root.Value < n.Value {
		root.Right = remove(root.Right, n)
		return root
	}
	return root
}

func swap(n, with *BSTreeNode) {
	if n == nil {
		n = with
		return
	}
	temp := n
	n = with
	n.Right = temp.Right
	n.Left = with.Right
	// n.Left.Right = with.Right
	if with.Left != nil && n.Left != nil {
		swap(n.Left.Left, with.Left)
	} else if n.Left != nil && n.Left.Left == nil {
		n.Left.Left = with.Left
		return
	}
}

// Search returns a node with val and the parent
// to that node
func (bst *BinaryTree) Search(val int) (*BSTreeNode, *BSTreeNode) {
	if val < 0 {
		return nil, nil
	}
	return search(bst.Root, bst.Root, val)
}

func search(root, parent *BSTreeNode, val int) (*BSTreeNode, *BSTreeNode) {
	if root.Value == val {
		return root, parent
	}
	if val < root.Value {
		// go left
		return search(root.Left, root, val)
	}
	// go right
	return search(root.Right, root, val)
}

func BSTBreadthFirstSearch(bst *BinaryTree) []int {
	var res []int
	queue := datastructures.Queue{}
	queue.Enqueue(bst.Root)

	for !queue.IsEmpty() {
		n, _ := queue.Dequeue()
		node := n.(*BSTreeNode)
		res = append(res, node.Value)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
	return res
}
