package main

import (
	"aoc2020/util"
	"fmt"
)

type Move struct {
	left int
	down int
}

func main() {
	data := util.GetData("3rd/input")
	part1(data)
	part2(data)
}

func part1(data []string) {
	trees := countTrees(data, Move{3, 1})
	fmt.Println(trees)
}

func part2(data []string) {
	moves := []Move{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	result := 1
	for _, move := range moves {
		result *= countTrees(data, move)
	}
	fmt.Println(result)
}

func countTrees(data []string, move Move) int {
	trees := 0
	position := 0
	for i := move.down; i < len(data); i += move.down {
		line := data[i]

		position += move.left
		if position >= len(line) {
			position -= len(line)
		}

		if line[position] == '#' {
			trees++
		}
	}
	return trees
}
