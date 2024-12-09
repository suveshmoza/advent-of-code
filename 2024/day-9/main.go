package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func helper(blockSize, freeSpace, temp int) string {
	line := ""
	for i := 0; i < blockSize; i++ {
		line += strconv.Itoa(temp)
	}
	for i := 0; i < freeSpace; i++ {
		line += "."
	}
	return line
}

func generateDiskMap(line string) ([]string, int) {
	res := []string{}
	counter := 0
	for i := 0; i < len(line); i++ {
		if i%2 == 0 {
			times, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			for j := 0; j < times; j++ {
				res = append(res, strconv.Itoa(counter))
			}
			counter++
		} else {
			times, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			for j := 0; j < times; j++ {
				res = append(res, ".")
			}
		}
	}

	return res, counter - 1
}

func moveLastToEmpty(dMap []string) []string {
	res := []string{}
	i := 0
	j := len(dMap) - 1

	for i <= j {
		if dMap[i] != "." {
			res = append(res, string(dMap[i]))
			i++
		} else {

			for dMap[j] == "." && j > i {
				j--
			}
			if j > i {
				res = append(res, string(dMap[j]))
				j--
			}
			i++
		}
	}

	return res
}

func solve1(dMap []string) int {
	ans := 0
	for i := 0; i < len(dMap); i++ {
		val, err := strconv.Atoi(string(dMap[i]))
		if err != nil {
			panic(err)
		}
		ans += (i * val)
	}
	return ans
}

func findADotSegmentWithXLength(dMap []string, xLength int) int {
	for i := 0; i < len(dMap); i++ {
		if dMap[i] == "." {
			j := i
			for j < len(dMap) && dMap[j] == "." {
				j++
			}
			if j-i == xLength {
				return i
			}
		}
	}
	return -1
}

type Segment struct {
	Start int
	End   int
}

func moveFilesCompactly(dMap []string, fileId int) []string {
	type Segment struct {
		Start, End int
	}

	mp := make(map[int]Segment)

	for i := 0; i < len(dMap); i++ {
		if dMap[i] != "." {
			id, _ := strconv.Atoi(dMap[i])
			if existingSegment, ok := mp[id]; ok {
				mp[id] = Segment{Start: existingSegment.Start, End: i}
			} else {
				mp[id] = Segment{Start: i, End: i}
			}
		}
	}

	for fileId >= 0 {
		segment, exists := mp[fileId]
		if !exists {
			fileId--
			continue
		}

		segmentLengthRequired := segment.End - segment.Start + 1

		for i := 0; i < len(dMap); i++ {
			if dMap[i] == "." {
				start := i
				end := i

				if start >= segment.Start {
					break
				}
				for end < len(dMap) && dMap[end] == "." {
					end++
				}
				end--
				emptyLength := end - start + 1
				if segmentLengthRequired > 0 && emptyLength >= segmentLengthRequired {
					fillValues(dMap, start, segmentLengthRequired, strconv.Itoa(fileId))
					clearValues(dMap, mp[fileId].Start, mp[fileId].End)
					delete(mp, fileId)
					break
				}
			}
		}
		fileId--
	}

	return dMap
}

func fillValues(dMap []string, start, len int, value string) {
	for i := start; i < start+len; i++ {
		dMap[i] = value
	}
}

func clearValues(dMap []string, start, end int) {
	for i := start; i <= end; i++ {
		dMap[i] = "."
	}
}

func solve2(dMap []string) int {
	checksum := 0
	for i, ch := range dMap {
		if ch != "." {
			id, _ := strconv.Atoi(ch)
			checksum += i * id
		}
	}
	return checksum
}

func main() {
	data, err := os.Open("input.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Scan()
	line := scanner.Text()

	dMap, fileId := generateDiskMap(line)
	dMap1 := moveLastToEmpty(dMap)

	fmt.Println("Answer for part 1: ", solve1(dMap1))

	dMap2 := moveFilesCompactly(dMap, fileId)

	fmt.Println("Answer for part 2: ", solve2(dMap2))
}
