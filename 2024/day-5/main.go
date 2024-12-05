package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isValidUpdate(dependencies map[int][]int, update []int) bool {
	for from, tos := range dependencies {
		posFrom := -1
		for i, page := range update {
			if page == from {
				posFrom = i
				break
			}
		}

		if posFrom == -1 {
			continue
		}

		for _, to := range tos {
			posTo := -1
			for i, page := range update {
				if page == to {
					posTo = i
					break
				}
			}

			if posTo == -1 {
				continue
			}

			if posFrom > posTo {
				return false
			}
		}
	}

	return true
}

func sortUsingDep(dependencies map[int][]int, update []int) []int {
	sort.Slice(update, func(i, j int) bool {
		for _, dependent := range dependencies[update[i]] {
			if dependent == update[j] {
				return false
			}
		}
		for _, dependent := range dependencies[update[j]] {
			if dependent == update[i] {
				return true
			}
		}
		return update[i] < update[j]
	})
	return update
}

func main() {
	data, e := os.Open("./input.txt")
	check(e)
	defer data.Close()

	dependencies := make(map[int][]int)
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("--Queries Started--")
			break
		}
		match := strings.Split(line, "|")
		from, err := strconv.Atoi(match[0])
		check(err)
		to, err := strconv.Atoi(match[1])
		check(err)
		dependencies[from] = append(dependencies[from], to)
	}

	var queries [][]int
	for scanner.Scan() {
		line := scanner.Text()
		pageStrs := strings.Split(line, ",")
		var pages []int
		for _, pageStr := range pageStrs {
			page, err := strconv.Atoi(pageStr)
			check(err)
			pages = append(pages, page)
		}
		queries = append(queries, pages)
	}

	// Part 1
	middleSum := 0
	for _, pages := range queries {
		if isValidUpdate(dependencies, pages) {
			middleSum += pages[len(pages)/2]
		}
	}
	fmt.Println("Answer for Problem 1: ", middleSum)

	// Part 2
	middleSum = 0
	for _, pages := range queries {
		if !isValidUpdate(dependencies, pages) {
			pages = sortUsingDep(dependencies, pages)
			middleSum += pages[len(pages)/2]
		}
	}
	fmt.Println("Answer for Problem 2: ", middleSum)
}
