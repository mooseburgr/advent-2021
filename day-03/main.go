package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	bitArray [][]uint8
)

func main() {
	start := time.Now()
	file, err := os.Open("./day-03/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var row []uint8
		for _, c := range scanner.Text() {
			integer, _ := strconv.Atoi(string(c))
			row = append(row, uint8(integer))
		}
		bitArray = append(bitArray, row)
	}

	// part 1. assuming all rows are same length
	var columnBits []ColumnBit
	for col := 0; col < len(bitArray[0]); col++ {
		colBit := ColumnBit{
			index:     col,
			zeroVotes: 0,
			oneVotes:  0,
		}
		for row := 0; row < len(bitArray); row++ {
			if bitArray[row][col] == 1 {
				colBit.oneVotes++
			} else {
				colBit.zeroVotes++
			}
		}
		columnBits = append(columnBits, colBit)
	}
	gammaDec, _ := strconv.ParseInt(GetGammaBinaryStr(columnBits), 2, 64)
	epsilonDec, _ := strconv.ParseInt(GetEpsilonBinaryStr(columnBits), 2, 64)

	fmt.Printf("gamma in decimal: %d \nepsilon in decimal: %d\n", gammaDec, epsilonDec)

	fmt.Println("finished after", time.Since(start))
}

type ColumnBit struct {
	index     int
	zeroVotes int
	oneVotes  int
}

func GetGammaBinaryStr(colBits []ColumnBit) string {
	var chars []string
	for _, cb := range colBits {
		chars = append(chars, cb.GetGammaBit())
	}
	return strings.Join(chars, "")
}

func GetEpsilonBinaryStr(colBits []ColumnBit) string {
	var chars []string
	for _, cb := range colBits {
		chars = append(chars, cb.GetEpsilonBit())
	}
	return strings.Join(chars, "")
}

func (c ColumnBit) GetGammaBit() string {
	if c.zeroVotes > c.oneVotes {
		return "0"
	} else {
		return "1"
	}
}

func (c ColumnBit) GetEpsilonBit() string {
	if c.zeroVotes > c.oneVotes {
		return "1"
	} else {
		return "0"
	}
}
