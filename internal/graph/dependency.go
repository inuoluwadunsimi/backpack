// Package graph provides dependency-aware ordering for restore operations.
// Tools have install dependencies (e.g. npm depends on Node, which may be
// installed via Homebrew), so restores must happen in topological order.
package graph

import "fmt"

// Node represents a tool in the dependency graph.
type Node struct {
	Name string
	Deps []string // names of tools this tool depends on
}

// DefaultDependencies returns the known dependency relationships.
// This is a static graph — we don't auto-discover deps at runtime.
func DefaultDependencies() []Node {
	return []Node{
		{Name: "system", Deps: nil},                      // no deps, just info
		{Name: "homebrew", Deps: []string{"system"}},      // needs Xcode CLI tools
		{Name: "shell", Deps: []string{"homebrew"}},       // oh-my-zsh etc. may come from brew
		{Name: "npm", Deps: []string{"homebrew"}},         // node via brew
		{Name: "pip", Deps: []string{"homebrew"}},         // python via brew
		{Name: "vscode", Deps: []string{"homebrew"}},      // `code` CLI via brew cask
	}
}

// ResolveOrder returns tool names in a valid installation order (topological sort).
// Returns an error if a cycle is detected.
func ResolveOrder(nodes []Node) ([]string, error) {
	// Build adjacency + in-degree
	inDegree := make(map[string]int)
	dependents := make(map[string][]string) // dep -> list of tools that depend on it

	for _, n := range nodes {
		if _, ok := inDegree[n.Name]; !ok {
			inDegree[n.Name] = 0
		}
		for _, dep := range n.Deps {
			dependents[dep] = append(dependents[dep], n.Name)
			inDegree[n.Name]++
			if _, ok := inDegree[dep]; !ok {
				inDegree[dep] = 0
			}
		}
	}

	// Kahn's algorithm
	var queue []string
	for name, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, name)
		}
	}

	var order []string
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		order = append(order, current)

		for _, dependent := range dependents[current] {
			inDegree[dependent]--
			if inDegree[dependent] == 0 {
				queue = append(queue, dependent)
			}
		}
	}

	if len(order) != len(inDegree) {
		return nil, fmt.Errorf("dependency cycle detected")
	}

	return order, nil
}
