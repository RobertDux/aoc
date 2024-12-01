package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read(filepath string) *os.File {
	f, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error while reading file.")
		os.Exit(1)
	}

	return f
}

func close(f *os.File) {
	err := f.Close()

	if err != nil {
		fmt.Println("Error while closing file.")
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Incorrect usage.")
		os.Exit(1)
	}

	day, err := strconv.Atoi(args[0])
	dayString := fmt.Sprintf("%02d", day)
	filepath := "./data/day" + dayString + ".txt"

	if err != nil || day < 0 {
		fmt.Println("Invalid day.")
	}

	f := read(filepath)
	defer close(f)

	// Day 01:
	var leftList []int
	var rightList []int
	var difference int
	var similarityScore int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		leftPart, _ := strconv.Atoi(parts[0])
		rightPart, _ := strconv.Atoi(parts[1])
		leftList = append(leftList, leftPart)
		rightList = append(rightList, rightPart)
	}

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	rightListCount := make(map[int]int)

	for _, num := range rightList {
		rightListCount[num]++
	}

	for i, elem := range leftList {
		difference += Abs(elem - rightList[i])
		similarityScore += elem * rightListCount[elem]
	}

	fmt.Println("Part 1:", difference)
	fmt.Println("Part 2:", similarityScore)
}
