package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("input")
	scanner := bufio.NewScanner(in)

	reg := regexp.MustCompile(`(\d+|[a-zA-Z]+\s[a-zA-Z]+)`)
	replacer := strings.NewReplacer("bags", "", "bag", "", ".", "")
	m1 := make(map[string]map[string]int)

	//create map of color rules
	for scanner.Scan() {
		raw := scanner.Text()
		str := replacer.Replace(raw)
		s := regexp.MustCompile("contain|,").Split(str, -1)

		m2 := make(map[string]int)

		for i, v := range s {
			v = strings.TrimSpace(v)
			if i == 0 {
				m1[v] = m2
			} else {
				a := reg.FindAllStringSubmatch(v, -1)
				if len(a) > 1 {
					color := a[1][0]
					instances, _ := strconv.Atoi(a[0][0])
					m2[color] = instances
				}
			}
		}
	}

	//recursively find count
	count := recur(m1, "shiny gold")
	fmt.Println(count)
}

func recur(m map[string]map[string]int, topColor string) int {
	counter := 0
	rules := m[topColor]
	if len(rules) == 0 {
		return 1
	}
	for k, v := range rules {
		r := recur(m, k)
		if r == 1 {
			counter += v
		} else {
			counter += (v + v*r)
		}
	}
	return counter
}
