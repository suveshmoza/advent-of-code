package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	c, r int
}

var (
	end = Point{70, 70}
	ds  = []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

func parseInput(filename string) ([]Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		c, err1 := strconv.Atoi(parts[0])
		r, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}
		data = append(data, Point{c, r})
	}
	return data, nil
}

func solve(bounds map[Point]struct{}) int {
	queue := []struct {
		p Point
		t int
	}{{Point{0, 0}, 0}}

	seen := make(map[Point]struct{})
	seen[Point{0, 0}] = struct{}{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.p == end {
			return current.t
		}

		for _, d := range ds {
			np := Point{current.p.c + d.c, current.p.r + d.r}

			if _, exists := seen[np]; exists {
				continue
			}

			if _, blocked := bounds[np]; blocked || np.c < 0 || np.c >= 71 || np.r < 0 || np.r >= 71 {
				continue
			}

			queue = append(queue, struct {
				p Point
				t int
			}{np, current.t + 1})
			seen[np] = struct{}{}
		}
	}

	return -1
}

func main() {
	data, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	bounds := make(map[Point]struct{})
	for i := 0; i < 1024 && i < len(data); i++ {
		bounds[data[i]] = struct{}{}
	}

	part1 := solve(bounds)

	if part1 == -1 {
		fmt.Println("No solution with initial bounds")
		return
	}

	fmt.Println("Part 1: ", part1)

	for i := 1024; i < len(data); i++ {
		bounds[data[i]] = struct{}{}
		if solve(bounds) == -1 {
			fmt.Println("Part 2:", data[i].c, ",", data[i].r)
			break
		}
	}
}
