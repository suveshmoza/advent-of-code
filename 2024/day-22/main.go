package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const PRUNE = 16777216

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var finalMp = make(map[[4]int]int)

func mixValue(value, secretNumber int) int {
	return value ^ secretNumber
}

func multValue(secretNumber int) int {
	return secretNumber * 2048
}

func findMaxBananas(sequence []int) {
	diffsequence := []int{}
	for i, diff := range sequence {
		if i == 0 {
			diffsequence = append(diffsequence, 0)
		} else {
			diffsequence = append(diffsequence, (diff%10)-(sequence[i-1]%10))
		}
	}

	mp := make(map[[4]int]int)

	for i := 0; i < len(sequence)-4; i++ {
		key := [4]int{
			diffsequence[i], diffsequence[i+1], diffsequence[i+2], diffsequence[i+3],
		}

		if _, ok := mp[key]; !ok {
			mp[key] = sequence[i+3] % 10
		}
	}

	for seq, val := range mp {
		if _, ok := finalMp[seq]; !ok {
			finalMp[seq] = 0
		}
		finalMp[seq] += val
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateNewSecretNumber(secretNumber int) int {
	currNumberSequences := []int{}
	currNumberSequences = append(currNumberSequences, secretNumber)
	for i := 1; i <= 2000; i++ {
		secretNumber = ((secretNumber * 64) ^ secretNumber) % PRUNE
		secretNumber = ((secretNumber / 32) ^ secretNumber) % PRUNE
		secretNumber = ((secretNumber * 2048) ^ secretNumber) % PRUNE

		currNumberSequences = append(currNumberSequences, secretNumber)
	}

	findMaxBananas(currNumberSequences)
	return secretNumber
}

func main() {
	data, err := os.Open("input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)
	initialSecretNumbers := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		secretNumber, err := strconv.Atoi(line)
		check(err)
		initialSecretNumbers = append(initialSecretNumbers, secretNumber)
	}

	part1 := 0
	for _, secretNumber := range initialSecretNumbers {
		val1 := generateNewSecretNumber(secretNumber)
		part1 += val1
	}

	part2 := -1
	for _, val := range finalMp {
		if val > part2 {
			part2 = val
		}
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)

}
