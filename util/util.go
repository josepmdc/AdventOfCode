package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetData(fileName string) []int {
	var data []int

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return data
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err == nil {
			data = append(data, val)
		}
	}
	return data
}

func GetDataString(fileName string) []string {
	var data []string

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return data
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
