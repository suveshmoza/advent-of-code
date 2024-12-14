package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	W    = 101
	H    = 103
	time = 100
)

type Robot struct {
	px, py int
	vx, vy int
}

func calculateSafetyFactor(robots []Robot) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	midRow := H / 2
	midCol := W / 2

	for _, robot := range robots {
		newX := (robot.px + robot.vx*time) % W
		if newX < 0 {
			newX += W
		}
		newY := (robot.py + robot.vy*time) % H
		if newY < 0 {
			newY += H
		}

		if newX > midCol && newY < midRow {
			q1++
		} else if newX < midCol && newY < midRow {
			q2++
		} else if newX < midCol && newY > midRow {
			q3++
		} else if newX > midCol && newY > midRow {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func main() {
	robots := []Robot{}

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		p := strings.TrimPrefix(parts[0], "p=")
		v := strings.TrimPrefix(parts[1], "v=")

		pos := strings.Split(p, ",")
		px, _ := strconv.Atoi(pos[0])
		py, _ := strconv.Atoi(pos[1])

		vel := strings.Split(v, ",")
		vx, _ := strconv.Atoi(vel[0])
		vy, _ := strconv.Atoi(vel[1])

		robots = append(robots, Robot{px: px, py: py, vx: vx, vy: vy})
	}

	fmt.Println("Part 1: ", calculateSafetyFactor(robots))

	seconds := 0

	for {
		seconds++
		grid := make([][]int, H)
		for i := range grid {
			grid[i] = make([]int, W)
		}

		bad := false
		for _, robot := range robots {
			nx := (robot.px + seconds*robot.vx) % W
			ny := (robot.py + seconds*robot.vy) % H

			if nx < 0 {
				nx += W
			}
			if ny < 0 {
				ny += H
			}

			grid[ny][nx]++
			if grid[ny][nx] > 1 {
				bad = true
			}
		}

		if !bad {
			fmt.Println("Part 2:", seconds)
			break
		}
	}
}
