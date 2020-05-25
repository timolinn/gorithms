package trees

import (
	"fmt"
)

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
	Vertices map[int]*Vertex
}

type Vertex struct {
	ID       int
	Value    interface{}
	Children map[int]*Vertex
}

func NewDAG() *DAG {
	return &DAG{
		Vertices: make(map[int]*Vertex),
	}
}

func NewVertex(id int, value interface{}) *Vertex {
	return &Vertex{
		ID:       id,
		Value:    value,
		Children: make(map[int]*Vertex, 0),
	}
}

func (dag *DAG) AddVertex(id int, value interface{}) error {
	if _, ok := dag.Vertices[id]; ok {
		return fmt.Errorf("vertex with %d already exist", id)
	}
	vertex := NewVertex(id, value)
	dag.Vertices[id] = vertex
	return nil
}

func (dag *DAG) getVertex(id int) *Vertex {
	if vertex, ok := dag.Vertices[id]; ok {
		return vertex
	}

	return nil
}

func (dag *DAG) AddEdge(from, to int) error {
	v1 := dag.getVertex(from)
	v2 := dag.getVertex(to)

	if v1 == nil || v2 == nil {
		return fmt.Errorf("missnig vertex")
	}
	// check if vertex already exists
	if _, ok := v1.Children[v2.ID]; ok {
		return fmt.Errorf("edge already exists")
	}

	v1.Children[v2.ID] = v2
	return nil
}
