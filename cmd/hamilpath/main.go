package main

import (
	"flag"
	"fmt"
	"log"

	"hamiltonian-path/internal/graph"
	"hamiltonian-path/internal/hamil"
)

func main() {
	file := flag.String("file", "", "arquivo de arestas (u v por linha)")
	directed := flag.Bool("directed", false, "grafo dirigido")
	start := flag.Int("start", -1, "vértice inicial (opcional)")
	flag.Parse()

	if *file == "" {
		log.Fatal("use -file para informar o arquivo de arestas")
	}

	gph, err := graph.LoadEdgeList(*file, *directed)
	if err != nil {
		log.Fatal(err)
	}

	if path, ok := hamil.HamiltonianPath(gph, *start); ok {
		fmt.Println("FOUND")
		for i, v := range path {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	} else {
		fmt.Println("NOT FOUND")
	}
}
