package main

import (
	"fmt"
	"math"
	"os"
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

		nums[i] = int(n)
	}

	return nums
}

func size(a int) int {
	var count int
	c := a

	for c > 0 {
		c /= 10
		count++
	}

	return count
}

func split(num int, s int) (int, int) {
	l := num / int(math.Pow10(s))
	r := num % int(math.Pow10(s))
	return l, r
}

func blink(store map[int]int) map[int]int {
	new := make(map[int]int)

	for num := range store {
		if num == 0 {
			new[1] += store[num]
			continue
		}

		s := size(num)
		x := store[num]

		if s%2 == 0 {
			l, r := split(num, int(s/2))
			new[l] += x
			new[r] += x

		} else {
			new[num*2024] += x
		}
	}

	return new
}

func run(store map[int]int, n int) int {
	var count int

	for i := 0; i < n; i++ {
		store = blink(store)
	}

	for _, v := range store {
		count += v
	}

	return count
}

func main() {
	contents, _ := os.ReadFile("data/day11.txt")
	nums := getNums(strings.Split(string(contents), " "))
	store := make(map[int]int)

	for _, num := range nums {
		store[num] = 1
	}

	fmt.Println("Part A:", run(store, 25))
	fmt.Println("Part B:", run(store, 75))
}
