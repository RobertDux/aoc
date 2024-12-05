package main

import (
	"fmt"
	"os"
	"slices"
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

func parse(lines []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var updates [][]int

	var a, b int
	isSplit := false

	for _, line := range lines {
		if line == "" {
			isSplit = true
			continue
		}

		if !isSplit {
			fmt.Sscanf(line, "%d|%d", &a, &b)
			rules[b] = append(rules[b], a)
			continue
		}

		nums := getNums(strings.Split(line, ","))
		updates = append(updates, nums)

	}

	return rules, updates
}

func loop(update []int, rules map[int][]int) int {
	for i, page := range update {
		for _, before := range rules[page] {
			if slices.Contains(update, before) && !slices.Contains(update[:i], before) {
				return 0
			}
		}
	}

	return update[(len(update)-1)/2]
}

func solve(update []int, rules map[int][]int) int {
	for {
		if loop(update, rules) != 0 {
			break
		}

		for i, page := range update {
			for _, before := range rules[page] {
				if slices.Contains(update, before) && !slices.Contains(update[:i], before) {
					idx := slices.Index(update, before)
					update[idx], update[i] = update[i], update[idx]
				}
			}
		}
	}

	return update[(len(update)-1)/2]
}

func main() {
	var part_a, part_b int
	var incorrect [][]int

	contents, _ := os.ReadFile("data/day05.txt")
	lines := strings.Split(string(contents), "\n")
	rules, updates := parse(lines)

	for _, update := range updates {
		score := loop(update, rules)

		if score == 0 {
			incorrect = append(incorrect, update)
		} else {
			part_a += score
		}
	}

	for _, item := range incorrect {
		part_b += solve(item, rules)
	}

	fmt.Println("Part 1:", part_a)
	fmt.Println("Part 2:", part_b)
}
