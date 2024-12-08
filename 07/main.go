package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getNums(line []string) []int {
	nums := make([]int, len(line))

	for i, line := range line {
		n, err := strconv.Atoi(line)

		if err != nil {
			return nil
		}

		nums[i] = n
	}

	return nums
}

func combine(a int, b int) int {
	var count int
	c := b

	for b > 0 {
		b /= 10
		count++
	}

	return int(math.Pow10(count))*a + c

}

func solveA(goal int, current int, parts []int) bool {
	if len(parts) == 0 {
		return current == goal
	}

	return solveA(goal, current+parts[0], parts[1:]) || solveA(goal, current*parts[0], parts[1:])

}

func solveB(goal int, current int, parts []int) bool {
	if len(parts) == 0 {
		return current == goal
	}

	return solveB(goal, current+parts[0], parts[1:]) ||
		solveB(goal, current*parts[0], parts[1:]) ||
		solveB(goal, combine(current, parts[0]), parts[1:])

}

func main() {
	var partA, partB int
	contents, _ := os.ReadFile("data/day07.txt")
	lines := strings.Split(string(contents), "\n")
	re := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		numbers := getNums(re.FindAllString(line, -1))
		if solveA(numbers[0], numbers[1], numbers[2:]) {
			partA += numbers[0]
		}

		if solveB(numbers[0], numbers[1], numbers[2:]) {
			partB += numbers[0]
		}
	}

	fmt.Println("Part A:", partA)
	fmt.Println("Part B:", partB)
}
