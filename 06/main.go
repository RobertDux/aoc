package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Location struct {
	x  int
	y  int
	dx int
	dy int
}

func getStart(grid map[Coord]string, val string) Coord {
	for k, v := range grid {
		if v == val {
			return k
		}
	}

	return Coord{-1, -1}
}

func partA(grid map[Coord]string) int {
	var partA int
	deltas := []Coord{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}
	pos := getStart(grid, "^")
	dir := 0

	for {
		dy := deltas[dir].y
		dx := deltas[dir].x

		if grid[pos] == "" {
			break
		}

		if grid[Coord{pos.x + dx, pos.y + dy}] == "#" {
			dir = (dir + 1) % 4
			continue
		}

		grid[pos] = "X"
		pos.x += dx
		pos.y += dy
	}

	for _, v := range grid {
		if v == "X" {
			partA++
		}
	}

	return partA
}

func isLoop(grid map[Coord]string, coord Coord) bool {
	seen := make(map[Location]bool)
	deltas := []Coord{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}
	pos := getStart(grid, "^")
	dir := 0

	for {
		dy := deltas[dir].y
		dx := deltas[dir].x
		step := Coord{pos.x + dx, pos.y + dy}

		if grid[pos] == "" {
			break
		}

		if grid[step] == "#" || step == coord {
			dir = (dir + 1) % 4
			continue
		}

		if (seen[Location{pos.x, pos.y, dx, dy}]) {
			return true
		}

		seen[Location{pos.x, pos.y, dx, dy}] = true
		pos.x += dx
		pos.y += dy
	}

	return false
}

func partB(grid map[Coord]string) int {
	var count int

	for coord, char := range grid {
		if char == "." && isLoop(grid, coord) {
			count++
		}
	}

	return count
}

func main() {
	grid := make(map[Coord]string)
	contents, _ := os.ReadFile("data/day06.txt")
	lines := strings.Split(string(contents), "\n")

	for row, line := range lines {
		for col, char := range strings.Split(line, "") {
			grid[Coord{col, row}] = char
		}
	}

	// fmt.Println("Part A:", partA(grid))
	fmt.Println("Part B:", partB(grid))
}
