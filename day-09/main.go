package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Location struct {
	Row    int
	Col    int
	Height int
	Risk   int
}

type Basin struct {
	Id        int
	Locations []Location
}

var (
	heightMap [][]int
	localMins []Location
	basins    []Basin
)

func main() {
	start := time.Now()
	file, err := os.Open("./day-09/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var rowHeights []int
		for _, c := range scanner.Text() {
			height, _ := strconv.Atoi(string(c))
			rowHeights = append(rowHeights, height)
		}
		//fmt.Printf("%v\n", rowHeights)
		heightMap = append(heightMap, rowHeights)
	}
	fmt.Printf("heightMap dimensions: %dr x %dc\n", len(heightMap), len(heightMap[0]))

	// part 1
	findAllLocalMins()

	// part 2
	for _, min := range localMins {
		fmt.Printf("on local min: %v\n", min)
		addOrAppendToBasins(min)
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i].Locations) > len(basins[j].Locations)
	})
	//fmt.Printf("all basins sorted: %v\n", basins)
	fmt.Printf("product of largest three sizes:%d\n\n",
		len(basins[0].Locations)*len(basins[1].Locations)*len(basins[2].Locations))

	for r := 0; r < len(heightMap); r++ {
		for c := 0; c < len(heightMap[r]); c++ {
			//printColor(r, c, basins)
			fmt.Print(string("\033[3"+strconv.Itoa(getContainerBasinId(r, c, basins)%7+1)+"m"), heightMap[r][c])
		}
		fmt.Println("", string("\033[0m"))
	}
	fmt.Println("finished after", time.Since(start))
}

func addOrAppendToBasins(location Location) {
	neighbz := getCardinalNeighbors(location.Row, location.Col)
	var adjacentBasin *Basin
	// check if location is adjacent to existing basin
	for i, basin := range basins {
		if containsAny(basin.Locations, neighbz) {
			adjacentBasin = &basins[i]
			if !contains(adjacentBasin.Locations, location) {
				// only add if not already tracked
				adjacentBasin.Locations = append(adjacentBasin.Locations, location)
			}
		}
	}

	// new basin! uh huh hunny
	if adjacentBasin == nil {
		newBasin := Basin{Id: len(basins)}
		newBasin.Locations = append(newBasin.Locations, location)
		basins = append(basins, newBasin)
	}

	// recursively expand through rest of basin until hitting walls (9s)
	for _, l := range getCardinalNeighbors(location.Row, location.Col) {
		if l.Height != 9 && !isAlreadyTrackedInBasin(l) {
			addOrAppendToBasins(l)
		}
	}
}

func findAllLocalMins() {
	for r := 0; r < len(heightMap); r++ {
		for c := 0; c < len(heightMap[r]); c++ {
			neighbz := getCardinalNeighbors(r, c)
			if isMinAmongNeighbz(heightMap[r][c], neighbz) {
				localMins = append(localMins, Location{
					Row:    r,
					Col:    c,
					Risk:   heightMap[r][c] + 1,
					Height: heightMap[r][c],
				})
			}
		}
	}
	// fmt.Printf("local maxes:%v\n", localMins)
	totalRisk := 0
	for _, min := range localMins {
		totalRisk += min.Risk
	}
	println("total risk", totalRisk)
}

func isMinAmongNeighbz(height int, neighbz []Location) bool {
	sort.Slice(neighbz, func(i, j int) bool {
		return neighbz[i].Height < neighbz[j].Height
	})
	result := height < neighbz[0].Height
	if result {
		//fmt.Printf("%d minimum next to %v: %t\n", height, neighbz, result)
	}
	return result
}

func getCardinalNeighbors(r int, c int) []Location {
	var neighbz []Location
	// north
	if r > 0 {
		neighbz = append(neighbz,
			Location{
				Row:    r - 1,
				Col:    c,
				Height: heightMap[r-1][c],
				Risk:   1 + heightMap[r-1][c],
			})
	}
	// west
	if c > 0 {
		neighbz = append(neighbz,
			Location{
				Row:    r,
				Col:    c - 1,
				Height: heightMap[r][c-1],
				Risk:   1 + heightMap[r][c-1],
			})
	}
	neighbz = append(neighbz, getSouthAndEastNeighbors(r, c)...)
	return neighbz
}

func getSouthAndEastNeighbors(r int, c int) []Location {
	var neighbz []Location
	// east
	if c < len(heightMap[r])-1 {
		neighbz = append(neighbz,
			Location{
				Row:    r,
				Col:    c + 1,
				Height: heightMap[r][c+1],
				Risk:   1 + heightMap[r][c+1],
			})
	}
	// south
	if r < len(heightMap)-1 {
		neighbz = append(neighbz,
			Location{
				Row:    r + 1,
				Col:    c,
				Height: heightMap[r+1][c],
				Risk:   1 + heightMap[r+1][c],
			})
	}
	return neighbz
}

func getContainerBasinId(r int, c int, basins []Basin) int {
	for _, basin := range basins {
		if contains(basin.Locations, Location{Row: r, Col: c}) {
			return basin.Id
		}
	}
	return -1
}

func isAlreadyTrackedInBasin(loc Location) bool {
	for _, basin := range basins {
		if contains(basin.Locations, loc) {
			return true
		}
	}
	return false
}

func (l *Location) Equals(other Location) bool {
	return l.Row == other.Row && l.Col == other.Col
}

func contains(s []Location, t Location) bool {
	for _, l := range s {
		if l.Equals(t) {
			return true
		}
	}
	return false
}

// containsAny checks if any of test locations `t` are contained within source locations `s`
func containsAny(s []Location, t []Location) bool {
	for _, l := range t {
		if contains(s, l) {
			return true
		}
	}
	return false
}
