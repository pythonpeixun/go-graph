package main

// bfs explores graph(adj) based on initial node. In the process,
// shortest-path reachable from vertex S will be collected: parent, and
// number of moves reachable from vertex S will be collected: level as well.
func bfs(s string, adj map[string][]string) {

	frontier := []string{s}
	level := make(map[string]int) // visited
	parent := make(map[string]string)
	i := 0
	level[s] = i
	for len(frontier) != 0 {
		for _, u := range frontier {

			next := []string{}
			for _, v := range adj[u] {
				if _, exist := level[v]; !exist {
					i++
					level[v] = i
					next = append(next, v)
					parent[v] = u
				}
			}
			frontier = next
		}

	}

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

//remove removes string:s from slice:a.
func remove(a []string, s string) []string {

	for i := range a {
		if a[i] == s {
			return append(a[:i], a[i+1:]...)
		}
	}
	return a
}

//pop pops(filo) an element from slice:a
func pop(a []string) (string, []string) {
	return a[len(a)-1], a[:len(a)-1]
}
