package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	segmentsIn0 = 6
	segmentsIn1 = 2
	segmentsIn2 = 5
	segmentsIn3 = 5
	segmentsIn4 = 4
	segmentsIn5 = 5
	segmentsIn6 = 6
	segmentsIn7 = 3
	segmentsIn8 = 7
	segmentsIn9 = 6
)

func main() {

	file, err := os.Open("./day-08/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total1s := 0
	total4s := 0
	total7s := 0
	total8s := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "|")

		patterns := strings.Split(strings.TrimSpace(split[0]), " ")
		outputs := strings.Split(strings.TrimSpace(split[1]), " ")
		fmt.Printf("\npatterns: %v\n outputs: %s\n", patterns, outputs)

		num1s, num4s, num7s, num8s := countOccurrences1478(outputs)

		total1s += num1s
		total4s += num4s
		total7s += num7s
		total8s += num8s
	}

	// part 1
	fmt.Printf("total 1s: %d \ntotal 4s: %d \ntotal 7s: %d \ntotal 8s: %d \n",
		total1s, total4s, total7s, total8s)

}

func countOccurrences1478(outputs []string) (int, int, int, int) {
	num1s := 0
	num4s := 0
	num7s := 0
	num8s := 0
	for _, str := range outputs {
		switch len(str) {
		case segmentsIn1:
			num1s++
		case segmentsIn4:
			num4s++
		case segmentsIn7:
			num7s++
		case segmentsIn8:
			num8s++
		}
	}
	return num1s, num4s, num7s, num8s
}
