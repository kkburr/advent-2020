package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("input")
	scanner := bufio.NewScanner(in)
	instructions := make([]map[string]string, 0)

	for scanner.Scan() {
		t := scanner.Text()
		instruction := make(map[string]string)
		a := strings.Split(t, " ")
		instruction["op"] = a[0]
		instruction["arg"] = a[1]
		instructions = append(instructions, instruction)
	}

	visited := make(map[int]struct{})
	i := 0
	acc := 0

	fmt.Println(move(instructions, visited, i, acc))
}

func move(arr []map[string]string, visited map[int]struct{}, i, acc int) int {
	for {
		if _, ok := visited[i]; ok {
			return acc
		} else {
			visited[i] = struct{}{}
		}

		m := arr[i]
		op := m["op"]
		arg, _ := strconv.Atoi(m["arg"])

		switch {
		case op == "nop":
			i = noOp(i)
			continue
		case op == "acc":
			i, acc = accum(i, arg, acc)
			continue
		case op == "jmp":
			i = jmp(i, arg)
			continue
		}
	}
	return acc
}

func noOp(i int) int {
	i += 1
	return i
}

func accum(i, arg, acc int) (int, int) {
	acc += arg
	i += 1
	return i, acc
}

func jmp(i, arg int) int {
	i += arg
	return i
}
