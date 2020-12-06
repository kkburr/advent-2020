package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("input")
	scanner := bufio.NewScanner(in)

	count := 0
	family := make(map[rune]struct{})
	firstRow := true

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			count += len(family)
			family = make(map[rune]struct{})
			firstRow = true
			continue
		} else if firstRow == true {
			for _, v := range raw {
				family[v] = struct{}{}
				firstRow = false
			}
		} else {
			//loop over row and add it to a map
			rowM := make(map[rune]struct{})
			for _, v := range raw {
				rowM[r] = struct{}{}
			}

			//loop over family map and delete entry if not in row map
			for k, _ := range family {
				if _, ok := rowM[k]; !ok {
					delete(family, k)
				}
			}
		}
	}
	count += len(family)
	fmt.Println(count)
}
