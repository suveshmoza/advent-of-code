package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const LOOPTIMES = 25

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var memo = make(map[string]int)

func blinking(stone, n int) int {
	key := fmt.Sprintf("%d,%d", stone, n)
	if val, exists := memo[key]; exists {
		return val
	}

	if n == LOOPTIMES {
		return 1
	}
	if stone == 0 {
		memo[key] = blinking(1, n+1)
		return memo[key]
	}

	digitLength := len(strconv.Itoa(stone))
	if digitLength%2 == 0 {
		convertedStone := strconv.Itoa(stone)
		leftHalf, _ := strconv.Atoi(convertedStone[:digitLength/2])
		rightHalf, _ := strconv.Atoi(convertedStone[digitLength/2:])
		memo[key] = blinking(leftHalf, n+1) + blinking(rightHalf, n+1)
	} else {
		memo[key] = blinking(stone*2024, n+1)
	}

	return memo[key]
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("Input file is empty.")
		return
	}
	line := scanner.Text()
	stones := strings.Split(line, " ")
	stoneArray := []int{}

	for _, stone := range stones {
		val, err := strconv.Atoi(stone)
		check(err)
		stoneArray = append(stoneArray, val)
	}

	sum := 0
	for _, stone := range stoneArray {
		sum += blinking(stone, 0)
	}
	fmt.Println(sum)
}
