package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	checkErr(err)
	numbs := convertToSliceOfInts(data)

	for i, v := range numbs {
		for ii := i + 1; ii < len(numbs); ii++ {
			vv := numbs[ii]
			sum := v + vv
			if sum == 2020 {
				fmt.Println("Part 1 answer:")
				fmt.Println(v * vv)
			} else if sum < 2020 {
				for iii := i; iii < len(numbs); iii++ {
					vvv := numbs[iii]
					if sum+vvv == 2020 {
						fmt.Println("Part 2 answer:")
						fmt.Println(v * vv * vvv)
						return
					}
				}
			}
		}
	}
}

func convertToSliceOfInts(a []byte) []int {
	//TODO: find way to convert byte array to int array w/o splitting into string array first
	fileData := strings.Split(string(a), "\n")
	digits := make([]int, 0)

	for _, v := range fileData {
		if v != "" {
			i, err := strconv.Atoi(v)
			checkErr(err)

			if i > 0 {
				digits = append(digits, i)
			}
		}
	}
	return digits
}

//TODO: pull this out into utils file
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
