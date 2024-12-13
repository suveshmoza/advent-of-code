package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Button struct {
	x int
	y int
}

type Prize struct {
	x int
	y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const offset = 10000000000000

func parseButton(line string) Button {
	parts := strings.Split(line, ", ")
	xStr := strings.Split(parts[0], "+")[1]
	yStr := strings.Split(parts[1], "+")[1]

	x, err := strconv.Atoi(xStr)
	check(err)
	y, err := strconv.Atoi(yStr)
	check(err)

	return Button{x: x, y: y}
}

func parsePrize(line string) Prize {
	parts := strings.Split(line, ", ")
	xStr := strings.Split(parts[0], "=")[1]
	yStr := strings.Split(parts[1], "=")[1]

	x, err := strconv.Atoi(xStr)
	check(err)
	y, err := strconv.Atoi(yStr)
	check(err)

	return Prize{x: x, y: y}
}

// Recursive function with memoization
func calcMinTokens(i, j int, buttonA, buttonB Button, prize Prize, memo map[string]int) int {
	if i == prize.x && j == prize.y {
		return 0
	}
	if i > prize.x || j > prize.y {
		return 1e9
	}

	key := fmt.Sprintf("%d,%d", i, j)
	if val, exists := memo[key]; exists {
		return val
	}

	costA := 3 + calcMinTokens(i+buttonA.x, j+buttonA.y, buttonA, buttonB, prize, memo)
	costB := 1 + calcMinTokens(i+buttonB.x, j+buttonB.y, buttonA, buttonB, prize, memo)

	memo[key] = min(costA, costB)
	return memo[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calcMinTokens2(p1, p2 Button, p3 Prize) int {
	p3.x += offset
	p3.y += offset
	if (((p2.x*(-p3.y))-(p2.y*(-p3.x)))%((p1.x*p2.y)-(p1.y*p2.x)) == 0) &&
		((((-p3.x)*p1.y)-((-p3.y)*p1.x))%((p1.x*p2.y)-(p1.y*p2.x)) == 0) {
		x := ((p2.x * (-p3.y)) - (p2.y * (-p3.x))) / ((p1.x * p2.y) - (p1.y * p2.x))
		y := (((-p3.x) * p1.y) - ((-p3.y) * p1.x)) / ((p1.x * p2.y) - (p1.y * p2.x))
		return 3*x + y
	}
	return 0
}

func main() {
	data, err := os.Open("./input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)

	var buttonA, buttonB Button
	var prize Prize

	totalTokens := 0
	minTokens2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Button A:") {
			buttonA = parseButton(line)
		} else if strings.HasPrefix(line, "Button B:") {
			buttonB = parseButton(line)
		} else if strings.HasPrefix(line, "Prize:") {
			prize = parsePrize(line)

			// Solve for this claw machine
			memo := make(map[string]int)
			minTokens := calcMinTokens(0, 0, buttonA, buttonB, prize, memo)
			minTokens2 += calcMinTokens2(buttonA, buttonB, prize)

			// Check if it's possible to win the prize
			if minTokens < 1e9 {
				totalTokens += minTokens
			}
		}
	}

	check(scanner.Err())

	fmt.Printf("Total Minimum Tokens: %d\n", totalTokens)
	fmt.Printf("Total Minimum Tokens: %d\n", minTokens2)
}
