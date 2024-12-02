package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	// Check if increasing
	isValid := true
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff <= 0 || diff > 3 {
			isValid = false
			break
		}
	}
	if isValid {
		return true
	}

	// Check if decreasing
	isValid = true
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff >= 0 || diff < -3 {
			return false
		}
	}
	return true
}

func isPossible(currRow []int) bool {
	if isSafe(currRow) {
		return true
	}

	for i := 0; i < len(currRow); i++ {
		var temp []int
		temp = append(temp, currRow[:i]...)
		temp = append(temp, currRow[i+1:]...)
		if isSafe(temp) {
			return true
		}
	}
	return false
}

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
		if isPossible(row) {
			count++
		}
	}

	fmt.Println(count)

}
