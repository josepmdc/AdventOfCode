package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := util.GetData("8th/input")
	part1(data)
	part2(data)
}

func part1(instructions []string) {
	count, _ := runProgram(instructions)
	fmt.Printf("Part 1: %d\n", count)
}

func part2(instructions []string) {
	for i, instruction := range instructions {
		cp := make([]string, len(instructions))
		copy(cp, instructions)

		inst := strings.Split(instruction, " ")
		if inst[0] == "jmp" {
			cp[i] = fmt.Sprintf("nop %s", inst[1])
		} else if inst[0] == "nop" {
			cp[i] = fmt.Sprintf("jmp %s", inst[1])
		}

		if count, isFinite := runProgram(cp); isFinite {
			fmt.Printf("Part 2: %d, changed the statement on position %d\n", count, i+1)
			return
		}
	}
}

func runProgram(instructions []string) (int, bool) {
	count, i := 0, 0
	instructionsCache := make(map[int]bool)
	for i < len(instructions) {
		if instructionsCache[i] {
			return count, false
		}
		instructionsCache[i] = true
		i += parseInstruction(instructions[i], &count)
	}
	return count, true
}

func parseInstruction(instruction string, count *int) int {
	i := strings.Split(instruction, " ")
	val, _ := strconv.Atoi(i[1])
	switch i[0] {
	case "jmp":
		return val
	case "acc":
		(*count) += val
		break
	}
	return 1
}
