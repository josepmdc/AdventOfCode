package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := util.GetData("14th/input")
	part1(data)
}

func part1(data []string) {
	var andMask int64
	var orMask int64
	memory := make(map[int]int64)
	for _, inst := range data {
		if strings.Contains(inst, "mask") {
			mask := strings.Split(inst, " = ")[1]
			andMask, _ = strconv.ParseInt(strings.ReplaceAll(mask, "X", "1"), 2, 64)
			orMask, _ = strconv.ParseInt(strings.ReplaceAll(mask, "X", "0"), 2, 64)
		} else {
			r := regexp.MustCompile(`^mem\[(.*?)\] = ([0-9]+)$`).FindStringSubmatch(inst)
			pos, _ := strconv.Atoi(r[1])
			v, _ := strconv.Atoi(r[2])
			val := int64(v)
			val &= andMask
			val |= orMask
			memory[pos] = val
		}
	}
	var sum int64
	for _, v := range memory {
		sum += v
	}
	fmt.Printf("Part 1: %d\n", sum)
}
