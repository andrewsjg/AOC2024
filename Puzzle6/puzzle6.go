package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	x         int
	y         int
	direction string
}

type Path map[Position]struct{}

func readInput(inputFile string) map[int][]rune {
	// Read the file line by line and return the data as a slice
	guardMap := make(map[int][]rune)

	// Open the file

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		guardMap[lineCount] = []rune(line)

		lineCount++
	}

	return guardMap

}

func findStart(guardMap map[int][]rune) Position {
	// Find the starting position
	startPos := Position{0, 0, "U"}

	for lineNo, line := range guardMap {
		if strings.Contains(string(line), "^") {
			guardPos := strings.Index(string(line), "^")
			startPos = Position{guardPos, lineNo, "U"}
			break
		}
	}

	return startPos
}

func getPath(guardMap map[int][]rune, startPos Position) Path {

	visited := Path{}

	move := true
	currentPosition := startPos

	visited[currentPosition] = struct{}{}

	// Move the guard
	for move == true {
		move, currentPosition = doMove(currentPosition, guardMap)

		if move {
			visited[currentPosition] = struct{}{}
		}

	}

	return visited
}

func countUniquePositions(path Path) int {

	coordsList := make(map[string]struct{})

	for pos := range path {

		coords := fmt.Sprintf("%d,%d", pos.x, pos.y)
		coordsList[coords] = struct{}{}
	}

	return len(coordsList)
}

func pos2string(pos Position) string {
	return fmt.Sprintf("%d,%d", pos.x, pos.y)
}

// Move the guard. With hindsight this could be a lot simpler!
func doMove(currentPosition Position, guardMap map[int][]rune) (bool, Position) {

	// Move the guard
	canMove := true

	switch currentPosition.direction {
	case "U":
		// Check if the position above is blocked if not update current position to the position above
		if currentPosition.y-1 >= 0 {

			if guardMap[currentPosition.y-1][currentPosition.x] != '#' {
				currentPosition = Position{currentPosition.x, currentPosition.y - 1, "U"}

			} else {
				// Move right

				currentPosition = Position{currentPosition.x, currentPosition.y, "R"}

			}

		} else {
			// gone off the map
			canMove = false
		}

	case "D":
		// Move down
		if currentPosition.y+1 < len(guardMap) {
			if guardMap[currentPosition.y+1][currentPosition.x] != '#' {
				currentPosition = Position{currentPosition.x, currentPosition.y + 1, "D"}

			} else {
				// Move left

				currentPosition = Position{currentPosition.x, currentPosition.y, "L"}

			}
		} else {
			// gone off the map
			canMove = false
		}

	case "L":
		// Move left
		if currentPosition.x-1 >= 0 {
			if guardMap[currentPosition.y][currentPosition.x-1] != '#' {
				currentPosition = Position{currentPosition.x - 1, currentPosition.y, "L"}

			} else {
				// Move up

				currentPosition = Position{currentPosition.x, currentPosition.y, "U"}

			}
		} else {
			// gone off the map
			canMove = false
		}

	case "R":
		// Move right
		if currentPosition.x+1 < len(guardMap[currentPosition.y]) {
			if guardMap[currentPosition.y][currentPosition.x+1] != '#' {
				currentPosition = Position{currentPosition.x + 1, currentPosition.y, "R"}

			} else {
				// Move down

				currentPosition = Position{currentPosition.x, currentPosition.y, "D"}

			}
		} else {
			// gone off the map
			canMove = false
		}
	}

	return canMove, currentPosition
}

func copyMap(original map[int][]rune) map[int][]rune {
	newMap := make(map[int][]rune)
	for key, value := range original {
		newValue := make([]rune, len(value))
		copy(newValue, value)
		newMap[key] = newValue
	}
	return newMap
}

func printMap(guardMap map[int][]rune) {

	for i := 0; i < len(guardMap); i++ {
		fmt.Println(string(guardMap[i]))
	}

	fmt.Println()
}

// Insert a new obstacle into the map at the given position
func insertObstacle(guardMap map[int][]rune, pos Position) map[int][]rune {
	newMap := copyMap(guardMap)

	startPosition := findStart(guardMap)

	if pos == startPosition {
		return newMap
	}

	line := newMap[pos.y]
	line[pos.x] = '#'
	newMap[pos.y] = line

	return newMap
}

// Test if a given map has a loop
func testForLoop(guardMap map[int][]rune, startPos Position) bool {

	newPos := startPos

	visitedPositions := map[Position]struct{}{}

	move := true

	for move {

		move, newPos = doMove(newPos, guardMap)

		if !move {
			return false
		}

		if _, ok := visitedPositions[newPos]; ok {

			return true
		}

		visitedPositions[newPos] = struct{}{}

	}

	return false
}

func part1(inputfile string) int {
	totalPositions := 0
	guardMap := readInput(inputfile)

	startPos := findStart(guardMap)

	// Generate the path

	path := getPath(guardMap, startPos)
	totalPositions = countUniquePositions(path)

	return totalPositions
}

func part2(inputfile string) int {
	guardMap := readInput(inputfile)
	loops := 0

	startPos := findStart(guardMap)

	// Get the unobstructed path
	path := getPath(guardMap, startPos)

	posMap := make(map[string]struct{})

	for pos := range path {

		posString := pos2string(pos)

		if _, ok := posMap[posString]; ok {
			// Do nothing at the moment
			// skip the position

		} else {
			testMap := insertObstacle(guardMap, pos)

			if testForLoop(testMap, startPos) {
				loops++

			}
			posMap[posString] = struct{}{}
		}

	}

	return loops
}

func main() {

	inputFile := "input.txt"

	fmt.Printf("Part 1 Total: %d\n\n", part1(inputFile))
	fmt.Printf("Part 2 Total: %v\n", part2(inputFile))
}
