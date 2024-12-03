package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func loop(lines []string, part2 bool) int {
	re := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)|do\\(\\)|don't\\(\\)")
	enabled := true
	count := 0

	for _, line := range lines {
		groups := re.FindAllString(line, -1)

		for _, group := range groups {
			if part2 && group == "do()" {
				enabled = true
				continue
			}

			if part2 && group == "don't()" {
				enabled = false
				continue
			}

			if enabled {
				var a, b int
				fmt.Sscanf(group, "mul(%d,%d)", &a, &b)
				count += a * b
			}
		}
	}

	return count
}

func main() {
	contents, _ := os.ReadFile("data/day03.txt")
	lines := strings.Split(string(contents), "\n")
	fmt.Println("Part 1:", loop(lines, false))
	fmt.Println("Part 2:", loop(lines, true))
}
