package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TopoMap map[int][]int

type Point struct {
	X, Y int
}

func (tp *TopoMap) Print() {

	topoVals := ""
	for i := 0; i < len(*tp); i++ {

		row := (*tp)[i]
		for _, item := range row {
			topoVals = topoVals + strconv.Itoa(item)
		}

		topoVals = topoVals + "\n"
	}

	fmt.Println(strings.TrimSpace(topoVals))
}

func (tp *TopoMap) GetValueAtPoint(p Point) int {

	return (*tp)[p.Y][p.X]
}

func (tp *TopoMap) GetTrailHeads() []Point {

	var trailHeads []Point

	for y := 0; y < len(*tp); y++ {

		for x, topoVal := range (*tp)[y] {

			if topoVal == 0 {
				trailHeads = append(trailHeads, Point{x, y})
			}
		}
	}

	return trailHeads
}

func (tp *TopoMap) MoveLeft(p Point) Point {

	if p.X-1 >= 0 && p.X-1 < len((*tp)[p.Y]) {
		return Point{p.X - 1, p.Y}
	}

	return Point{-1, -1}
}

func (tp *TopoMap) MoveRight(p Point) Point {

	if p.X+1 >= 0 && p.X+1 < len((*tp)[p.Y]) {
		return Point{p.X + 1, p.Y}
	}

	return Point{-1, -1}
}

func (tp *TopoMap) MoveUp(p Point) Point {

	if p.Y-1 >= 0 && p.Y-1 < len(*tp) {
		return Point{p.X, p.Y - 1}
	}

	return Point{-1, -1}
}

func (tp *TopoMap) MoveDown(p Point) Point {

	if p.Y+1 >= 0 && p.Y+1 < len(*tp) {
		return Point{p.X, p.Y + 1}
	}

	return Point{-1, -1}
}

// This will follow a trail if the path remains within the bounds of the map
// and the next value on the path is no more than 1 more than the previous.
// If there is no valid path, it will return a point with coordinates -1, -1

func (tp *TopoMap) TopoMove(p Point, direction rune) Point {

	switch direction {
	case 'U':
		if p.Y-1 >= 0 && p.Y-1 < len(*tp) {
			if tp.GetValueAtPoint(Point{p.X, p.Y - 1}) == tp.GetValueAtPoint(p)+1 {
				return Point{p.X, p.Y - 1}
			}
		}
	case 'D':
		if p.Y+1 >= 0 && p.Y+1 < len(*tp) {
			if tp.GetValueAtPoint(Point{p.X, p.Y + 1}) == tp.GetValueAtPoint(p)+1 {
				return Point{p.X, p.Y + 1}
			}
		}
	case 'L':
		if p.X-1 >= 0 && p.X-1 < len((*tp)[p.Y]) {
			if tp.GetValueAtPoint(Point{p.X - 1, p.Y}) == tp.GetValueAtPoint(p)+1 {
				return Point{p.X - 1, p.Y}
			}
		}
	case 'R':

		if p.X+1 >= 0 && p.X+1 < len((*tp)[p.Y]) {
			if tp.GetValueAtPoint(Point{p.X + 1, p.Y}) == tp.GetValueAtPoint(p)+1 {
				return Point{p.X + 1, p.Y}
			}
		}
	}

	return Point{-1, -1}
}

func strToIntArray(s string) []int {
	var intArray []int
	for _, r := range s {
		intArray = append(intArray, runeToInt(r))
	}
	return intArray
}

func runeToInt(r rune) int {
	return int(r - '0')
}

func getMap(inputFile string) TopoMap {

	topMap := make(TopoMap)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		topMap[y] = strToIntArray(scanner.Text())

		y++
	}

	return topMap
}

func followTrail(tp TopoMap, trailStart Point, exits map[Point]bool, part2 bool) int {

	score := 0

	if tp.GetValueAtPoint(trailStart) == 9 {

		// For part 2 we actually want to count all the ways to reach an exit, not just the exits
		// So count every exit even if we've already seen it.
		if part2 {
			return 1
		}

		if _, exists := exits[trailStart]; !exists {
			exits[trailStart] = true
			return 1
		}

		return 0

	}

	if tp.TopoMove(trailStart, 'U').X != -1 {

		score += followTrail(tp, tp.TopoMove(trailStart, 'U'), exits, part2)
	}

	if tp.TopoMove(trailStart, 'D').X != -1 {

		score += followTrail(tp, tp.TopoMove(trailStart, 'D'), exits, part2)
	}

	if tp.TopoMove(trailStart, 'L').X != -1 {

		score += followTrail(tp, tp.TopoMove(trailStart, 'L'), exits, part2)
	}

	if tp.TopoMove(trailStart, 'R').X != -1 {

		score += followTrail(tp, tp.TopoMove(trailStart, 'R'), exits, part2)
	}

	return score
}

func main() {
	inputFile := "input.txt"
	topoMap := getMap(inputFile)

	th := topoMap.GetTrailHeads()
	score := 0

	for _, trailHead := range th {
		exits := make(map[Point]bool)
		score += followTrail(topoMap, trailHead, exits, false)
	}

	fmt.Println("Part 1 Score:", score)

	score = 0
	for _, trailHead := range th {
		exits := make(map[Point]bool)
		score += followTrail(topoMap, trailHead, exits, true)
	}

	fmt.Println("Part 2 Score:", score)
}
