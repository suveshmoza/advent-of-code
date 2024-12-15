package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

func printMat(mat [][]string) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			fmt.Print(mat[i][j])
		}
		fmt.Println()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func expandGrid(input [][]string) (finalGrid [][]string, start Position) {
	for j, row := range input {
		finalRow := []string{}
		for i, char := range row {
			switch char {
			case ".":
				finalRow = append(finalRow, []string{".", "."}...)
			case "#":
				finalRow = append(finalRow, []string{"#", "#"}...)
			case "O":
				finalRow = append(finalRow, []string{"[", "]"}...)
			case "@":
				start.x = i * 2
				start.y = j
				finalRow = append(finalRow, []string{"@", "."}...)
			}

		}
		finalGrid = append(finalGrid, finalRow)
	}
	return finalGrid, start
}

func generateNextPos(robotPos Position, op string) Position {
	switch op {
	case "^":
		return Position{robotPos.x - 1, robotPos.y}
	case "v":
		return Position{robotPos.x + 1, robotPos.y}
	case "<":
		return Position{robotPos.x, robotPos.y - 1}
	case ">":
		return Position{robotPos.x, robotPos.y + 1}
	}
	return Position{-1, -1}
}

func isValidMove(mat [][]string, pos Position) bool {
	return pos.x >= 0 && pos.x < len(mat) && pos.y >= 0 && pos.y < len(mat[0]) && mat[pos.x][pos.y] != "#"
}
func moveRobot(mat [][]string, robotPos *Position, op string) bool {
	nextPos := generateNextPos(*robotPos, op)
	if !isValidMove(mat, nextPos) {
		return false
	}

	if (mat)[nextPos.x][nextPos.y] == "." {
		mat[robotPos.x][robotPos.y] = "."
		mat[nextPos.x][nextPos.y] = "@"
		*robotPos = nextPos
		return true
	}

	if (mat)[nextPos.x][nextPos.y] == "[" || (mat)[nextPos.x][nextPos.y] == "]" {
		nextDotOrWall := findNextDotOrWall(mat, nextPos, op)
		if nextDotOrWall.x == -1 && nextDotOrWall.y == -1 {
			return false
		}

		// Move the boxes
		moveAllRobots(mat, nextPos, op, nextDotOrWall)
		mat[robotPos.x][robotPos.y] = "."
		*robotPos = nextPos
		return true
	}

	return false
}

func moveAllRobots(mat [][]string, start Position, op string, endPos Position) {
	if op == ">" {
		for j := endPos.y; j >= start.y; j-- {
			mat[endPos.x][j] = "O"
			mat[endPos.x][j-1] = "."
		}
		mat[start.x][start.y] = "@"

	} else if op == "<" {
		for j := endPos.y; j <= start.y; j++ {
			mat[endPos.x][j] = "O"
			mat[endPos.x][j+1] = "."
		}
		mat[start.x][start.y] = "@"

	} else if op == "^" {
		for i := endPos.x; i <= start.x; i++ {
			mat[i][endPos.y] = "O"
			mat[i+1][endPos.y] = "."
		}
		mat[start.x][start.y] = "@"

	} else if op == "v" {
		for i := endPos.x; i >= start.x; i-- {
			mat[i][endPos.y] = "O"
			mat[i-1][endPos.y] = "."
		}
		mat[start.x][start.y] = "@"

	}
}

func findNextDotOrWall(mat [][]string, pos Position, dir string) Position {
	if dir == "^" {
		for i := pos.x - 1; i >= 0; i-- {
			if mat[i][pos.y] == "." {
				return Position{i, pos.y}
			}
			if mat[i][pos.y] == "#" {
				return Position{-1, -1}
			}
		}
	}
	if dir == "v" {
		for i := pos.x + 1; i < len(mat); i++ {
			if mat[i][pos.y] == "." {
				return Position{i, pos.y}
			}
			if mat[i][pos.y] == "#" {
				return Position{-1, -1}
			}
		}
	}
	if dir == "<" {
		for j := pos.y - 1; j >= 0; j-- {
			if mat[pos.x][j] == "." {
				return Position{pos.x, j}
			}
			if mat[pos.x][j] == "#" {
				return Position{-1, -1}
			}
		}
	}
	if dir == ">" {
		for j := pos.y + 1; j < len(mat[0]); j++ {
			if mat[pos.x][j] == "." {
				return Position{pos.x, j}
			}
			if mat[pos.x][j] == "#" {
				return Position{-1, -1}
			}
		}
	}
	return Position{-1, -1}
}

func calcAnsForPart1(mat [][]string) int {
	sum := 0

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == "O" {
				sum += ((100 * i) + j)
			}
		}
	}

	return sum
}

func main() {
	data, err := os.Open("./input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var mat [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		data := strings.Split(line, "")
		mat = append(mat, data)
	}

	var ops []string
	for scanner.Scan() {
		line := scanner.Text()
		separated := strings.Split(line, "")
		ops = append(ops, separated...)
	}

	var robotPos Position
	found := false
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == "@" {
				robotPos = Position{x: i, y: j}
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	for _, op := range ops {
		moveRobot(mat, &robotPos, op)
	}

	printMat(mat)
	fmt.Println("Part 1", calcAnsForPart1(mat))

}
