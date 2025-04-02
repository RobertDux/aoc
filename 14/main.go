package main

import (
	"fmt"
	"os"
	"strings"
)

func parse() [][]int {
	var robots [][]int
	contents, _ := os.ReadFile("data/day14.txt")
	lines := strings.Split(string(contents), "\n")

	for _, line := range lines {
		var x, y, dx, dy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &dx, &dy)
		robot := []int{x, y, dx, dy}
		robots = append(robots, robot)
	}

	return robots
}

func check(robots [][]int, w int, h int) int {
	q := make(map[int]int)

	for _, robot := range robots {
		x, y := robot[0], robot[1]

		if x < w/2 {
			if y < h/2 {
				q[0] += 1
			}

			if y > h/2 {
				q[2] += 1
			}
		}

		if x > w/2 {
			if y < h/2 {
				q[1] += 1
			}

			if y > h/2 {
				q[3] += 1
			}
		}
	}

	return q[0] * q[1] * q[2] * q[3]
}

func move(robots [][]int, w int, h int, n int) [][]int {
	for _, robot := range robots {
		dx, dy := robot[2], robot[3]
		robot[0] += (n * dx) % w
		robot[1] += (n * dy) % h
		robot[0] = (w + robot[0]) % w
		robot[1] = (h + robot[1]) % h
	}

	return robots
}

type Coord struct {
	x int
	y int
}

func overlap(robots [][]int) bool {
	seen := make(map[Coord]int)

	for _, robot := range robots {
		seen[Coord{robot[0], robot[1]}] += 1
	}

	return len(robots) != len(seen)
}

func partA(robots [][]int, w int, h int) int {
	robots = move(robots, w, h, 100)
	return check(robots, w, h)
}

func partB(robots [][]int, w int, h int) int {
	for i := 1; i < 10_000; i++ {
		robots = move(robots, w, h, 1)
		if !overlap(robots) {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println("Part A:", partA(parse(), 101, 103))
	fmt.Println("Part B:", partB(parse(), 101, 103))
}
