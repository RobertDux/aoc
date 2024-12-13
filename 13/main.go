package main

import (
	"fmt"
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

func parse(lines []string) [][]int {
	var nums [][]int
	var pos []int
	re := regexp.MustCompile(`\d+`)

	for i, line := range lines {
		t := (i + 1) % 4

		if t == 0 {
			nums = append(nums, pos)
			pos = nil
		} else {
			ints := getNums(re.FindAllString(line, -1))
			pos = append(pos, ints...)
		}
	}

	return nums
}

func claw(p []int, part2 bool) int {
	xa, ya, xb, yb, px, py := p[0], p[1], p[2], p[3], p[4], p[5]

	if part2 {
		px += 10000000000000
		py += 10000000000000
	}

	// xa * a + xb * b = px
	// ya * a + yb * b = py
	a := (py*xb - px*yb) / (xb*ya - xa*yb)
	b := (py*xa - px*ya) / (xa*yb - xb*ya)

	if !part2 && (a > 100 || b > 100) {
		return 0
	}

	if xa*a+xb*b == px && ya*a+yb*b == py {
		return 3*a + b
	}

	return 0
}

func main() {
	var tokens, moreTokens int
	contents, _ := os.ReadFile("data/day13.txt")
	lines := strings.Split(string(contents), "\n")
	positions := parse(lines)

	for _, pos := range positions {
		tokens += claw(pos, false)
		moreTokens += claw(pos, true)
	}

	fmt.Println("Part A", tokens)
	fmt.Println("Part B", moreTokens)
}
