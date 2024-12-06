package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Queue Implementation
type Queue[T any] struct {
    Items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.Items = append(q.Items, item)
}

func (q *Queue[T]) Dequeue() T {
    item := q.Items[0]
    q.Items = q.Items[1:]
    return item
}

func (q *Queue[T]) IsEmpty() bool {
    return len(q.Items) == 0
}

func deepCopy(mat [][]string) [][]string {
	copyMat := make([][]string, len(mat))
	for i := range mat {
		copyMat[i] = append([]string(nil), mat[i]...)
	}
	return copyMat
}


func printMat(mat [][]string){
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[i]);j++{
			fmt.Print(mat[i][j])
		}
		fmt.Println()
	}
}

func check(e error){
	if(e!=nil){
		panic(e)
	}
}


func traverse(mat [][]string,x int,y int) [][]string {
	q:=Queue[[]int]{}
	dirIndex:=0
	dirs:=[][]int {{-1,0},{0,1},{1,0},{0,-1}}
	q.Enqueue([]int{x,y})
	for !q.IsEmpty(){
		item:=q.Dequeue()
		currX,currY:=item[0],item[1]
		
		newX,newY:=currX+dirs[dirIndex][0],currY+dirs[dirIndex][1]
		if newX < 0 || newX >= len(mat) || newY < 0 || newY >= len(mat[0]) {
			return mat
		}

		if mat[newX][newY]=="#"{
			dirIndex=(dirIndex+1)%4
			newX,newY=currX+dirs[dirIndex][0],currY+dirs[dirIndex][1]
		}
		q.Enqueue([]int{newX,newY})
		mat[currX][currY]="X"
	}

	return mat
}

func traverse2(mat [][]string, x, y int) bool {
    q := Queue[[]int]{}
    dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
    visited := make(map[string]bool)

    dirIndex := 0
    q.Enqueue([]int{x, y, dirIndex})

    for !q.IsEmpty() {
        current := q.Dequeue()
        currX, currY, currDir := current[0], current[1], current[2]

        stateKey := fmt.Sprintf("%d,%d,%d", currX, currY, currDir)
        if visited[stateKey] {
            return true 
        }
        visited[stateKey] = true

        newX := currX + dirs[currDir][0]
        newY := currY + dirs[currDir][1]

        if newX < 0 || newX >= len(mat) || newY < 0 || newY >= len(mat[0]) {
            return false 
        }

        if mat[newX][newY] == "#" {
            newDir := (currDir + 1) % 4
            q.Enqueue([]int{currX, currY, newDir})
        } else {
            q.Enqueue([]int{newX, newY, currDir})
        }
    }

    return false 
}

func countPossibleObstructions(mat [][]string, startX, startY int) int {
	totalLoops := 0
	
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == "#" || mat[i][j]=="^" {
				continue
			}
			testMat := deepCopy(mat)
			testMat[i][j] = "#"
			if traverse2(testMat,startX,startY){
				totalLoops++
			}
			testMat[i][j]="."
		}
	}
	return totalLoops
}

func main(){
	data,e:=os.Open("./input.txt")
	check(e)
	defer data.Close()

	scanner:=bufio.NewScanner(data)

	var mat [][]string

	for scanner.Scan(){
		line:=scanner.Text()
		mat = append(mat, strings.Split(line,""))
	}

	var x,y int

	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[i]);j++{
			if mat[i][j]=="^"{
				x=i
				y=j
				break
			}
		}
	}

	ogMat := deepCopy(mat)
	mat = traverse(mat, x, y)	
	count:=0
	
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[i]);j++{
			if mat[i][j]=="X"{
				count++
			}
		}
	}
	fmt.Println("Answer for Part 1:",count+1)
	fmt.Println("Answer for Part 2:",countPossibleObstructions(ogMat, x, y))
}