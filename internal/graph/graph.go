package graph

import (
	"bufio"
	"fmt"
	"os"
)

type Graph struct {
	N        int
	Directed bool
	Adj      [][]bool
}

func New(n int, directed bool) *Graph {
	adj := make([][]bool, n)
	for i := range adj {
		adj[i] = make([]bool, n)
	}
	return &Graph{N: n, Directed: directed, Adj: adj}
}

func (g *Graph) AddEdge(u, v int) {
	g.Adj[u][v] = true
	if !g.Directed {
		g.Adj[v][u] = true
	}
}

func LoadEdgeList(path string, directed bool) (*Graph, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	edges := [][2]int{}
	maxV := -1
	for scanner.Scan() {
		var u, v int
		line := scanner.Text()
		if line == "" {
			continue
		}
		_, err := fmt.Sscan(line, &u, &v)
		if err != nil {
			return nil, fmt.Errorf("linha inválida: %q", line)
		}
		edges = append(edges, [2]int{u, v})
		if u > maxV {
			maxV = u
		}
		if v > maxV {
			maxV = v
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	g := New(maxV+1, directed)
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}
	return g, nil
}
