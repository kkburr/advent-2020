package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters := scanInput()
	sort.Ints(adapters)
	part1 := findPart1(adapters)
	fmt.Printf("Part 1: %v\n", part1)
}

func scanInput() []int {
	in, _ := os.Open("input")
	scanner := bufio.NewScanner(in)
	adapters := make([]int, 0)
	for scanner.Scan() {
		t := scanner.Text()
		d, _ := strconv.Atoi(t)
		adapters = append(adapters, d)
	}
	return adapters
}

func findPart1(adapters []int) int {
	diffOnes := 0
	diffThrees := 0
	prev := 0
	for i, v := range adapters {
		if i > 0 {
			prev = adapters[i-1]
		}
		diff := v - prev
		if diff == 1 {
			diffOnes++
		} else if diff == 3 {
			diffThrees++
		}
	}
	diffThrees++
	return diffOnes * diffThrees
}
