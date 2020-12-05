package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in, _ := os.Open("input")
	scanner := bufio.NewScanner(in)
	ids := make([]float64, 0)
	for scanner.Scan() {
		raw := scanner.Text()
		row := process(raw[:7], 'F', 0.0, 127.0)
		column := process(raw[7:], 'L', 0.0, 7.0)
		id := row*8 + column
		ids = append(ids, id)
	}
	sort.Float64s(ids)
	fmt.Printf("Part 1: %v\n", ids[len(ids)-1])

	for i := 0; i < len(ids)-1; i++ {
		id := ids[i]
		if (ids[i+1] - 1) != id {
			fmt.Printf("Part 2: %v\n", id+1)
		}
	}
}

func process(letters string, letter byte, first, second float64) float64 {
	if len(letters) == 1 {
		if letters[0] == letter {
			return first
		} else {
			return second
		}

	}
	next := math.Round((second - first) / 2)
	if letters[0] == letter {
		second = second - next
	} else {
		first = first + next
	}

	return process(letters[1:], letter, first, second)
}
