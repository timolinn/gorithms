package trees

type NodeType int

const (
	// TerminatingTrieNode the "*" node
	TerminatingTrieNode NodeType = iota
	DataNode
)

type TrieNode struct {
	char     string
	children []*TrieNode
	NodeType
}
