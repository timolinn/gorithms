package trees

// Directed Acyclic Graph (DAG)
// This is a type of graph that has no cycles
// it is widely used in the blockchain space
// where they are used in-place of maybe linked-lists
// for distributed ledger data structures
type DAG struct {
	nodes map[int]*DAGNode
}

type DAGNode struct {
	ID       int
	Value    interface{}
	adjacent map[int]*DAGNode
}
