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

func count(visited map[Coord]int, size int) int {
	count := 0

	for key := range visited {
		if key.x >= 0 && key.x < size && key.y >= 0 && key.y < size {
			count++
		}
	}

	return count
}

func main() {
	contents, _ := os.ReadFile("data/day08.txt")
	lines := strings.Split(string(contents), "\n")
	grid := make(map[Coord]string)
	size := len(lines[0])

	for row, line := range lines {
		for col, char := range strings.Split(line, "") {
			if char != "." {
				grid[Coord{col, row}] = char
			}
		}
	}

	visited_a := make(map[Coord]int)
	visited_b := make(map[Coord]int)

	for coord, value := range grid {
		for c, v := range grid {
			if v == value && c.x != coord.x && c.y != coord.y {
				dx := coord.x - c.x
				dy := coord.y - c.y

				visited_a[Coord{coord.x + dx, coord.y + dy}] += 1
				visited_a[Coord{c.x - dx, c.y - dy}] += 1

				for i := 0; i < 25; i++ {
					visited_b[Coord{coord.x + (dx * i), coord.y + (dy * i)}] += 1
					visited_b[Coord{c.x - (dx * i), c.y - (dy * i)}] += 1
				}
			}
		}

	}

	fmt.Println("Part A:", count(visited_a, size))
	fmt.Println("Part B:", count(visited_b, size))
}
