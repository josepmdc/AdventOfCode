package main

import "fmt"

func main() {
	part1()
}

func part1() {
	cache := map[int]int{11: 1, 18: 2, 0: 3, 20: 4, 1: 5, 7: 6}
	last := 16
	n := 0
	for i := 8; i <= 2020; i++ {
		if cache[last] != 0 {
			n = i - 1 - cache[last]
		} else {
			n = 0
		}
		cache[last] = i - 1
		last = n
	}
	fmt.Printf("Part 1: %d\n", n)
}
