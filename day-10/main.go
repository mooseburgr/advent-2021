package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	parenL     = "("
	parenR     = ")"
	bracketL   = "["
	bracketR   = "]"
	braceL     = "{"
	braceR     = "}"
	chevL      = "<"
	chevR      = ">"
	parenPts   = 3
	bracketPts = 57
	bracePts   = 1197
	chevPts    = 25137
)

func main() {
	start := time.Now()
	file, err := os.Open("./day-10/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalPts := 0
	for scanner.Scan() {

		parenBalance := 0
		bracketBalance := 0
		braceBalance := 0
		chevBalance := 0

		line := scanner.Text()
		for _, c := range line {
			c := string(c)
			switch c {
			case parenL:
				parenBalance--
			case parenR:
				parenBalance++
			case bracketL:
				bracketBalance--
			case bracketR:
				bracketBalance++
			case braceL:
				braceBalance--
			case braceR:
				braceBalance++
			case chevL:
				chevBalance--
			case chevR:
				chevBalance++
			default:
				panic("ruh roh, what is this?? " + c)
			}

			_, pts := findInvalidCharAndPts(parenBalance, bracketBalance, braceBalance, chevBalance)
			if pts != 0 {
				totalPts += pts
				println(line)
				break
			}
		}
	}

	println("total points:", totalPts)

	fmt.Println("finished after", time.Since(start))
}

func findInvalidCharAndPts(parenBalance int, bracketBalance int, braceBalance int, chevBalance int) (string, int) {
	if parenBalance > 0 {
		return parenR, parenPts
	}
	if bracketBalance > 0 {
		return bracketR, bracketPts
	}
	if braceBalance > 0 {
		return braceR, bracePts
	}
	if chevBalance > 0 {
		return chevR, chevPts
	}

	return "", 0
}
