package trees

// TreeNode represents a node in a tree
// data structure
type TreeNode struct {
	value    interface{}
	children []*TreeNode
}

func NewTreeNode(value interface{}) *TreeNode {
	return &TreeNode{
		value:    value,
		children: []*TreeNode{},
	}
}

func (t *TreeNode) Add(value interface{}) {
	t.children = append(t.children, NewTreeNode(value))
}
