package graph

import (
	"fmt"
	"testing"
)

func TestGraphSearch(t *testing.T) {
	graph := newGraph(8)
	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 4)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 7)

	graph.BFS(0, 7)
	fmt.Println()
	graph.BFS(1, 3)
	fmt.Println()
	graph.DFS(0, 7)
	fmt.Println()
	graph.DFS(1, 3)
	fmt.Println()
}
