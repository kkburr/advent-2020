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
	waypoint     map[byte]int
}

func main() {
	c := initCompass()
	part1 := c.part1()
	fmt.Printf("Part 1: %v\n", part1)
	part2 := c.part2()
	fmt.Printf("Part 2: %v\n", part2)
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
		waypoint:     map[byte]int{'E': 10, 'N': 1, 'W': 0, 'S': 0},
	}
	return &c
}

func (c *compass) part1() float64 {
	for _, v := range c.instructions {
		dir := v[0]
		amt, _ := strconv.Atoi(v[1:])
		c.move(dir, amt)
	}
	eastWest := c.log['E'] - c.log['W']
	northSouth := c.log['N'] - c.log['S']
	return math.Abs(float64(eastWest)) + math.Abs(float64(northSouth))
}

func (c *compass) part2() float64 {
	c.log = map[byte]int{'E': 0, 'N': 0, 'W': 0, 'S': 0}
	for _, v := range c.instructions {
		dir := v[0]
		amt, _ := strconv.Atoi(v[1:])
		c.move2(dir, amt)
	}
	eastWest := c.log['E'] - c.log['W']
	northSouth := c.log['N'] - c.log['S']
	return math.Abs(float64(eastWest)) + math.Abs(float64(northSouth))
}

func (c *compass) move(dir byte, amt int) {
	switch {
	case dir == 'L' || dir == 'R':
		c.F = c.changeDirections(dir, c.F, amt/90)
	case dir == 'F':
		c.log[c.F] = c.log[c.F] + amt
	default:
		c.log[dir] = c.log[dir] + amt
	}
}

func (c *compass) move2(dir byte, amt int) {
	switch {
	case dir == 'L' || dir == 'R':
		tempMap := map[byte]int{'E': 10, 'N': 1, 'W': 0, 'S': 0}
		for k, v := range c.waypoint {
			newDir := c.changeDirections(dir, k, amt/90)
			tempMap[newDir] = v
		}
		c.waypoint = tempMap
	case dir == 'F':
		for k, v := range c.waypoint {
			if v > 0 {
				c.log[k] = c.log[k] + (amt * v)
			}
		}
	default:
		oppositeDir := c.changeDirections('L', dir, 2)
		oppositeAmt := c.waypoint[oppositeDir]
		if oppositeAmt > 0 {
			newAmt := oppositeAmt - amt
			if newAmt < 0 {
				c.waypoint[dir] = int(math.Abs(float64(newAmt)))
				c.waypoint[oppositeDir] = 0
			} else {
				c.waypoint[oppositeDir] = newAmt
			}
		} else {
			c.waypoint[dir] = c.waypoint[dir] + amt
		}
	}
}

func (c *compass) changeDirections(dir, position byte, numberOfTurns int) byte {
	var index int

	for i, v := range c.directions {
		if v == position {
			index = i
		}
	}

	if dir == 'L' {
		index -= numberOfTurns
	} else if dir == 'R' {
		index += numberOfTurns
	}

	index = adjustIndex(index)
	return c.directions[index]
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
