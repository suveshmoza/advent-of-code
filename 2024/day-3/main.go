package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func calc(s string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindAllString(s, -1)

	if len(match) < 2 {
		return 0
	}

	m1, e := strconv.Atoi(match[0])
	if e != nil {
		panic(e)
	}

	m2, e := strconv.Atoi(match[1])
	if e != nil {
		panic(e)
	}

	return m1 * m2
}

func main() {
	data, e := os.Open("./input.txt")
	if e != nil {
		panic(e)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	// Part 1
	pattern1 := `mul\(\d+,\d+\)`
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(pattern1)
		match := re.FindAllString(line, -1)
		for _, val := range match {
			ans += calc(val)
		}
	}
	fmt.Println("Answer for p1:", ans)

	// Part 2
	ans = 0
	doOps := true
	data.Seek(0, 0)
	scanner = bufio.NewScanner(data)

	pattern2 := `mul\(\d+,\d+\)|don't\(\)|do\(\)`
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(pattern2)
		match := re.FindAllString(line, -1)
		for _, val := range match {
			if val == "don't()" {
				doOps = false
			} else if val == "do()" {
				doOps = true
			} else if doOps {
				ans += calc(val)
			}
		}
	}
	fmt.Println("Answer for p2:", ans)
}
