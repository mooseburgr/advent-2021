package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	X int
	Y int
}

func (v1 Vector) add(v2 Vector) Vector {
	return Vector{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func main() {

	file, err := os.Open("./day-02/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	position := Vector{X: 0, Y: 0}
	aim := 0
	for scanner.Scan() {
		handleCommand(scanner.Text(), &position, &aim)
	}
	println("final position:", position)
	println("product:", position.X*position.Y)
}

func handleCommand(input string, position *Vector, aim *int) {
	split := strings.Split(input, " ")
	distance, _ := strconv.Atoi(split[1])

	switch split[0] {
	case "forward":
		position.X = position.X + distance
		position.Y = position.Y + distance*(*aim)

	case "up":
		*aim -= distance
	case "down":
		*aim += distance
	}
}
