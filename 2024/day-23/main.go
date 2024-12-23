package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var adj = make(map[string]map[string]bool)

func buildGraph(connections []string) {
	for _, line := range connections {
		parts := strings.Split(line, "-")
		u := parts[0]
		v := parts[1]

		if adj[u] == nil {
			adj[u] = make(map[string]bool)
		}
		if adj[v] == nil {
			adj[v] = make(map[string]bool)
		}

		adj[u][v] = true
		adj[v][u] = true
	}
}

func readInputFromFile(filename string) ([]string, error) {
	var connections []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		connections = append(connections, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}

func findTriangles() [][]string {
	triangles := make(map[string][]string)

	for node, neighbors := range adj {
		for n1 := range neighbors {
			for n2 := range neighbors {
				if _, ok := adj[n1][n2]; ok {
					nodes := []string{node, n1, n2}
					sort.Strings(nodes)
					key := strings.Join(nodes, ",")
					triangles[key] = nodes
				}
			}
		}
	}

	var result [][]string
	for _, triangle := range triangles {
		result = append(result, triangle)
	}

	return result
}

func countTrianglesWithT(triangles [][]string) int {
	count := 0
	for _, triangle := range triangles {
		for _, node := range triangle {
			if strings.HasPrefix(node, "t") {
				count++
				break
			}
		}
	}
	return count
}

func isClique(nodes []string) bool {
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if !adj[nodes[i]][nodes[j]] {
				return false
			}
		}
	}
	return true
}

func findLargestClique() []string {
	var largestClique []string
	nodes := make([]string, 0, len(adj))
	for node := range adj {
		nodes = append(nodes, node)
	}

	n := len(nodes)
	for size := 3; size <= n; size++ {
		v := make([]bool, n)
		for i := 0; i < size; i++ {
			v[i] = true
		}

		for {
			var clique []string
			for i := 0; i < n; i++ {
				if v[i] {
					clique = append(clique, nodes[i])
				}
			}

			if isClique(clique) && len(clique) > len(largestClique) {
				largestClique = clique
			}

			i := n - 1
			for i >= 0 && !v[i] {
				i--
			}
			if i < 0 {
				break
			}
			v[i] = false
			for j := i + 1; j < n; j++ {
				v[j] = true
			}
		}
	}

	return largestClique
}

func main() {
	filename := "test.txt"
	connections, err := readInputFromFile(filename)
	if err != nil {
		panic(err)
	}

	buildGraph(connections)

	triangles := findTriangles()

	result := countTrianglesWithT(triangles)
	fmt.Println("Part 1", result)

	largestClique := findLargestClique()

	sort.Strings(largestClique)
	password := strings.Join(largestClique, ",")

	fmt.Println("Part 2:", password)
}
