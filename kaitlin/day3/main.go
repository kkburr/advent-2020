package main

import (
	"fmt"
	"github.com/kkburr/advent-2020/kaitlin/utils"
	"io/ioutil"
	"strings"
	"sync"
)

type slope struct {
	axis []int
}

type slopes struct {
	slopes []slope
}

func main() {
	f, err := ioutil.ReadFile("input")
	utils.CheckError(err)
	data := strings.Split(string(f), "\n")
	s := slopes{[]slope{slope{[]int{1, 1}}, slope{[]int{3, 1}}, slope{[]int{5, 1}}, slope{[]int{7, 1}}, slope{[]int{1, 2}}}}
	l := len(s.slopes)
	counts := make(chan int, l)
	var wg sync.WaitGroup
	wg.Add(l)
	for _, v := range s.slopes {
		go scan(&wg, v, data, counts)
	}
	wg.Wait()
	product := 1
	for i := 0; i < l; i++ {
		select {
		case val := <-counts:
			product *= val
		}
	}
	fmt.Printf("Part 2 answer: %v", product)
}

func scan(wg *sync.WaitGroup, s slope, data []string, counts chan int) {
	x := 0
	y := 0
	xSlope := s.axis[0]
	ySlope := s.axis[1]
	t := 0
	for i, row := range data {
		if row == "" {
			fmt.Println(i)
			continue
		}
		if (y % ySlope) == 0 {
			if row[x] == '#' {
				t += 1
			}
			x = (x + xSlope) % len(row)
		}
		y++
	}
	counts <- t
	fmt.Printf("Trees hit in slope x %v & y %v: %v\n", xSlope, ySlope, t)
	wg.Done()
}
