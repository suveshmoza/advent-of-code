package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(eq []int, target int, index int,currVal int, part2 bool) bool {
	if index == len(eq) {
		return currVal == target
	}
	
	// Basic operations that both parts use
	result := solve(eq, target, index+1, currVal+eq[index], part2) || 
			  solve(eq, target, index+1, currVal*eq[index], part2)
	
	// Add concatenation operation only for part 2
	if part2 {
		num, _ := strconv.Atoi(strconv.Itoa(currVal) + strconv.Itoa(eq[index]))
		result = result || solve(eq, target, index+1, num,part2)
	}
	
	return result
}



func main() {
	data, e := os.Open("input.txt")
	if e != nil {
		panic(e)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	eqs := make(map[int][]int)
	for scanner.Scan() {
		line:=scanner.Text()
		parts:=strings.Split(line,": ")

		row,e:=strconv.Atoi(parts[0])
		if e!=nil{
			panic(e)
		}
		nums:=make([]int,0)
		for _,n:=range strings.Split(parts[1]," "){
			num,e:=strconv.Atoi(n)
			if e!=nil{
				panic(e)
			}
			nums=append(nums,num)
		}
		eqs[row]=nums
	}

	ans:=0
	for key,value := range eqs {
		if solve(value, key, 1, value[0], false){
			ans += key
		}
	}
	fmt.Println("Part 1: ", ans)
	
	ans = 0
	for key, value := range eqs {
		if solve(value, key, 1, value[0], true){
			ans += key
		}
	}
	
	fmt.Println("Part 2: ", ans)
}
