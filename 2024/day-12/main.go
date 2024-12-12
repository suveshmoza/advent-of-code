package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Plot struct {
	area      int
	perimeter int
	side      int
}

func dfs(mat *[][]string, i, j int, label string, visited [][]bool, plot *Plot) {
	if i < 0 || j < 0 || i >= len(*mat) || j >= len((*mat)[0]) ||
		visited[i][j] || (*mat)[i][j] != label {
		return
	}

	visited[i][j] = true
	plot.area++

	if i == 0 || (*mat)[i-1][j] != label {
		plot.perimeter++
	}

	if j == 0 || (*mat)[i][j-1] != label {
		plot.perimeter++
	}

	if i == len(*mat)-1 || (*mat)[i+1][j] != label {
		plot.perimeter++
	}

	if j == len((*mat)[0])-1 || (*mat)[i][j+1] != label {
		plot.perimeter++
	}

	// Outside corners - when current plot meets two different plots at a corner
	// Top-left corner
	if (i > 0 && j > 0 && (*mat)[i-1][j] != label && (*mat)[i][j-1] != label) ||
		(i > 0 && j == 0 && (*mat)[i-1][j] != label) ||
		(i == 0 && j > 0 && (*mat)[i][j-1] != label) ||
		(i == 0 && j == 0) {
		plot.side++
	}

	// Top-right corner
	if (i > 0 && j < len((*mat)[0])-1 && (*mat)[i-1][j] != label && (*mat)[i][j+1] != label) ||
		(i > 0 && j == len((*mat)[0])-1 && (*mat)[i-1][j] != label) ||
		(i == 0 && j < len((*mat)[0])-1 && (*mat)[i][j+1] != label) ||
		(i == 0 && j == len((*mat)[0])-1) {
		plot.side++
	}

	// Bottom-left corner
	if (i < len(*mat)-1 && j > 0 && (*mat)[i+1][j] != label && (*mat)[i][j-1] != label) ||
		(i < len(*mat)-1 && j == 0 && (*mat)[i+1][j] != label) ||
		(i == len(*mat)-1 && j > 0 && (*mat)[i][j-1] != label) ||
		(i == len(*mat)-1 && j == 0) {
		plot.side++
	}

	// Bottom-right corner
	if (i < len(*mat)-1 && j < len((*mat)[0])-1 && (*mat)[i+1][j] != label && (*mat)[i][j+1] != label) ||
		(i < len(*mat)-1 && j == len((*mat)[0])-1 && (*mat)[i+1][j] != label) ||
		(i == len(*mat)-1 && j < len((*mat)[0])-1 && (*mat)[i][j+1] != label) ||
		(i == len(*mat)-1 && j == len((*mat)[0])-1) {
		plot.side++
	}

	// Inside corners - when current plot has two same-type neighbors that form a corner with a different plot
	// Top-left inside corner
	if i < len(*mat)-1 && j < len((*mat)[0])-1 &&
		(*mat)[i][j+1] == label && (*mat)[i+1][j] == label &&
		(*mat)[i+1][j+1] != label {
		plot.side++
	}

	// Top-right inside corner
	if i < len(*mat)-1 && j > 0 &&
		(*mat)[i][j-1] == label && (*mat)[i+1][j] == label &&
		(*mat)[i+1][j-1] != label {
		plot.side++
	}

	// Bottom-left inside corner
	if i > 0 && j < len((*mat)[0])-1 &&
		(*mat)[i][j+1] == label && (*mat)[i-1][j] == label &&
		(*mat)[i-1][j+1] != label {
		plot.side++
	}

	// Bottom-right inside corner
	if i > 0 && j > 0 &&
		(*mat)[i][j-1] == label && (*mat)[i-1][j] == label &&
		(*mat)[i-1][j-1] != label {
		plot.side++
	}

	dfs(mat, i-1, j, label, visited, plot)
	dfs(mat, i+1, j, label, visited, plot)
	dfs(mat, i, j-1, label, visited, plot)
	dfs(mat, i, j+1, label, visited, plot)
}

func main() {
	data, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var mat [][]string

	for scanner.Scan() {
		line := scanner.Text()
		mat = append(mat, strings.Split(line, ""))
	}

	var ans1, ans2 int

	visited := make([][]bool, len(mat))
	for i := range visited {
		visited[i] = make([]bool, len(mat[0]))
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if !visited[i][j] {
				plot := &Plot{area: 0, perimeter: 0}
				dfs(&mat, i, j, mat[i][j], visited, plot)
				ans1 += plot.area * plot.perimeter
				ans2 += plot.area * plot.side
			}
		}
	}

	fmt.Println("Answer for Part 1: ", ans1)
	fmt.Println("Answer for Part 2: ", ans2)
}
