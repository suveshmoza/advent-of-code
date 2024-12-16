package main

import (
	"bufio"
	"fmt"
	"os"
)

type State struct {
	x, y      int
	direction string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func findStartEnd(maze [][]rune) (start, end [2]int) {
	for y, row := range maze {
		for x, cell := range row {
			if cell == 'S' {
				start = [2]int{x, y}
			} else if cell == 'E' {
				end = [2]int{x, y}
			}
		}
	}
	return
}

func getNextStates(state State, maze [][]rune) []struct {
	state State
	cost  int
} {
	var moves []struct {
		state State
		cost  int
	}

	dirOffsets := map[string][2]int{
		"N": {0, -1},
		"E": {1, 0},
		"S": {0, 1},
		"W": {-1, 0},
	}
	dx, dy := dirOffsets[state.direction][0], dirOffsets[state.direction][1]
	newX, newY := state.x+dx, state.y+dy
	if newY >= 0 && newY < len(maze) && newX >= 0 && newX < len(maze[0]) && maze[newY][newX] != '#' {
		moves = append(moves, struct {
			state State
			cost  int
		}{State{newX, newY, state.direction}, 1})
	}

	rotations := map[string][]string{
		"N": {"E", "W"},
		"E": {"S", "N"},
		"S": {"W", "E"},
		"W": {"N", "S"},
	}
	for _, newDir := range rotations[state.direction] {
		moves = append(moves, struct {
			state State
			cost  int
		}{State{state.x, state.y, newDir}, 1000})
	}
	return moves
}

func findOptimalPaths(maze [][]rune) (optimalPositions map[[2]int]bool, minEndScore int) {
	start, end := findStartEnd(maze)

	initialState := State{start[0], start[1], "E"}
	minScores := make(map[State]int)
	minScores[initialState] = 0
	queue := []struct {
		state State
		score int
	}{{initialState, 0}}

	cameFrom := make(map[State][]struct {
		state State
		cost  int
	})
	endStates := make(map[State]bool)
	minEndScore = int(^uint(0) >> 1)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.score > minScores[current.state] {
			continue
		}

		if current.state.x == end[0] && current.state.y == end[1] {
			if current.score <= minEndScore {
				minEndScore = current.score
				endStates[current.state] = true
			}
			continue
		}

		for _, move := range getNextStates(current.state, maze) {
			newScore := current.score + move.cost
			if prevScore, exists := minScores[move.state]; !exists || newScore <= prevScore {
				if newScore < prevScore || !exists {
					minScores[move.state] = newScore
					cameFrom[move.state] = nil
					queue = append(queue, struct {
						state State
						score int
					}{move.state, newScore})
				}
				cameFrom[move.state] = append(cameFrom[move.state], struct {
					state State
					cost  int
				}{current.state, move.cost})
			}
		}
	}

	optimalPositions = make(map[[2]int]bool)
	var backtrack func(State, int)
	backtrack = func(state State, score int) {
		if score > minScores[state] {
			return
		}
		optimalPositions[[2]int{state.x, state.y}] = true
		if score == 0 {
			return
		}
		for _, prev := range cameFrom[state] {
			prevScore := score - prev.cost
			if minScores[prev.state] == prevScore {
				backtrack(prev.state, prevScore)
			}
		}
	}

	for endState := range endStates {
		if minScores[endState] == minEndScore {
			backtrack(endState, minEndScore)
		}
	}
	return
}

func main() {
	data, err := os.Open("input.txt")
	check(err)
	defer data.Close()

	var maze [][]rune
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		currLine := scanner.Text()
		maze = append(maze, []rune(currLine))
	}

	optimalPositions, minScore := findOptimalPaths(maze)

	fmt.Println("Part 1:", minScore)
	fmt.Println("Part 2:", len(optimalPositions))
}
