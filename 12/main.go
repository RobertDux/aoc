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
type Grid map[Coord]string
type Seen map[Coord]bool

var directions = []Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func getNeighbours(grid Grid, coord Coord) int {
	var count int

	for _, d := range directions {
		if grid[coord] == grid[Coord{coord.y + d.y, coord.x + d.x}] {
			count++
		}
	}

	return count
}

func getCorners(grid Grid, coord Coord) int {
	count := 0

	for idx, dir := range directions {
		dx, dy := dir.x, dir.y
		ddx, ddy := directions[(idx+1)%4].x, directions[(idx+1)%4].y
		left := grid[Coord{coord.y + dy, coord.x + dx}]
		right := grid[Coord{coord.y + ddy, coord.x + ddx}]
		mid := grid[Coord{coord.y + dy + ddy, coord.x + dx + ddx}]

		if (left != grid[coord] && right != grid[coord]) || (left == grid[coord] && right == grid[coord] && mid != grid[coord]) {
			count++
		}
	}

	return count
}

func walk(grid Grid, coord Coord, curr string, seen Seen) (int, int, int) {
	if seen[coord] || grid[coord] != curr {
		return 0, 0, 0
	}

	seen[coord] = true
	var area int = 1
	var perimeter int = 4 - getNeighbours(grid, coord)
	var side int = getCorners(grid, coord)

	for _, direction := range directions {
		a, p, s := walk(grid, Coord{coord.y + direction.y, coord.x + direction.x}, curr, seen)
		area, perimeter, side = area+a, perimeter+p, side+s
	}

	return area, perimeter, side
}

func main() {
	var partA, partB int
	contents, _ := os.ReadFile("data/day12.txt")
	lines := strings.Split(string(contents), "\n")
	grid := make(Grid)
	seen := make(Seen)
	h := len(lines)
	w := len(lines[0])

	for row, line := range lines {
		for col, char := range strings.Split(line, "") {
			grid[Coord{row, col}] = char
		}
	}

	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			area, perimeter, side := walk(grid, Coord{row, col}, grid[Coord{row, col}], seen)
			partA += area * perimeter
			partB += area * side
		}
	}

	fmt.Println("Part A:", partA)
	fmt.Println("Part B:", partB)
}
