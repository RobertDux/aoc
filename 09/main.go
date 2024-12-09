package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func space(size int, i int) []int {
	var val int
	arr := make([]int, size)

	if i%2 == 0 {
		val = i / 2
	} else {
		val = -1
	}

	for i := range arr {
		arr[i] = val
	}

	return arr

}

func getLastIndex(arr []int) int {
	idx := len(arr) - 1

	for idx > 0 {
		if arr[idx] != -1 {
			return idx
		}

		idx--
	}

	return -1
}

func solve(data string, part func(diskmap []int) []int) int {
	var diskMap []int
	var count int

	for i, char := range strings.Split(data, "") {
		size, _ := strconv.Atoi(char)
		arr := space(size, i)
		diskMap = append(diskMap, arr...)
	}

	diskMap = part(diskMap)

	for i, v := range diskMap {
		if v != -1 {
			count += i * v
		}
	}

	return count
}

func findSpace(diskMap []int, size int, end int) int {
	for i, disk := range diskMap {
		if disk == -1 && i < end {
			var fits = true

			for j := 0; j < size; j++ {
				if diskMap[i+j] != -1 {
					fits = false
				}
			}

			if fits {
				return i
			}
		}
	}

	return -1
}

func partA(diskMap []int) []int {
	for i, v := range diskMap {
		if v == -1 {
			idx := getLastIndex(diskMap)

			if idx < i {
				break
			}

			diskMap[i], diskMap[idx] = diskMap[idx], diskMap[i]
		}
	}

	return diskMap
}

func partB(diskMap []int) []int {
	var lenFile = 1
	var prev = diskMap[len(diskMap)-1]

	for i := len(diskMap) - 2; i >= 0; i-- {
		// more than 1 empty space: ignore
		if diskMap[i] == -1 && diskMap[i] == prev {
			prev = diskMap[i]
			continue
		}

		// current file is over
		if diskMap[i] != prev {
			// if not empty space
			if lenFile > 0 {
				start := findSpace(diskMap, lenFile, i)
				// and room to fill
				if start != -1 {
					// fill
					for j := 0; j < lenFile; j++ {
						diskMap[start+j], diskMap[i+j+1] = diskMap[i+j+1], diskMap[start+j]
					}
				}
			}

			// reset file length
			if diskMap[i] == -1 {
				lenFile = 0
			} else {
				lenFile = 1
			}
		} else {
			// count current file length
			lenFile++
		}

		prev = diskMap[i]
	}

	return diskMap
}

func main() {
	contents, _ := os.ReadFile("data/day09.txt")
	data := string(contents)
	fmt.Println("Part A:", solve(data, partA))
	fmt.Println("Part B:", solve(data, partB))
}
