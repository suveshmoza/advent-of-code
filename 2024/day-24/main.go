package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var wires = make(map[string]int)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Query struct {
	wireA   string
	wireB   string
	resWire string
	op      string
}

func calcOperationResult(x, y, op string) int {
	if op == "AND" {
		return wires[x] & wires[y]
	} else if op == "OR" {
		return wires[x] | wires[y]
	} else if op == "XOR" {
		return wires[x] ^ wires[y]
	}
	return 0
}

func main() {
	data, err := os.Open("input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		splittedString := strings.Split(line, " ")

		key := splittedString[0][0 : len(splittedString[0])-1]
		val, err := strconv.Atoi(splittedString[1])
		check(err)
		wires[key] = val
	}

	var processLater []Query

	for scanner.Scan() {
		line := scanner.Text()

		splittedString := strings.Split(line, " ")

		// x op y -> z
		x := splittedString[0]
		op := splittedString[1]
		y := splittedString[2]
		z := splittedString[4]

		_, xExists := wires[x]
		_, yExists := wires[y]

		if xExists && yExists {
			wires[z] = calcOperationResult(x, y, op)
		} else {
			processLater = append(processLater, Query{x, y, z, op})
		}
	}

	for len(processLater) > 0 {
		q := processLater[0]
		processLater = processLater[1:]

		_, xExists := wires[q.wireA]
		_, yExists := wires[q.wireB]

		if xExists && yExists {
			wires[q.resWire] = calcOperationResult(q.wireA, q.wireB, q.op)
		} else {
			processLater = append(processLater, q)
		}
	}

	keys := make([]string, 0, len(wires))
	for k := range wires {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var binaryString strings.Builder
	for i := len(keys) - 1; i >= 0; i-- {
		if keys[i][0] == 'z' {
			binaryString.WriteString(strconv.FormatInt(int64(wires[keys[i]]), 2))
		}
	}

	i, err := strconv.ParseInt(binaryString.String(), 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)

}
