package trees

// Directed Acyclic Graph (DAG)
// This is a type of graph that has no cycles
// it is widely used in the blockchain space
// where they are used in-place of maybe linked-lists
// for distributed ledger data structures
// In order to support the ability to push
// and pull changesets between multiple instances of the same repository,
// we need a specially designed structure for representing multiple versions of things.
// The structure we use is called a Directed Acyclic Graph (DAG),
// a design which is more expressive than a purely linear model.
// The history of everything in the repository is modeled as a DAG. This sounds like Git.
type DAG struct {
	nodes map[int]*DAGNode
}

type DAGNode struct {
	ID       int
	Value    interface{}
	adjacent map[int]*DAGNode
}
