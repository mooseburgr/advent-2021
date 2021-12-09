package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"sync"
)

type Result struct {
	TargetPos int
	Fuel      int
}

var (
	initialPositions []int

	results []Result

	mutex = sync.RWMutex{}
)

func main() {

	file, err := os.Open("./day-07/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, _ := csvReader.Read()

	for _, val := range data {
		pos, _ := strconv.Atoi(val)
		initialPositions = append(initialPositions, pos)
	}
	sort.Ints(initialPositions)

	minTarget := initialPositions[0]
	maxTarget := initialPositions[len(initialPositions)-1]

	fmt.Printf("target somewhere between %d and %d\n", minTarget, maxTarget)

	for pos := minTarget; pos <= maxTarget; pos++ {
		// go 	lol, meh
		calculateTotalFuel(pos)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Fuel < results[j].Fuel
	})
	println("most efficient result:", results[0])

	sort.Slice(results, func(i, j int) bool {
		return results[i].Fuel > results[j].Fuel
	})
	println("least efficient result:", results[0])
}

func calculateTotalFuel(targetPosition int) {
	totalFuel := 0
	for _, initialPos := range initialPositions {
		totalFuel += calcSingleTripFuel(int(math.Abs(float64(targetPosition - initialPos))))
	}
	mutex.Lock()
	results = append(results, Result{TargetPos: targetPosition, Fuel: totalFuel})
	mutex.Unlock()
}

func calcSingleTripFuel(tripLength int) int {
	// in part 2, single trip fuel cost = 1 + 2 + 3 + ... + n
	return tripLength * (tripLength + 1) / 2
}
