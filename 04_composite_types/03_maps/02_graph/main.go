package main

import (
	"fmt"
)

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	graph[from][to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("A", "B")
	addEdge("B", "C")
	addEdge("C", "D")
	addEdge("D", "A")

	fmt.Printf("From:%v to:%v -> %v\n", "A", "B", hasEdge("A", "B"))
	fmt.Printf("From:%v to:%v -> %v\n", "B", "D", hasEdge("B", "D"))
}
