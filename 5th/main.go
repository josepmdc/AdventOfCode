package main

import (
	"aoc2020/util"
	"fmt"
)

func main() {
	//testSeats := []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	seats := util.GetData("5th/input")
	part1(seats)
	part2(seats)
}

type bounds struct {
	lower int
	upper int
}

type seat struct {
	row    int
	column int
}

func part1(seats []string) {
	highestID := 0
	for _, seat := range seats {
		column := getColumn(seat)
		row := getRow(seat)
		seatID := (row * 8) + column
		if seatID > highestID {
			highestID = seatID
		}
	}
	fmt.Printf("Part 1: The highest seat ID is %d\n", highestID)
}

func part2(seatCodes []string) {
	seats := getSeats(seatCodes)
	// We skipt the first 8 rows and the last 10
	for i := 9; i < 117; i++ {
		for j := 0; j < 7; j++ {
			s := seat{row: i, column: j}
			occupied := seats[s]
			if !occupied {
				seatID := (s.row * 8) + s.column
				fmt.Printf("Part 2: row = %d, column = %d, seatID = %d\n", s.row, s.column, seatID)
				return
			}
		}
	}
}

func getSeats(seats []string) map[seat]bool {
	s := make(map[seat]bool, 0)
	for _, seatCode := range seats {
		key := seat{row: getRow(seatCode), column: getColumn(seatCode)}
		s[key] = true
	}
	return s
}

func getColumn(seat string) int {
	b := bounds{lower: 0, upper: 7}
	for _, letter := range seat[len(seat)-3:] {
		middle := (b.upper + b.lower) / 2
		if letter == 'L' {
			b.upper = middle
		} else {
			b.lower = middle + 1
		}
	}
	return b.lower
}

func getRow(seat string) int {
	b := bounds{lower: 0, upper: 127}
	for _, letter := range seat[:len(seat)-3] {
		middle := (b.upper + b.lower) / 2
		if letter == 'F' {
			b.upper = middle
		} else {
			b.lower = middle + 1
		}
	}
	return b.lower
}
