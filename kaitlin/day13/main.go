package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input")
	a := strings.Split(string(data), "\n")
	t, _ := strconv.ParseFloat(a[0], 64)
	part1 := findBusRoutes(a[1], t)

	fmt.Printf("Part 1: %v\n", part1)
}

func findBusRoutes(str string, timestamp float64) float64 {
	temp := strings.Split(str, ",")
	diffs := make([]float64, 0)
	buses := make(map[float64]float64)
	for _, v := range temp {
		if v != "x" {
			id, _ := strconv.ParseFloat(v, 64)
			diff := timestamp / id
			key := diff - math.Floor(diff)
			buses[key] = id
			diffs = append(diffs, key)
		}
	}
	sort.Float64s(diffs)
	bus := buses[diffs[len(diffs)-1]]
	return (math.Round(timestamp/bus)*bus - timestamp) * bus
}
