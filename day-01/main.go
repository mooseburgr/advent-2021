package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./day-01/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var measurements []int

	for scanner.Scan() {
		//fmt.Println("read:", scanner.Text())
		depth, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, depth)
	}

	lenMeasurements := len(measurements)
	var rollingSums []int
	prev := -1
	increases := -1
	for i, depth := range measurements {
		if depth > prev {
			increases++
		}
		prev = depth

		if i < lenMeasurements-2 {
			rollingSums = append(rollingSums, measurements[i]+measurements[i+1]+measurements[i+2])
		}
	}

	fmt.Println("part 1: # increases =", increases) // 1832
	fmt.Println("rolling sums length =", len(rollingSums))

	prev = -1
	increases = -1
	for _, sum := range rollingSums {
		if sum > prev {
			increases++
		}
		prev = sum
	}
	fmt.Println("part 2: # increases =", increases) // 1858
}
