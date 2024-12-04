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

func xmas(grid map[Coord]string, row int, col int) int {
	count := 0

	directions := [][]Coord{
		{{0, 1}, {0, 2}, {0, 3}},
		{{0, -1}, {0, -2}, {0, -3}},
		{{1, 0}, {2, 0}, {3, 0}},
		{{-1, 0}, {-2, 0}, {-3, 0}},
		{{1, 1}, {2, 2}, {3, 3}},
		{{1, -1}, {2, -2}, {3, -3}},
		{{-1, 1}, {-2, 2}, {-3, 3}},
		{{-1, -1}, {-2, -2}, {-3, -3}},
	}

	if grid[Coord{row, col}] != "X" {
		return 0
	}

	for _, dir := range directions {
		if grid[Coord{row + dir[0].x, col + dir[0].y}] == "M" && grid[Coord{row + dir[1].x, col + dir[1].y}] == "A" && grid[Coord{row + dir[2].x, col + dir[2].y}] == "S" {
			count++
		}
	}

	return count
}

func mas(grid map[Coord]string, row int, col int) int {
	count := 0

	directions := []map[Coord]string{
		{Coord{-1, -1}: "M", Coord{1, 1}: "S", Coord{-1, 1}: "M", Coord{1, -1}: "S"},
		{Coord{-1, -1}: "M", Coord{1, 1}: "S", Coord{-1, 1}: "S", Coord{1, -1}: "M"},
		{Coord{-1, -1}: "S", Coord{1, 1}: "M", Coord{-1, 1}: "M", Coord{1, -1}: "S"},
		{Coord{-1, -1}: "S", Coord{1, 1}: "M", Coord{-1, 1}: "S", Coord{1, -1}: "M"},
	}

	if grid[Coord{row, col}] != "A" {
		return 0
	}

	for _, dir := range directions {
		good := true

		for coord := range dir {
			if (grid[Coord{row + coord.x, col + coord.y}] != dir[coord]) {
				good = false
				break
			}
		}

		if good {
			count++
		}

	}

	return count
}

func main() {
	var count_xmas, count_mas int
	grid := make(map[Coord]string)

	contents, _ := os.ReadFile("data/day04.txt")
	lines := strings.Split(string(contents), "\n")

	for row, line := range lines {
		for col, char := range strings.Split(line, "") {
			grid[Coord{row, col}] = char
		}
	}

	width := len(lines[0])
	height := len(lines)

	for row := range height {
		for col := range width {
			count_xmas += xmas(grid, row, col)
			count_mas += mas(grid, row, col)
		}
	}

	fmt.Println("Part 1:", count_xmas)
	fmt.Println("Part 2:", count_mas)
}
