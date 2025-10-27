package hamil

import "hamiltonian-path/internal/graph"


func HamiltonianPath(g *graph.Graph, start int) ([]int, bool) {
	n := g.N
	visited := make([]bool, n)
	path := make([]int, n)

	tryFrom := func(s int) ([]int, bool) {
		for i := range visited {
			visited[i] = false
		}
		path[0] = s
		visited[s] = true
		if dfs(g, 1, visited, path) {
			return append([]int(nil), path...), true
		}
		return nil, false
	}

	if start >= 0 {
		return tryFrom(start)
	}
	for s := 0; s < n; s++ {
		if p, ok := tryFrom(s); ok {
			return p, true
		}
	}
	return nil, false
}

func dfs(g *graph.Graph, pos int, visited []bool, path []int) bool {
	if pos == g.N {
		return true
	}
	last := path[pos-1]
	for v := 0; v < g.N; v++ {
		if visited[v] {
			continue
		}
		if !g.Adj[last][v] {
			continue
		}
		visited[v] = true
		path[pos] = v
		if dfs(g, pos+1, visited, path) {
			return true
		}
		visited[v] = false
	}
	return false
}
