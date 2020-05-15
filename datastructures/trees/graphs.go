package trees

import (
	"fmt"

	"github.com/timolinn/gorithms/datastructures"
)

/**
Refrence: https://github.com/twmb/algoimpl/blob/master/go/graph/graph.go
*/

type GraphType int

const (
	Undirected GraphType = iota
	Directed
	DirectedWeighted
)

type edge struct {
	weight int
	end    *GraphNode
}

// GraphNode also known as Vertex in graph theory
type GraphNode struct {
	// edges         []edge
	// reversedEdges []edge
	ID       int
	adjacent []*GraphNode
	Value    *interface{}
}

// type Graph struct {
// 	nodes []*GraphNode
// 	Kind  GraphType
// }

type Graph struct {
	nodes map[int]*GraphNode
	Kind  GraphType
}

func (g *Graph) newGraphNode(value *interface{}) *GraphNode {
	return &GraphNode{
		ID:       len(g.nodes),
		Value:    value,
		adjacent: make([]*GraphNode, 0),
	}
}

// NewGraph constructs a new Graph type
func NewGraph(kind GraphType) *Graph {
	g := Graph{
		nodes: make(map[int]*GraphNode, 0),
		Kind:  kind,
	}
	return &g
}

// AddNode adds a new node or vertex to the graph
func (g *Graph) AddNode(value interface{}) *GraphNode {
	node := g.newGraphNode(&value)
	g.nodes[node.ID] = node
	return node
}

// AddEdge adds connection between two nodes
func (g *Graph) AddEdge(source, destination int) error {
	s := g.getNode(source)
	d := g.getNode(destination)

	if s == nil || d == nil {
		fmt.Println("invalid node")
		return fmt.Errorf("invalid source or destination ID")
	}

	s.adjacent = append(s.adjacent, d)
	if g.Kind == Undirected {
		d.adjacent = append(d.adjacent, s)
	}

	return nil
}

// RemoveEdge removes an edge between two vertecies
func (g *Graph) RemoveEdge(start, end int) {
	v1 := g.getNode(start)
	v2 := g.getNode(end)

	for index, v := range v1.adjacent {
		if v.Value == v2.Value {
			v1.adjacent[index] = v1.adjacent[len(v1.adjacent)-1]
			v1.adjacent = v1.adjacent[:len(v1.adjacent)-1]
		}
	}

	if g.Kind == Undirected {
		for index, v := range v2.adjacent {
			if v.Value == v1.Value {
				v2.adjacent[index] = v2.adjacent[len(v2.adjacent)-1]
				v2.adjacent = v2.adjacent[:len(v2.adjacent)-1]
			}
		}
	}
}

func (g *Graph) getNode(id int) *GraphNode {
	return g.nodes[id]
}

// Contains returns true if a value exists in a graph
// otherwise return false
func (g *Graph) Contains(value interface{}) bool {
	queue := datastructures.Queue{}
	visited := make(map[int]struct{}, 0)
	// you can use concurrency here
	// that is by
	for _, vertex := range g.nodes {
		queue.Enqueue(vertex)
	}

	for !queue.IsEmpty() {
		v, _ := queue.Dequeue()
		vertex := v.(*GraphNode)
		if _, ok := visited[vertex.ID]; ok {
			continue
		}
		visited[vertex.ID] = struct{}{}
		if *vertex.Value == value {
			return true
		}
	}
	return false
}

// HasPathDFS uses depth first search to check whether
// there is an existing path from node to node1
//
// use a Set-like datastructure to keep track of visited nodes
// to prevent infinite loop
func (g *Graph) HasPathDFS(node, node1 int) bool {
	v1 := g.getNode(node)
	v2 := g.getNode(node1)
	visited := make(map[int]struct{})
	return hasPathDFS(v1, v2, visited)
}

// HasPathBFS uses breadth first search to check whether
// there is an existing path from node to node1
func (g *Graph) HasPathBFS(node, node1 int) bool {
	return hasPathBFS(g.getNode(node), g.getNode(node1))
}

func hasPathBFS(source, destination *GraphNode) bool {
	visited := make(map[int]struct{})
	queue := datastructures.Queue{}
	queue.Enqueue(source)
	for !queue.IsEmpty() {
		v, _ := queue.Dequeue()
		vertex := v.(*GraphNode)
		if vertex == destination {
			return true
		}
		if _, ok := visited[vertex.ID]; ok {
			continue
		}
		visited[vertex.ID] = struct{}{}

		for _, child := range vertex.adjacent {
			queue.Enqueue(child)
		}
	}
	return false
}

func hasPathDFS(source, destination *GraphNode, visited map[int]struct{}) bool {
	if _, ok := visited[source.ID]; ok {
		return false
	}
	visited[source.ID] = struct{}{}
	if source == destination {
		return true
	}

	for _, child := range source.adjacent {
		if hasPathDFS(child, destination, visited) {
			return true
		}
	}
	return false
}

func (g *Graph) Size() int {
	return len(g.nodes)
}
