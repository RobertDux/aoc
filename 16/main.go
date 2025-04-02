package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Coord struct {
	y int
	x int
}

type Grid map[Coord]string

var directions = []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func parse(contents string) (Grid, Coord, Coord) {
	grid := make(Grid)
	var start, end Coord

	for y, row := range strings.Split(contents, "\n") {
		for x, char := range row {
			if string(char) == "S" {
				start = Coord{y, x}
			}

			if string(char) == "E" {
				end = Coord{y, x}
				grid[Coord{y, x}] = "."
				continue
			}

			grid[Coord{y, x}] = string(char)
		}
	}

	return grid, start, end
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func canTravel(g Grid, v map[Coord]bool, c Coord) bool {
	return g[c] == "." && !v[c]
}

func neighbours(g Grid, v map[Coord]bool, c Coord) []Coord {
	
}

func dijkstra(g Grid, start Coord, end Coord) int {
	visited := make(map[Coord]bool)
	queue := []Coord{start}
	times := make(map[Coord]int)

	for {
		next := queue[0]
		queue = queue[1:]
		time := times[next]
		
		if visited[next] {
			continue
		}

		visited[next] = true

		if next == end {
			return time
		}

		for _, n := range neighbours(next) {

		}
	}

}

func main() {
	var partB int
	contents, _ := os.ReadFile("data/day16.txt")
	grid, start, end := parse(string(contents))

	fmt.Println("Part A:", solve(grid, start, end))
	fmt.Println("Part B:", partB)
}
