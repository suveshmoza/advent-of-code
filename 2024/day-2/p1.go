package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	var mat [][]int

	data, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	// count:=0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, " ")
		var list1 []int
		for _, val := range temp {
			val, e := strconv.Atoi(val)
			if e != nil {
				panic(e)
			}
			list1 = append(list1, val)
		}
		mat = append(mat, list1)
	}
	count := 0
	for _, row := range mat {
		isIncreasing := row[0] < row[1]
		isValid := true

		for i := 0; i < len(row)-1; i++ {
			diff := (row[i+1] - row[i])
			if isIncreasing {
				if diff <= 0 || diff > 3 {
					isValid = false
					break
				}
			} else {
				if diff >= 0 || -diff > 3 {
					isValid = false
					break
				}
			}
		}
		if isValid {
			count++
		}
	}

	print(count)

}
