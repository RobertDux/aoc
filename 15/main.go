package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	y int
	x int
}

func parse(data []string) (map[Coord]string, Coord) {
	grid := make(map[Coord]string)
	var start Coord

	for y, row := range strings.Split(data[0], "\n") {
		for x, char := range row {
			char := string(char)

			if char == "@" {
				start = Coord{y, x}
				char = "."
			}

			grid[Coord{y, x}] = char
		}
	}

	return grid, start
}

func parseLarge(data []string) (map[Coord]string, Coord) {
	grid := make(map[Coord]string)
	var start Coord

	for y, row := range strings.Split(data[0], "\n") {
		for x, char := range row {
			char := string(char)

			if char == "@" {
				start = Coord{y, x}
				char = ".."
			}

			if char == "#" {
				char = "##"
			}

			if char == "." {
				char = ".."
			}

			if char == "O" {
				char = "[]"
			}

			grid[Coord{y, x}] = char
		}
	}

	return grid, start
}

func nextCoord(start Coord, dir string) Coord {
	if dir == "^" {
		return Coord{start.y - 1, start.x}
	}

	if dir == ">" {
		return Coord{start.y, start.x + 1}
	}

	if dir == "v" {
		return Coord{start.y + 1, start.x}
	}

	return Coord{start.y, start.x - 1}
}

func push(grid map[Coord]string, coord Coord, dir string) bool {
	if grid[nextCoord(coord, dir)] == "." {
		return false
	}

	count := 0
	search := coord

	for {
		if grid[nextCoord(search, dir)] == "#" {
			return true
		}

		if grid[nextCoord(search, dir)] == "." {
			break
		}

		count++
		search = nextCoord(search, dir)
	}

	q := nextCoord(coord, dir)
	grid[q] = "."

	for i := 0; i < count; i++ {
		q = nextCoord(q, dir)
		grid[q] = "O"
	}

	return false
}

func gps(grid map[Coord]string) int {
	sum := 0

	for pos, char := range grid {
		if char == "O" {
			sum += pos.y*100 + pos.x
		}
	}

	return sum
}

func display(grid map[Coord]string, start Coord) {
	for y := 0; y < 7; y++ {
		for x := 0; x < 15; x++ {
			if y == start.y && x == start.x {
				fmt.Print("@")
			} else {
				fmt.Print(grid[Coord{y, x}])
			}
		}
		fmt.Println()
	}
}

func walk(grid map[Coord]string, start Coord, path string) Coord {
	for _, char := range path {
		dir := string(char)

		if dir == "\n" {
			continue
		}

		if grid[nextCoord(start, dir)] == "#" {
			continue
		}

		if push(grid, start, dir) {
			continue
		}

		start = nextCoord(start, dir)
	}

	return start
}

func main() {
	contents, _ := os.ReadFile("data/day15.txt")
	data := strings.Split(string(contents), "\n\n")
	grid, start := parse(data)
	start = walk(grid, start, data[1])
	display(grid, start)
	fmt.Println("Part A:", gps(grid))

	gridB, startB := parseLarge(data)
	display(gridB, startB)
}
