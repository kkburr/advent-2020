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

	var wg sync.WaitGroup
	slopes := slopes{[]slope{slope{[]int{1, 1}}, slope{[]int{3, 1}}, slope{[]int{5, 1}}, slope{[]int{7, 1}}, slope{[]int{1, 2}}}}
	counts := make(chan int, len(slopes.slopes))

	for _, v := range slopes.slopes {
		wg.Add(1)
		go scan(&wg, v, data, counts)
	}

	wg.Wait()

	product := 1
	for i := 0; i <= len(slopes.slopes)-1; i++ {
		select {
		case val := <-counts:
			fmt.Println(product)
			fmt.Println(val)
			product = product * val
		}
	}
	fmt.Printf("Part 2 answer: %v", product)
}

func trimX(l, p int) int {
	for {
		if p >= l {
			p = p - l
		} else {
			return p
		}
	}
}

func scan(wg *sync.WaitGroup, s slope, data []string, counts chan int) {
	x := 0
	y := 0
	xSlope := s.axis[0]
	ySlope := s.axis[1]
	treeCounter := 0
	for _, row := range data {
		row = strings.TrimSpace(row)
		if row != "" {
			if (y % ySlope) == 0 {
				x = trimX(len(row), x)
				if "#" == string(row[x]) {
					treeCounter = treeCounter + 1
				}
			}
			x = x + xSlope
			y = y + ySlope
		}
	}
	counts <- treeCounter
	fmt.Printf("Trees hit in slope x %v & y %v: %v\n", xSlope, ySlope, treeCounter)
	wg.Done()
}
