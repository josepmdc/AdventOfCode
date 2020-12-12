package main

import (
	"aoc2020/util"
	"fmt"
	"reflect"
)

type coordinate struct {
	x int
	y int
}

type adjMethod func(x, y int, adj coordinate) coordinate

func main() {
	data := util.GetData("11th/input")
	seatMap := generateSeatMap(data)
	xSize := len(data[0])
	ySize := len(data)
	part1(seatMap, xSize, ySize)
	part2(seatMap, xSize, ySize)
}

func part1(seatMap map[coordinate]string, xSize, ySize int) {
	adjMethod := func(x, y int, adj coordinate) coordinate {
		return coordinate{x: x + adj.x, y: y + adj.y}
	}
	result := runSimulation(seatMap, xSize, ySize, 4, adjMethod)
	fmt.Printf("Part 1: %d\n", result)
}

func part2(seatMap map[coordinate]string, xSize, ySize int) {
	adjMethod := func(x, y int, adj coordinate) coordinate {
		adjCoordinate := coordinate{x: x + adj.x, y: y + adj.y}
		for seatMap[adjCoordinate] == "." && seatMap[adjCoordinate] != "" {
			adjCoordinate.x += adj.x
			adjCoordinate.y += adj.y
		}
		return adjCoordinate
	}
	result := runSimulation(seatMap, xSize, ySize, 5, adjMethod)
	fmt.Printf("Part 2: %d\n", result)
}

func runSimulation(seatMap map[coordinate]string, xSize, ySize, occupiedThreshold int, adjMethod adjMethod) int {
	for {
		tmpSeatMap := nextRound(seatMap, xSize, ySize, occupiedThreshold, adjMethod)
		if reflect.DeepEqual(tmpSeatMap, seatMap) {
			return countOccupiedSeats(seatMap)
		}
		seatMap = tmpSeatMap
	}
}

func nextRound(seatMap map[coordinate]string, xSize, ySize, occupiedThreshold int, adjMethod adjMethod) map[coordinate]string {
	adjecentCoordinates := []coordinate{{0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}}
	tmpSeatMap := make(map[coordinate]string)

	for k, v := range seatMap {
		occupiedCount := 0
		for _, adjecent := range adjecentCoordinates {
			adjecentCoordinate := adjMethod(k.x, k.y, adjecent)
			if seatMap[adjecentCoordinate] == "#" {
				occupiedCount++
			}
		}
		if seatMap[k] == "#" && occupiedCount >= occupiedThreshold {
			v = "L"
		}
		if seatMap[k] == "L" && occupiedCount == 0 {
			v = "#"
		}
		tmpSeatMap[k] = v
	}
	return tmpSeatMap
}

func countOccupiedSeats(seatMap map[coordinate]string) int {
	count := 0
	for _, state := range seatMap {
		if state == "#" {
			count++
		}
	}
	return count
}

func printSeatMap(seatMap map[coordinate]string, xSize int, ySize int) {
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			c := coordinate{x: x, y: y}
			fmt.Print(seatMap[c])
		}
		fmt.Println()
	}
	fmt.Println()
}

func generateSeatMap(data []string) map[coordinate]string {
	seats := make(map[coordinate]string)
	for x, row := range data {
		for y, seat := range row {
			coordinates := coordinate{x: x, y: y}
			seats[coordinates] = string(seat)
		}
	}
	return seats
}
