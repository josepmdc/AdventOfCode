package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type modifier struct {
	pos int
	val rune
}

func main() {
	data := util.GetData("14th/input_test")
	part1(data)
}

func part1(data []string) {
	var mask string
	var modifiers []modifier
	memory := make(map[int]int)
	for _, inst := range data {
		if strings.Contains(inst, "mask") {
			mask = strings.Split(data[0], " = ")[1]
			for i, char := range mask {
				if char != 'X' {
					modifiers = append(modifiers, modifier{pos: i, val: char})
				}
			}
		} else {
			r := regexp.MustCompile(`^mem\[(.*?)\] = ([0-9]+)$`).FindStringSubmatch(inst)
			pos, _ := strconv.Atoi(r[1])
			v, _ := strconv.Atoi(r[2])
			val := []rune("00000000000000000000000000000000000000000000000000000000" +
				strconv.FormatInt(int64(v), 2))
			val = val[len(val)-len(mask):]
			for _, m := range modifiers {
				val[m.pos] = m.val
			}
			tmp, _ := strconv.ParseInt(string(val), 2, 64)
			memory[pos] = int(tmp)
		}
	}
	sum := 0
	for _, v := range memory {
		sum += v
	}
	fmt.Printf("Part 1: %d\n", sum)
}
