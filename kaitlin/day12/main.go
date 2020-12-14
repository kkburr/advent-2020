package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type compass struct {
	F            byte         // forward facing direction
	directions   []byte       // array of all possible directions
	instructions []string     // array of input instructions
	log          map[byte]int // record of moves in each direction
}

func main() {
	c := initCompass()
	part1 := c.part1()
	fmt.Println(part1)
}

func initCompass() *compass {
	in, _ := os.Open("input")
	scanner := bufio.NewScanner(in)
	instructions := make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		instructions = append(instructions, t)
	}
	c := compass{
		F:            'E',
		directions:   []byte{'W', 'N', 'E', 'S'},
		instructions: instructions,
		log:          map[byte]int{'E': 0, 'N': 0, 'W': 0, 'S': 0},
	}
	return &c
}

func (c *compass) part1() float64 {
	for _, v := range c.instructions {
		dir := v[0]
		amt, _ := strconv.Atoi(v[1:])
		c.move(dir, amt)
		fmt.Println(c.log)
	}
	eastWest := c.log['E'] - c.log['W']
	northSouth := c.log['N'] - c.log['S']
	return math.Abs(float64(eastWest)) + math.Abs(float64(northSouth))
}

func (c *compass) move(dir byte, amt int) {
	switch {
	case dir == 'L' || dir == 'R':
		c.changeDirections(dir, amt/90)
	case dir == 'F':
		c.log[c.F] = c.log[c.F] + amt
	default:
		c.log[dir] = c.log[dir] + amt
	}
}

func (c *compass) changeDirections(dir byte, numberOfTurns int) {
	var index int

	for i, v := range c.directions {
		if v == c.F {
			index = i
		}
	}

	if dir == 'L' {
		index -= numberOfTurns
	} else if dir == 'R' {
		index += numberOfTurns
	}

	index = adjustIndex(index)
	c.F = c.directions[index]
}

func adjustIndex(index int) int {
	if index < 0 {
		for {
			index += 4
			if index >= 0 {
				break
			}
		}
	} else if index >= 4 {
		for {
			index -= 4
			if index < 4 {
				break
			}
		}
	}
	return index
}
