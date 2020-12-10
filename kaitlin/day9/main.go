package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	in, _ := os.Open("input")
	scanner := bufio.NewScanner(in)
	x := make([]int, 0)
	preambleLen := 25

	for scanner.Scan() {
		d, _ := strconv.Atoi(scanner.Text())
		x = append(x, d)
	}

	part1 := findPart1(x, preambleLen)
	fmt.Printf("Part 1: %v\n", part1)
	part2 := findPart2(x, part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func findPart1(digits []int, l int) int {
	p := findPreamble(digits, l)
	pointer := digits[l]
	for a := 0; a < l-1; a++ {
		for b := 1; b < l; b++ {
			if p[a]+p[b] == pointer {
				return findPart1(digits[1:], l)
			}
		}
	}
	return pointer
}

func findPreamble(digits []int, l int) []int {
	a := digits[0:l]
	n := make([]int, l)
	copy(n, a)
	sort.Ints(n)
	return n
}

func findPart2(digits []int, sum int) int {
	startIndex := 0
	pointerSum := 0
	for index := 0; index < len(digits); index++ {
		pointerSum += digits[index]
		if pointerSum > sum {
			for {
				pointerSum -= digits[startIndex]
				startIndex++
				if pointerSum <= sum {
					break
				}
			}
		}
		if pointerSum == sum {
			subRange := digits[startIndex:index]
			sort.Ints(subRange)
			return subRange[0] + subRange[len(subRange)-1]
		}
	}
	return 0
}
