package main

import (
	"aoc2020/util"
	"fmt"
	"math"
	"strconv"
)

type position struct {
	x         int
	y         int
	direction string
}

type move func(position, int) position

func main() {
	data := util.GetData("12th/input")
	part1(data)
	part2(data)
}

func part1(data []string) {
	moves := make(map[string]move)
	moves["N"] = func(p position, a int) position {
		p.y += a
		return p
	}
	moves["S"] = func(p position, a int) position {
		p.y -= a
		return p
	}
	moves["E"] = func(p position, a int) position {
		p.x += a
		return p
	}
	moves["W"] = func(p position, a int) position {
		p.x -= a
		return p
	}
	moves["L"] = func(p position, a int) position {
		p.direction = getRotations()[p.direction+"L"][getRotationPosition(a)]
		return p
	}
	moves["R"] = func(p position, a int) position {
		p.direction = getRotations()[p.direction+"R"][getRotationPosition(a)]
		return p
	}
	moves["F"] = func(p position, a int) position {
		return moves[p.direction](p, a)
	}

	p := position{x: 0, y: 0, direction: "E"}
	for _, instruction := range data {
		direction := string(instruction[0])
		shift, _ := strconv.Atoi(instruction[1:])
		p = moves[direction](p, shift)
	}
	fmt.Printf("Part 1: %v\n", math.Abs(float64(p.x))+math.Abs(float64(p.y)))
}

func part2(data []string) {
	waypoint := position{x: 10, y: 1}
	moves := make(map[string]move)
	moves["N"] = func(p position, a int) position {
		waypoint.y += a
		return p
	}
	moves["S"] = func(p position, a int) position {
		waypoint.y -= a
		return p
	}
	moves["E"] = func(p position, a int) position {
		waypoint.x += a
		return p
	}
	moves["W"] = func(p position, a int) position {
		waypoint.x -= a
		return p
	}
	moves["L"] = func(p position, a int) position {
		for i := 0; i < a/90; i++ {
			tmpY := waypoint.y
			waypoint.y = waypoint.x
			waypoint.x = tmpY * -1
		}
		return p
	}
	moves["R"] = func(p position, a int) position {
		for i := 0; i < a/90; i++ {
			tmpX := waypoint.x
			waypoint.x = waypoint.y
			waypoint.y = tmpX * -1
		}
		return p
	}
	moves["F"] = func(p position, a int) position {
		p.x += a * waypoint.x
		p.y += a * waypoint.y
		return p
	}

	p := position{x: 0, y: 0}
	for _, instruction := range data {
		direction := string(instruction[0])
		shift, _ := strconv.Atoi(instruction[1:])
		p = moves[direction](p, shift)
	}
	fmt.Printf("Part 2: %v\n", math.Abs(float64(p.x))+math.Abs(float64(p.y)))
}

func getRotationPosition(degrees int) int {
	if degrees == 90 {
		return 0
	} else if degrees == 180 {
		return 1
	} else {
		return 2
	}
}

func getRotations() map[string][]string {
	return map[string][]string{
		"NL": {"W", "S", "E"},
		"SL": {"E", "N", "W"},
		"EL": {"N", "W", "E"},
		"WL": {"S", "E", "W"},
		"NR": {"E", "S", "W"},
		"SR": {"W", "N", "E"},
		"ER": {"S", "W", "N"},
		"WR": {"N", "E", "S"},
	}
}
