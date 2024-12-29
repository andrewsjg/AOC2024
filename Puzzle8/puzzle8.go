package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Point struct {
	X, Y int
}

type Node struct {
	loc      Point
	NodeChar rune
}

type NodeList []Node

// Return a map structure and the x and y dimensions of the map
func getMap(inputFile string) (map[Point]rune, int, int) {
	// Read the file line by line and return the data as a map
	antennaMap := make(map[Point]rune)

	xDim := 0
	yDim := 0

	// Open the file

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		if xDim == 0 {
			xDim = len(line)
		}

		for x, char := range line {
			antennaMap[Point{x, y}] = char
		}

		y++
	}

	yDim = y

	return antennaMap, xDim, yDim

}

func getNodeList(antennaMap map[Point]rune) NodeList {
	nodeList := make(NodeList, 0)

	for loc, char := range antennaMap {
		if char != '.' {
			nodeList = append(nodeList, Node{loc, char})
		}
	}

	return nodeList
}

func getAntennaLocs(antennaMap map[Point]rune) map[rune][]Point {
	antennaLocs := make(map[rune][]Point)

	for loc, char := range antennaMap {
		if char != '.' {
			antennaLocs[char] = append(antennaLocs[char], loc)
		}
	}

	return antennaLocs
}

func copySlice(slice []Point) []Point {
	newSlice := make([]Point, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// Print a map of the given X and Y dimensions
func printMap(data map[Point]rune, xDim int, yDim int) {
	// Print the map
	for y := 0; y < yDim; y++ {
		for x := 0; x < xDim; x++ {
			fmt.Printf("%c", data[Point{x, y}])
		}
		fmt.Println()
	}
}

func getAntiNodes(nodes map[rune][]Point, xDim int, yDim int, part2 bool, debug bool) map[Point]struct{} {
	antiNodes := make(map[Point]struct{})

	for node, loc := range nodes {

		if debug {
			fmt.Printf("Node Type: %c, Locations: %v\n\n", node, loc)
		}

		tmpLocs := copySlice(loc)

		for len(tmpLocs) > 1 {

			testNode := tmpLocs[0]
			tmpLocs = tmpLocs[1:]

			for _, sameNode := range tmpLocs {

				if debug {
					XDist, YDist := XYDistance(testNode, sameNode)
					fmt.Printf("Finding Antinodes for nodes: %v, and: %v XYDistance of: %d,%d Manhattan distance of: %d \n", testNode, sameNode, XDist, YDist, manhattanDistance(testNode, sameNode))
				}

				an1 := Point{}
				an2 := Point{}

				xd, yd := XYDistance(testNode, sameNode)

				an1 = Point{testNode.X + xd, testNode.Y + yd}
				an2 = Point{sameNode.X - xd, sameNode.Y - yd}

				if an1.X >= 0 && an1.X < xDim && an1.Y >= 0 && an1.Y < yDim {

					antiNodes[an1] = struct{}{}
				}

				if an2.X >= 0 && an2.X < xDim && an2.Y >= 0 && an2.Y < yDim {
					antiNodes[an2] = struct{}{}
				}

				if part2 {

					antiNodes[testNode] = struct{}{}
					antiNodes[sameNode] = struct{}{}

					for an1.X >= 0 && an1.X < xDim && an1.Y >= 0 && an1.Y < yDim {

						antiNodes[an1] = struct{}{}

						an1.X += xd
						an1.Y += yd

					}

					for an2.X >= 0 && an2.X < xDim && an2.Y >= 0 && an2.Y < yDim {

						antiNodes[an2] = struct{}{}

						an2.X -= xd
						an2.Y -= yd

					}

				}

			}
		}

		if debug {
			fmt.Println()
		}
	}

	return antiNodes
}

func manhattanDistance(p1, p2 Point) int {
	return int(math.Abs(float64(p2.X-p1.X)) + math.Abs(float64(p2.Y-p1.Y)))
}

func XYDistance(p1, p2 Point) (int, int) {
	return p1.X - p2.X, p1.Y - p2.Y
}
func main() {

	antennaMap, xDim, yDim := getMap("input.txt")
	antennaLocs := getAntennaLocs(antennaMap)

	p1antiNodes := getAntiNodes(antennaLocs, xDim, yDim, false, false)
	p2antiNodes := getAntiNodes(antennaLocs, xDim, yDim, true, false)

	fmt.Printf("Total Antinodes for Part 1: %d\nTotal Antinodes ofr Part 2: %d\n", len(p1antiNodes), len(p2antiNodes))

}
