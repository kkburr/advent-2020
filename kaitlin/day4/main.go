package main

import (
	"fmt"
	"github.com/kkburr/advent-2020/kaitlin/utils"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passwordRules struct {
	year   *regexp.Regexp
	height *regexp.Regexp
	hair   *regexp.Regexp
	eye    *regexp.Regexp
	pid    *regexp.Regexp
	data   []string
}

func main() {
	f, err := ioutil.ReadFile("input")
	utils.CheckError(err)
	fmt.Printf("Total valid Passports: %v", Process(f))
}

func Process(f []byte) int {
	count := 0
	re := regexp.MustCompile(`\s+`)
	data := strings.Split(string(f), "\n\n")
	rule := &passwordRules{
		year:   regexp.MustCompile(`\A\d{4}\z`),
		height: regexp.MustCompile(`\A(\d+)(cm|in)\z`),
		hair:   regexp.MustCompile(`\A#[0-9a-f]{6}\z`),
		eye:    regexp.MustCompile(`\A(amb|blu|brn|gry|grn|hzl|oth){1}\z`),
		pid:    regexp.MustCompile(`\A\d{9}\z`),
		data:   data,
	}

	for _, v := range rule.data {
		if v == "" {
			continue
		}
		fields := make(map[string]string)
		s := re.Split(v, -1)
		for _, str := range s {
			if str == "" {
				continue
			}
			parts := strings.Split(str, ":")
			fields[parts[0]] = parts[1]
		}
		if rule.isValid(fields) {
			count += 1
		}
	}
	return count
}

func (r *passwordRules) isValid(m map[string]string) bool {
	if len(m) < 7 || len(m) == 7 && m["cid"] != "" {
		return false
	}
	count := 0

	byr := m["byr"]
	if r.year.MatchString(byr) {
		a, _ := strconv.Atoi(byr)
		if a >= 1920 && a <= 2002 {
			count++
		}
	}

	iyr := m["iyr"]
	if r.year.MatchString(iyr) {
		a, _ := strconv.Atoi(iyr)
		if a >= 2010 && a <= 2020 {
			count++
		}
	}

	eyr := m["eyr"]
	if r.year.MatchString(eyr) {
		a, _ := strconv.Atoi(eyr)
		if a >= 2020 && a <= 2030 {
			count++
		}
	}

	hgt := m["hgt"]
	if r.height.MatchString(hgt) {
		arr := r.height.FindStringSubmatch(hgt)
		if arr[2] == "cm" {
			l, _ := strconv.Atoi(arr[1])
			if l >= 150 && l <= 193 {
				count++
			}
		} else if arr[2] == "in" {
			l, _ := strconv.Atoi(arr[1])
			if l >= 59 && l <= 76 {
				count++
			}
		}
	}

	hcl := m["hcl"]
	if r.hair.MatchString(hcl) {
		count++
	}

	ecl := m["ecl"]
	if r.eye.MatchString(ecl) {
		count++
	}

	p := m["pid"]
	if r.pid.MatchString(p) {
		count++
	}
	return count == 7
}
