package trees

import (
	"fmt"
	"log"
	"testing"
)

var graph *Graph
var v0 *GraphNode
var v1 *GraphNode
var v2 *GraphNode
var v3 *GraphNode
var v4 *GraphNode

func setup(kind GraphType) {
	graph = NewGraph(kind)
	v0 = graph.AddNode(3)
	v1 = graph.AddNode(2)
	v2 = graph.AddNode(1)
	v3 = graph.AddNode(6)
	v4 = graph.AddNode(7)
	graph.AddEdge(v4.ID, v0.ID)
	graph.AddEdge(v3.ID, v4.ID)
	graph.AddEdge(v1.ID, v2.ID)
	graph.AddEdge(v1.ID, v3.ID)
}

func TestHasPathBFS(t *testing.T) {
	setup(Undirected)
	tests := []struct {
		input1 int
		input2 int
		want   bool
	}{
		{v4.ID, v0.ID, true},
		{v0.ID, v4.ID, true},
		{v3.ID, v4.ID, true},
		{v4.ID, v3.ID, true},
		{v1.ID, v2.ID, true},
		{v2.ID, v1.ID, true},
		{v1.ID, v3.ID, true},
		{v3.ID, v1.ID, true},
		{v3.ID, v2.ID, true},
		{v2.ID, v4.ID, true},
		{v1.ID, v0.ID, true},
		{v1.ID, v4.ID, true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("v1=%v, v2=%v", tc.input1, tc.input2), func(t *testing.T) {
			got := graph.HasPathBFS(tc.input1, tc.input2)
			if got != tc.want {
				t.Fatalf("graph.HasPathBFS(): got=%v, want=%v", got, tc.want)
			}
		})
	}
}

func TestAddNode(t *testing.T) {
	graph := NewGraph(Directed)
	graph.AddNode(3)
	t.Run("size of graph should be exactly 1", func(t *testing.T) {
		if graph.Size() != 1 {
			log.Fatalf("graph size should be 1 got=%v", graph.Size())
		}
	})
}

func TestAddEdge(t *testing.T) {
	graph := NewGraph(Directed)
	v0 := graph.AddNode(3)
	v1 := graph.AddNode(5)
	t.Run("3 must be linked to 5", func(t *testing.T) {
		graph.AddEdge(v0.ID, v1.ID)
		if !graph.HasPathBFS(v0.ID, v1.ID) {
			log.Fatalf("path must exist between 3 and 5")
		}
	})
}
