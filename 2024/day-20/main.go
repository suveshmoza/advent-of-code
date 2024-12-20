package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

var directions = []Position{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

// Find positions of 'S' (Start) and 'E' (End) in the grid.
func findPositions(grid [][]rune) (Position, Position, bool) {
	var start, end Position
	foundStart, foundEnd := false, false
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				start = Position{i, j}
				foundStart = true
			} else if cell == 'E' {
				end = Position{i, j}
				foundEnd = true
			}
		}
	}
	return start, end, foundStart && foundEnd
}

// Check if the cell is a valid track cell.
func isTrack(cell rune) bool {
	return cell == '.' || cell == 'S' || cell == 'E'
}

// Breadth-first search for pathfinding.
func bfs(grid [][]rune, start Position, end *Position, isValidCell func(rune) bool, ignoreWalls bool) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]int, rows)
	for i := range visited {
		visited[i] = make([]int, cols)
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}

	queue := []Position{start}
	visited[start.x][start.y] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if end != nil && current == *end {
			break
		}

		for _, dir := range directions {
			nx, ny := current.x+dir.x, current.y+dir.y
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
				if isValidCell != nil {
					if !isValidCell(grid[nx][ny]) {
						continue
					}
				} else if !ignoreWalls && grid[nx][ny] == '#' {
					continue
				}
				if visited[nx][ny] == -1 || visited[nx][ny] > visited[current.x][current.y]+1 {
					visited[nx][ny] = visited[current.x][current.y] + 1
					queue = append(queue, Position{nx, ny})
				}
			}
		}
	}
	return visited
}

// Reconstruct path from BFS visited data.
func reconstructPath(visited [][]int, start, end Position) []Position {
	path := []Position{}
	x, y := end.x, end.y
	step := visited[x][y]
	if step == -1 {
		return nil
	}
	path = append(path, end)
	currentStep := step

	for x != start.x || y != start.y {
		foundPrev := false
		for _, dir := range directions {
			nx, ny := x+dir.x, y+dir.y
			if nx >= 0 && nx < len(visited) && ny >= 0 && ny < len(visited[0]) && visited[nx][ny] == currentStep-1 {
				path = append(path, Position{nx, ny})
				x, y = nx, ny
				currentStep--
				foundPrev = true
				break
			}
		}
		if !foundPrev {
			return nil
		}
	}
	// Reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func solvePart1(grid [][]rune, start, end Position) {
	visited := bfs(grid, start, &end, func(ch rune) bool { return ch != '#' }, false)
	T := visited[end.x][end.y]
	if T == -1 {
		fmt.Println("No path exists from Start to End.")
		return
	}
	path := reconstructPath(visited, start, end)
	if path == nil {
		fmt.Println("No path reconstructed.")
		return
	}
	stepNum := make(map[Position]int)
	for idx, pos := range path {
		stepNum[pos] = idx
	}
	cheatCount := 0
	rows, cols := len(grid), len(grid[0])

	for idx, a := range path {
		targetIdx := idx + 102
		if targetIdx >= len(path) {
			break
		}
		for _, d1 := range directions {
			for _, d2 := range directions {
				bx := a.x + d1.x + d2.x
				by := a.y + d1.y + d2.y
				if bx >= 0 && bx < rows && by >= 0 && by < cols {
					b := Position{bx, by}
					if step, exists := stepNum[b]; exists && step-idx >= 102 {
						cheatCount++
					}
				}
			}
		}
	}
	fmt.Println(cheatCount)
}

func solvePart2(grid [][]rune, start, end Position) {
	rows, cols := len(grid), len(grid[0])
	distFromStart := bfs(grid, start, nil, isTrack, false)
	distFromEnd := bfs(grid, end, nil, isTrack, false)

	normalDist := distFromStart[end.x][end.y]
	if normalDist == -1 {
		fmt.Println(0)
		return
	}

	cheats := make(map[[4]int]struct{})
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if distFromStart[x][y] != -1 && isTrack(grid[x][y]) {
				distIgnoreWalls := bfs(grid, Position{x, y}, nil, func(rune) bool { return true }, true)
				for nx := 0; nx < rows; nx++ {
					for ny := 0; ny < cols; ny++ {
						d := distIgnoreWalls[nx][ny]
						if d != -1 && d >= 1 && d <= 20 && isTrack(grid[nx][ny]) && distFromEnd[nx][ny] != -1 {
							routeWithCheat := distFromStart[x][y] + d + distFromEnd[nx][ny]
							if normalDist-routeWithCheat >= 100 {
								cheats[[4]int{x, y, nx, ny}] = struct{}{}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(len(cheats))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	start, end, found := findPositions(grid)
	if !found {
		fmt.Println("Start or End not found in the grid.")
		return
	}

	solvePart1(grid, start, end)
	solvePart2(grid, start, end)
}
