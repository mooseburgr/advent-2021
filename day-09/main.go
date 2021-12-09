package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
)

func main() {
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
	part1()

	// part 2
	var basins []Basin
	for r := 0; r < len(heightMap); r++ {
		for c := 0; c < len(heightMap[r]); c++ {
			// anything < 9 is in a basin
			if heightMap[r][c] < 9 {
				addOrAppendToBasins(&basins, Location{
					Row:    r,
					Col:    c,
					Height: heightMap[r][c],
				})
			}
		}
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
}

func addOrAppendToBasins(basins *[]Basin, location Location) {
	neighbz := getCardinalNeighbors(location.Row, location.Col)
	var adjacentBasin *Basin
	// check if location is adjacent to existing basin
	for i, basin := range *basins {
		if containsAny(basin.Locations, neighbz) {
			adjacentBasin = &(*basins)[i]
			adjacentBasin.Locations = append(adjacentBasin.Locations, location)
		}
	}
	// new basin! uh huh hunny
	if adjacentBasin == nil {
		newBasin := Basin{Id: len(*basins)}
		newBasin.Locations = append(newBasin.Locations, location)
		*basins = append(*basins, newBasin)
	}
}

func part1() {
	var localMins []Location
	for r := 0; r < len(heightMap); r++ {
		for c := 0; c < len(heightMap[r]); c++ {
			neighbz := getCardinalNeighbors(r, c)
			if isMinAmongNeighbz(heightMap[r][c], neighbz) {
				localMins = append(localMins, Location{
					Row:  r,
					Col:  c,
					Risk: heightMap[r][c] + 1,
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
	//fmt.Printf("neighbz for r%d c%d are %v\n", r, c, neighbz)
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
