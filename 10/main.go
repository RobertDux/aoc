package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func walk(grid map[Coord]int, start Coord, curr int, visited map[Coord]bool) int {

	if (grid[Coord{start.x, start.y}] == 9 && curr == 9) {
		visited[Coord{start.x, start.y}] = true
		return 1
	}

	if (grid[Coord{start.x, start.y}] != curr) {
		return 0
	}

	return walk(grid, Coord{start.x + 1, start.y}, curr+1, visited) +
		walk(grid, Coord{start.x, start.y + 1}, curr+1, visited) +
		walk(grid, Coord{start.x - 1, start.y}, curr+1, visited) +
		walk(grid, Coord{start.x, start.y - 1}, curr+1, visited)

}

func main() {
	contents, _ := os.ReadFile("data/day10.txt")
	lines := strings.Split(string(contents), "\n")
	var partA, partB int
	grid := make(map[Coord]int)

	for row, line := range lines {
		for col, char := range strings.Split(line, "") {
			num, _ := strconv.Atoi(char)
			grid[Coord{col, row}] = num
		}
	}

	for coord, height := range grid {
		visited := make(map[Coord]bool)

		if height == 0 {
			partB += walk(grid, coord, 0, visited)
			partA += len(visited)
		}
	}

	fmt.Println("Part A:", partA)
	fmt.Println("Part B:", partB)
}
