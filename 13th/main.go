package main

import (
	"aoc2020/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type trip struct {
	busID float64
	time  float64
}

func main() {
	data := util.GetData("13th/input")
	time, _ := strconv.ParseFloat(data[0], 64)
	buses := strings.Split(data[1], ",")
	part1(time, buses)
	//part2(buses)
}

func part1(time float64, buses []string) {
	var trips []trip
	for _, bus := range buses {
		if bus != "x" {
			b, _ := strconv.ParseFloat(string(bus), 64)
			t := math.Round(time/b) * b
			if t < time {
				t += b
			}
			trips = append(trips, trip{busID: b, time: t})
		}
	}

	minDiffTrip := findMinDiffTrip(trips, time)
	waitTime := int(minDiffTrip.time - time)
	result := int(minDiffTrip.busID) * waitTime

	fmt.Printf("%d, %d, %d, %d\n", int(minDiffTrip.busID), int(minDiffTrip.time), waitTime, result)
}

func findMinDiffTrip(trips []trip, time float64) trip {
	minDiff := math.Abs(trips[0].time - time)
	minDiffTrip := trips[0]
	for _, trip := range trips {
		diff := math.Abs(trip.time - time)
		if diff < minDiff {
			minDiff = diff
			minDiffTrip = trip
		}
	}
	return minDiffTrip
}
