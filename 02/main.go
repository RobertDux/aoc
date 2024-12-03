package main

import (
	"fmt"
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

		nums[i] = n
	}

	return nums
}

func getDir(a int, b int) int {
	if a > b {
		return -1
	}

	if a < b {
		return 1
	}

	return 0
}

func differ(a int, b int) bool {
	if a-b > 3 || b-a > 3 {
		return true
	}

	return false
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func getPerms(line []int) [][]int {
	perms := make([][]int, len(line))

	for i := 0; i < len(line); i++ {
		perm := RemoveIndex(line, i)
		perms[i] = perm
	}

	return perms
}

func isSafe(line []int) bool {
	dir := getDir(line[0], line[1])

	if dir == 0 {
		return false
	}

	for i := 1; i < len(line); i++ {
		if getDir(line[i-1], line[i]) != dir {
			return false
		}

		if differ(line[i-1], line[i]) {
			return false
		}
	}

	return true
}

func main() {
	var count_a, count_b int
	data, _ := os.ReadFile("data/day02.txt")
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		nums := getNums(strings.Split(line, " "))

		if isSafe(nums) {
			count_a++
			count_b++
		} else {
			perms := getPerms(nums)
			for _, perm := range perms {
				if isSafe(perm) {
					count_b++
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", count_a)
	fmt.Println("Part 2:", count_b)
}
