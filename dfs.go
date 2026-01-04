/*
Togosort-DFS is a library specifically made for package managers to detect
circular dependencies and sort dependencies from least to most dependent
Copyright (C) 2025-2026 ApertureOS Team
Licensed under the MIT license. See LICENSE file in the project root for details.
*/

package togosort

import (
	"fmt"
)

/*

Before you try to understand the code, please read the following explanation of
topological sorting and the conventions used in this package. It is crucial for
understanding how to use the code correctly. Read it even if you dont wanna develop
on this package, because it contains important information about dependency
management in general, understanding this will mean you can avoid many common pitfalls
when dealing with dependencies in your own projects and also you'll be able to understand
how to use this package properly, which is crucial for avoiding bugs and issues.
Remember in everything you do. YOU MUST ALWAYS UNDERSTAND what you are working with,
especially when it comes to dependencies. Blindly using code without understanding
it is a recipe for disaster. (same for vibe coding btw, dont vibecode if you dont know
what you are doing). Thanks, you may proceed, ill try to make this lesson as understandable
as possible since it was a pain in the ass for me to understand in the first place.

******** ALWAYS RUN DFS (DEPTH FIRST SEARCH) BEFORE TOPOSORTING ********

---------------------- TECHNICAL DEFINITIONS ----------------------

Graph:
    A collection of nodes (also called vertices) connected by directed edges.
    Example:
        A -> B -> C
    This means there is a direction to the relationships.

    In package managers and most real systems, we use this convention:
        Dependency -> Dependent

    Meaning:
        If B depends on A, the edge is:
            A -> B
        A must come BEFORE B.

Nodes (Vertices):
    Individual elements in a graph.
    In:
        A -> B -> C
    The nodes are A, B, and C.

Array:
    A data structure that holds items in a specific order.
    Often used to store nodes or results.

    Example (Go):
        deps := []string{"A", "B", "C"}

    Access:
        deps[0] == "A"

    Note:
        Go arrays/slices are 0-indexed.
        Lua is 1-indexed (unfortunately).

Edges:
    Directed connections between nodes.
    They represent relationships.

    Example:
        A -> B

Incoming Edge:
    An edge that points INTO a node.
    In:
        A -> B
    The edge A->B is an incoming edge for node B.

Dependency:
    A relationship where one node must be processed before another.

    Using the standard convention:
        Dependency -> Dependent

    Example:
        Algebra1 -> Algebra2
        Algebra2 depends on Algebra1

        libc -> bash
        bash depends on libc

    So:
        If X depends on Y, the edge is:
            Y -> X

---------------------- END OF TECHNICAL DEFINITIONS ----------------------

--------------------------- HOW DOES DFS WORK? ---------------------------
Let's declare some Arrays
ToBeVisited (the array you'll pass to the dfs() function)
Visiting
Visited
assuming the graph A -> B -> C,D-> F,G
A depends on B
B depends on C, D
C has no dependencies
D depends on F, G
F has no dependencies
G has no dependencies

In order to DFS this graph we would have to
Step 1: dfs(A)
- Move A from ToBeVisited to Visiting
  Visiting = [A]
- A depends on B -> dfs(B)

Step 2: dfs(B)
- Move B to Visiting
  Visiting = [A, B]
- B depends on C, D

Step 3: dfs(C)
- Move C to Visiting
  Visiting = [A, B, C]
- C has no dependencies
- Move C from Visiting to Visited
  Visited  = [C]
  Visiting = [A, B]

Step 4: dfs(D)
- Move D to Visiting
  Visiting = [A, B, D]
- D depends on F, G

Step 5: dfs(F)
- Move F to Visiting
  Visiting = [A, B, D, F]
- F has no dependencies
- Move F to Visited
  Visited  = [C, F]
  Visiting = [A, B, D]

Step 6: dfs(G)
- Move G to Visiting
  Visiting = [A, B, D, G]
- G has no dependencies
- Move G to Visited
  Visited  = [C, F, G]
  Visiting = [A, B, D]

Step 7: finish D
- All dependencies of D processed
- Move D to Visited
  Visited  = [C, F, G, D]
  Visiting = [A, B]

Step 8: finish B
- All dependencies of B processed
- Move B to Visited
  Visited  = [C, F, G, D, B]
  Visiting = [A]

Step 9: finish A
- All dependencies of A processed
- Move A to Visited
  Visited  = [C, F, G, D, B, A]

Final result:
Visited = [C, F, G, D, B, A] (dfs function will return this)
This order guarantees dependencies are handled before dependents.

*/

// DFS performs a depth-first search starting from the given roots.
// Returns:
//
//		error if there's a cycle.
//	 	if no error it returns nil
//
// HOW TO USE:
// Initialize a roots variable
// roots: the packages you want to check/install; DFS will traverse all dependencies
// reachable from these roots and return an error if any cycle is found. For a package manager it'd
// *** Probably *** be the arguments the user input, eg.
// $ sudo blink install gtk4, gtk5, bash, zsh, fish
// roots would be the packages, so gtk4, gtk5, bash, zsh, fish
// roots := []string{ ***packages here*** }
func (g *Graph) DFS(roots []string) error {

	visiting := make(map[string]bool)
	visited := make(map[string]bool)

	var dfs func(string) error
	dfs = func(node string) error {
		// Cycle detected
		if visiting[node] {
			return fmt.Errorf("cycle detected at %q", node)
		}

		// Already fully processed
		if visited[node] {
			return nil
		}

		// Enter node
		visiting[node] = true

		// Visit dependents
		for _, dep := range g.Edges[node] {
			if err := dfs(dep); err != nil {
				return err
			}
		}

		// Exit node
		visiting[node] = false
		visited[node] = true
		return nil
	}

	for _, root := range roots {
		if err := dfs(root); err != nil {
			return err
		}
	}

	return nil
}
