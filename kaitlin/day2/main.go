package main

import (
	"fmt"
	"github.com/kkburr/advent-2020/kaitlin/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	utils.CheckError(err)

	part1, part2 := findGoodPasswords(data)
	fmt.Println("Part 1:")
	fmt.Println(len(part1))
	fmt.Println("Part 2:")
	fmt.Println(len(part2))
}

func findGoodPasswords(a []byte) ([]string, []string) {
	data := strings.Split(string(a), "\n")
	part1 := make([]string, 0)
	part2 := make([]string, 0)
	for _, v := range data {
		d := strings.Split(v, ":")
		if len(d) > 1 {
			rule := d[0]
			pw := d[1]
			p1, p2 := isGoodPassword(rule, pw)
			if p1 {
				part1 = append(part1, pw)
			}
			if p2 {
				part2 = append(part2, pw)
			}
		}
	}
	return part1, part2
}

func isGoodPassword(rule, pw string) (bool, bool) {
	r1 := strings.Split(rule, " ")
	letter := r1[1]
	r2 := strings.Split(r1[0], "-")
	lower, _ := strconv.Atoi(r2[0])
	upper, _ := strconv.Atoi(r2[1])

	part1 := false
	count := numberOfInstances(letter, pw)
	if count >= lower && count <= upper {
		part1 = true
	}

	part2 := false
	//string starts with a space, so we don't have to offset the indexes
	pos1 := string(pw[lower])
	pos2 := string(pw[upper])
	if (pos1 == letter && pos2 != letter) || (pos1 != letter && pos2 == letter) {
		part2 = true
	}

	return part1, part2
}

func numberOfInstances(letter, pw string) int {
	count := 0
	for _, v := range pw {
		if string(v) == letter {
			count++
		}
	}
	return count
}
