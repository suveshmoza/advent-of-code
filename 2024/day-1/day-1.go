package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, []int, map[int]int, error) {
	var list1 []int
	var list2 []int
	list2Map := make(map[int]int)

	data, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("file doesn't exist: %v", err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, "   ")
		val1, err := strconv.Atoi(temp[0])
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error parsing value from first column: %v", err)
		}
		list1 = append(list1, val1)

		val2, err := strconv.Atoi(temp[1])
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error parsing value from second column: %v", err)
		}
		list2 = append(list2, val2)
		list2Map[val2]++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return list1, list2, list2Map, nil
}

func calculateAbsDifference(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for i := range list1 {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return sum
}

func calculateWeightedSum(list1 []int, list2Map map[int]int) int {
	weightedSum := 0
	for _, val := range list1 {
		weightedSum += val * list2Map[val]
	}
	return weightedSum
}

func main() {
	list1, list2, list2Map, err := readFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Task 1
	sum := calculateAbsDifference(list1, list2)
	fmt.Printf("Sum of absolute differences: %d\n", sum)

	// Task 2
	weightedSum := calculateWeightedSum(list1, list2Map)
	fmt.Printf("Weighted sum: %d\n", weightedSum)
}
