package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isPossibleFurther(design string, pattern []string, possible map[string]int) int {
	if val, ok := possible[design]; ok {
		return val
	}

	isPossibleCount := 0

	for _, t := range pattern {
		if len(t) > len(design) {
			continue
		}

		if strings.HasPrefix(design, t) {
			if len(t) == len(design) {
				isPossibleCount++
				continue
			}
			isPossibleCount += isPossibleFurther(design[len(t):], pattern, possible)
		}
	}
	possible[design] = isPossibleCount
	return isPossibleCount
}

func getPossibleCount(patterns []string, designs []string) (int, int) {
	count1, count2 := 0, 0
	possible := make(map[string]int)
	for _, d := range designs {
		isPossibleCount := isPossibleFurther(d, patterns, possible)
		count2 += isPossibleCount
		if isPossibleCount > 0 {
			count1 += 1
		}
	}
	return count1, count2
}

func main() {
	patterns := []string{}
	designs := []string{}

	data, err := os.Open("input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)

	if scanner.Scan() {
		patterns = append(patterns, strings.Split(scanner.Text(), ", ")...)
	}

	scanner.Scan()

	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}

	result1, result2 := getPossibleCount(patterns, designs)
	fmt.Println("Part 1", result1)
	fmt.Println("Part 2", result2)

}
