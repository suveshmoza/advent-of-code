package main

import (
	"bufio"
	"fmt"
	"os"
)

func countPossible(mat [][]string, i int, j int) int {
	count := 0
	target1, target2 := "XMAS", "SAMX"

	// Check forward
	curr := ""
	y := j
	for y < len(mat[0]) && len(curr) < 4 {
		curr += mat[i][y]
		y++
	}
	if curr == target1 || curr == target2 {
		count++
	}

	// Check backward
	curr = ""
	y = j
	for y >= 0 && len(curr) < 4 {
		curr += mat[i][y]
		y--
	}
	if curr == target1 || curr == target2 {
		count++
	}

	// Check up
	curr = ""
	x := i
	for x >= 0 && len(curr) < 4 {
		curr += mat[x][j]
		x--
	}
	if curr == target1 || curr == target2 {
		count++
	}

	// Check down
	curr = ""
	x = i
	for x < len(mat) && len(curr) < 4 {
		curr += mat[x][j]
		x++
	}
	if curr == target1 || curr == target2 {
		count++
	}

	// Check upward diagonals
	curr = ""
	x, y = i, j
	for x >= 0 && y >= 0 && len(curr) < 4 {
		curr += mat[x][y]
		x--
		y--
	}
	if curr == target1 || curr == target2 {
		count++
	}

	// Check downward diagonals
	curr = ""
	x, y = i, j
	for x < len(mat) && y < len(mat[0]) && len(curr) < 4 {
		curr += mat[x][y]
		x++
		y++
	}
	if curr == target1 || curr == target2 {
		count++
	}

	// Check upward-right diagonals
	curr = ""
	x, y = i, j
	for x >= 0 && y < len(mat[0]) && len(curr) < 4 {
		curr += mat[x][y]
		x--
		y++
	}
	if curr == target1 || curr == target2 {
		count++
	}

	// Check downward-left diagonals
	curr = ""
	x, y = i, j
	for x < len(mat) && y >= 0 && len(curr) < 4 {
		curr += mat[x][y]
		x++
		y--
	}
	if curr == target1 || curr == target2 {
		count++
	}

	return count
}

func isXMas(mat [][]string, i int, j int) bool {
	if i-1 < 0 || j-1 < 0 || i+1 >= len(mat) || j+1 >= len(mat[0]) {
		return false
	}

	leftDiag := mat[i-1][j-1] + mat[i][j] + mat[i+1][j+1]
	rightDiag := mat[i+1][j-1] + mat[i][j] + mat[i-1][j+1]

	if (leftDiag == "MAS" || leftDiag == "SAM") && (rightDiag == "MAS" || rightDiag == "SAM") {
		return true
	}
	return false
}

func main() {
	data, e := os.Open("./input.txt")
	if e != nil {
		panic(e)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	var mat [][]string

	for scanner.Scan() {
		line := scanner.Text()
		var currLine []string
		for _, val := range line {
			currLine = append(currLine, string(val))
		}
		mat = append(mat, currLine)
	}

	count := 0

	for i, row := range mat {
		for j := range row {
			if mat[i][j] == "X" {
				count += countPossible(mat, i, j)
			}
		}
	}

	fmt.Println("Answer for p1 :", count)
	count = 0
	for i, row := range mat {
		for j := range row {
			if mat[i][j] == "A" && isXMas(mat, i, j) {
				count++
			}
		}
	}

	fmt.Println("Answer for p2 :", count)

}
