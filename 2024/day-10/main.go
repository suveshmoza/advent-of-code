package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Point struct {
	x, y int
}

var directions = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isValid(x, y int, grid [][]int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func dfs(x, y int, grid [][]int, visited [][]bool, reachableNines *map[Point]bool, currentHeight int, totalWays *int) {
	visited[x][y] = true

	if grid[x][y] == 9 {
		(*totalWays)++
		(*reachableNines)[Point{x, y}] = true
	}

	for _, dir := range directions {
		newX := x + dir.x
		newY := y + dir.y

		if isValid(newX, newY, grid) && !visited[newX][newY] && grid[newX][newY] == currentHeight+1 {
			dfs(newX, newY, grid, visited, reachableNines, currentHeight+1, totalWays)
		}
	}

	visited[x][y] = false
}

func countHikingTrails(grid [][]int) (int, int) {
	rows := len(grid)
	cols := len(grid[0])
	distinctPositions := 0
	totalWays := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				visited := make([][]bool, rows)
				for k := range visited {
					visited[k] = make([]bool, cols)
				}
				reachableNines := make(map[Point]bool)
				dfs(i, j, grid, visited, &reachableNines, 0, &totalWays)
				distinctPositions += len(reachableNines)
			}
		}
	}

	return distinctPositions, totalWays
}

func main() {
	data, err := os.Open("input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var matrix [][]int

	for scanner.Scan() {
		line := scanner.Text()
		temp := []int{}
		for _, char := range line {
			val, err := strconv.Atoi(string(char))
			check(err)
			temp = append(temp, val)
		}
		matrix = append(matrix, temp)
	}

	distinctPositions, waysCount := countHikingTrails(matrix)

	fmt.Println("Sum of scores of all trail heads:", distinctPositions)
	fmt.Println("Sum of ratings of all trail heads:", waysCount)

}
