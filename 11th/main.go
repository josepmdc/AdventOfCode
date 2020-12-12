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

func main() {
	data := util.GetData("11th/input")
	part1(data)
	part2(data)
}

func part1(data []string) {
	seatMap := generateSeatMap(data)
	xSize := len(data[0])
	ySize := len(data)
	for {
		tmpSeatMap := nextRound(seatMap, xSize, ySize)
		if reflect.DeepEqual(seatMap, tmpSeatMap) {
			fmt.Printf("Part 1: %d\n", countOccupiedSeats(seatMap))
			return
		}
		seatMap = tmpSeatMap
	}
}

func countOccupiedSeats(seatMap map[string]string) int {
	count := 0
	for _, state := range seatMap {
		if state == "#" {
			count++
		}
	}
	return count
}

func nextRound(seatMap map[string]string, xSize int, ySize int) map[string]string {
	adjecentCoordinates := getAdjecentCoordinates()
	tmpSeatMap := make(map[string]string)
	for k, v := range seatMap {
		tmpSeatMap[k] = v
	}
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			currentCoordinate := getCoordinates(x, y)
			occupiedCount := 0
			for _, adjecent := range adjecentCoordinates {
				adjecentCoordinate := getCoordinates(x+adjecent.x, y+adjecent.y)

				if seatMap[adjecentCoordinate] == "#" {
					occupiedCount++
				}
			}
			if seatMap[currentCoordinate] == "#" && occupiedCount >= 4 {
				tmpSeatMap[currentCoordinate] = "L"
			}
			if seatMap[currentCoordinate] == "L" && occupiedCount == 0 {
				tmpSeatMap[currentCoordinate] = "#"
			}
		}
	}
	return tmpSeatMap
}

func printSeatMap(seatMap map[string]string, xSize int, ySize int) {
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			coordinate := getCoordinates(x, y)
			fmt.Print(seatMap[coordinate])
		}
		fmt.Println()
	}
	fmt.Println()
}

func getAdjecentCoordinates() []coordinate {
	return []coordinate{
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
}

func getCoordinates(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func part2(data []string) {

}

func generateSeatMap(data []string) map[string]string {
	seats := make(map[string]string)
	for x, row := range data {
		for y, seat := range row {
			coordinates := getCoordinates(x, y)
			seats[coordinates] = string(seat)
		}
	}
	return seats
}
