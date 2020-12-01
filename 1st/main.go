package main

import (
	"aoc2020/util"
	"fmt"
)

func main() {
	var data = util.GetData("1st/input")
	part1(data)
	part2(data)
}

func part1(data []int) {
	for i := 0; i < len(data); i++ {
		num1 := data[i]
		for j := i + 1; j < len(data); j++ {
			num2 := data[j]
			if num1+num2 == 2020 {
				fmt.Printf("Part 1 -> %d\n", num1*num2)
				return
			}
		}
	}
}

func part2(data []int) {
	for i := 0; i < len(data); i++ {
		num1 := data[i]
		for j := i + 1; j < len(data); j++ {
			num2 := data[j]
			for k := j + 1; k < len(data); k++ {
				num3 := data[k]
				if num1+num2+num3 == 2020 {
					fmt.Printf("Part 2 -> %d\n", num1*num2*num3)
					return
				}
			}
		}
	}
}
