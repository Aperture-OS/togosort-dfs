/*
  Togosort-DFS is a library specifically made for package managers to detect
  circular dependencies and sort dependencies from least to most dependent
  Copyright (C) 2025-2026 ApertureOS Team
  Licensed under the MIT license. See LICENSE file in the project root for details.
*/

package togosort

import ()

// Graph uses the convention:
//
//	dependency -> dependent
//
// If B depends on A:
//
//	A -> B
type Graph struct {
	Edges map[string][]string
}

// NewGraph initializes an empty graph
// HOW TO USE:
// just make a new graph.
// graph := togosort.NewGraph()
// from now on graph is your graph variable. Use this with the other functions of this library.
func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[string][]string), // edges would look like "a": "b","d","...and so on"
		//					   "c": "d"
	}
}

// AddNode ensures a node exists
// HOW TO USE:
// You do NOT normally call this!
// You only use it if: a package has no dependencies but you still want it to appear in the graph
// example, packages without any dependencies, you could just ignore them, or sort them aswell, your choice ;)
// graph.AddNode("Node Name here")
func (g *Graph) AddNode(node string) {
	if _, ok := g.Edges[node]; !ok {
		g.Edges[node] = []string{}
	}
}

// AddEdge adds a dependency edge:
//
//	dependency -> dependent
//
// HOW TO USE:
// graph.AddEdge("dependency", "dependent")
// to clear some doubts, if a depends on multiple dependencies such as a -> b,c,d
// you're gonna have to call this multiple times, like this.
// graph.AddEdge("a", "b")
// graph.AddEdge("a", "c")
// graph.AddEdge("a", "d")
// TIP: automate this with a loop
func (g *Graph) AddEdge(dependency, dependent string) {
	g.AddNode(dependency)
	g.AddNode(dependent)
	g.Edges[dependency] = append(g.Edges[dependency], dependent)
}
