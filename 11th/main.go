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
		if reflect.DeepEqual(tmpSeatMap, seatMap) {
			fmt.Printf("Part 1: %d\n", countOccupiedSeats(seatMap))
			return
		}
		seatMap = tmpSeatMap
	}
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

func nextRound(seatMap map[coordinate]string, xSize int, ySize int) map[coordinate]string {
	adjecentCoordinates := getAdjecentCoordinates()
	tmpSeatMap := make(map[coordinate]string)
	for k, v := range seatMap {
		occupiedCount := 0
		for _, adjecent := range adjecentCoordinates {
			adjecentCoordinate := coordinate{x: k.x + adjecent.x, y: k.y + adjecent.y}

			if seatMap[adjecentCoordinate] == "#" {
				occupiedCount++
			}
		}
		if seatMap[k] == "#" && occupiedCount >= 4 {
			v = "L"
		}
		if seatMap[k] == "L" && occupiedCount == 0 {
			v = "#"
		}
		tmpSeatMap[k] = v
	}
	return tmpSeatMap
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

func getAdjecentCoordinates() []coordinate {
	return []coordinate{{0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}}
}

func getCoordinates(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func part2(data []string) {

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
