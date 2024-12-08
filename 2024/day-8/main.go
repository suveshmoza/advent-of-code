package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(filename string) [][]string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var matrix [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix
}

func groupCharacters(matrix [][]string) map[string][]Position {
	charMap := make(map[string][]Position)

	for i, row := range matrix {
		for j, char := range row {
			if char != "." {
				charMap[char] = append(charMap[char], Position{x: i, y: j})
			}
		}
	}

	return charMap
}

func addPosition(pos Position, matrix [][]string, set map[string]bool) {
	if pos.x >= 0 && pos.x < len(matrix) && pos.y >= 0 && pos.y < len(matrix[0]) {
		key := fmt.Sprintf("%d,%d", pos.x, pos.y)
		set[key] = true
	}
}

func findMarkedPositions(charMap map[string][]Position, matrix [][]string, considerAll bool) int {
	visited := make(map[string]bool)

	for _, positions := range charMap {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				dx := positions[j].x - positions[i].x
				dy := positions[j].y - positions[i].y
				// consider all possible positions for part 2
				if considerAll {
					k := 0
					for {
						newPos := Position{x: positions[i].x - dx*k, y: positions[i].y - dy*k}
						if newPos.x < 0 || newPos.x >= len(matrix) || newPos.y < 0 || newPos.y >= len(matrix[0]) {
							break
						}
						addPosition(newPos, matrix, visited)
						k++
					}
				} else {
					addPosition(Position{x: positions[i].x - dx, y: positions[i].y - dy}, matrix, visited)
				}

				if considerAll {
					k := 0
					for {
						newPos := Position{x: positions[j].x + dx*k, y: positions[j].y + dy*k}
						if newPos.x < 0 || newPos.x >= len(matrix) || newPos.y < 0 || newPos.y >= len(matrix[0]) {
							break
						}
						addPosition(newPos, matrix, visited)
						k++
					}
				} else {
					addPosition(Position{x: positions[j].x + dx, y: positions[j].y + dy}, matrix, visited)
				}
			}
		}
	}

	return len(visited)
}

func main() {
	matrix := parseInput("input.txt")
	charMap := groupCharacters(matrix)
	part1 := findMarkedPositions(charMap, matrix, false)
	fmt.Println("Answer for Part 1:", part1)
	part2 := findMarkedPositions(charMap, matrix, true)
	fmt.Println("Answer for Part 2:", part2)
}
