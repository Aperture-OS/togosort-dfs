/*
  Togosort-DFS is a library specifically made for package managers to detect
  circular dependencies and sort dependencies from least to most dependent
  Copyright (C) 2025-2026 ApertureOS Team
  Licensed under the MIT license. See LICENSE file in the project root for details.
*/

package togosort

import ()

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


---------------------- HOW TO TOPOLOGICALLY SORT A GRAPH ----------------------

We assume the standard convention:
    Dependency -> Dependent

1. Count how many dependencies each node has

    Example: a->b, b->c, c->d
             in-degree: a=0, b=1, c=1, d=1

2. Grab all nodes with no dependencies

    Example: Only 'a' has 0 dependencies, so queue = [a]

3. Take a node, add it to the result, and reduce its dependents' counts

    Example: Take 'a', result=[a]
             b depends on a, so reduce b's count: b=0

4. If a dependent now has no dependencies, add it to the queue

    Example: b now has 0, add to queue: queue=[b]

5. Reverse the result so least dependent comes first

    Example: Result before reverse=[a,b,c,d]
             After reverse=[d,c,b,a]

If you don't get how it makes sense, think of it this way.

A depends on B, so B needs to get installed before A for A to function
B depends on C, so C needs to get installed before B for B to function
C depends on D, so D needs to get installed before C for C to function
D has no dependencies

we gotta install A, so first we gotta install B, so first we gotta install C, so first we gotta install D.
D Installed
We gotta install A, so first we gotta install B, so first we gotta install C, D is already satisfied (AKA already installed), so we install C
C Installed
We gotta install A, so first we gotta install B, C already satisfied, so we install B
B Installed
We gotta install A, B is satisfied, so we install A
A installed.
Done!

Order is D, C, B, A

Notice how we installed the Packages with less dependencies first
Dependencies list:
A: 3, the following
B: 2, the following
C: 1, the following
D: 0

From less to most dependencies, 0,1,2,3: D,C,B,A

*/

// TopoSort returns nodes from least dependent -> most dependent.
// Assumes you already checked for cycles with DFS.
// HOW TO USE:
// after you made your graph variable
// sorted := graph.TopoSort()
// sorted now contains the sorted array of packages/nodes
// it is a array, shows itself like this [d c b a].
// 'd c b a' will be different depending on your situation
func (g *Graph) TopoSort() []string {
	// Count how many dependencies each node has
	inDegree := make(map[string]int)
	for node := range g.Edges {
		if _, ok := inDegree[node]; !ok {
			inDegree[node] = 0
		}
		for _, dep := range g.Edges[node] {
			inDegree[dep]++
		}
	}

	// Grab all nodes with no dependencies
	queue := []string{}
	for node, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}

	result := []string{}

	// Keep taking nodes with no dependencies
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		// Remove their effect from dependents
		for _, dep := range g.Edges[node] {
			inDegree[dep]--
			if inDegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	// Flip it so least dependent comes first
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
