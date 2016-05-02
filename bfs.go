package main

import "fmt"

// breadth first search - algorithm that explores node.
func bfs(s string, adj map[string][]string) {

	frontier := []string{s}

	fmt.Println(frontier)

}

// relax returns node S's neighbors based on adjacency-list adj.
func relax(s string, adj map[string][]string) []string {
	return adj[s]
}

// filter filters candidate-list by excluding nodes from visited-list.
func filter(candidate []string, visited map[string]struct{}) []string {
	var next []string
	for _, u := range candidate {
		if _, ok := visited[u]; !ok {
			next = append(next, u)
			visited[u] = struct{}{}
		}
	}
	return next
}
