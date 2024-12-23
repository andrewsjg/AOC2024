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

type Path []Position

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

// Append a string to a slice of strings if it doesn't already exist
func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

// Not sure if I need to use strings here. But it works. Makes the appendIfMissing function easier maybe?
func pos2string(pos Position) string {
	return fmt.Sprintf("%d,%d", pos.x, pos.y)
}

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
				if currentPosition.x+1 < len(guardMap[currentPosition.y]) {
					currentPosition = Position{currentPosition.x + 1, currentPosition.y, "R"}
				} else {
					// gone off the map
					canMove = false
				}
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
				if currentPosition.x-1 >= 0 {
					currentPosition = Position{currentPosition.x - 1, currentPosition.y, "L"}
				} else {
					// gone off the map
					canMove = false
				}
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
				if currentPosition.y-1 >= 0 {
					currentPosition = Position{currentPosition.x, currentPosition.y - 1, "U"}
				} else {
					// gone off the map
					canMove = false
				}
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
				if currentPosition.y+1 < len(guardMap) {
					currentPosition = Position{currentPosition.x, currentPosition.y + 1, "D"}
				} else {
					// gone off the map
					canMove = false
				}
			}
		} else {
			// gone off the map
			canMove = false
		}
	}

	return canMove, currentPosition
}

func part1_old(inputfile string) (int, Path) {
	totalPositions := 0
	guardMap := readInput(inputfile)

	// This was used to solve part 1. Could replace with the path slice instead?
	visitedPositions := make([]string, 0)
	moveDirection := "U"
	path := make(Path, 0)

	currentPosition := Position{0, 0, "U"}

	// Find the Staring position
	for lineNo, line := range guardMap {
		if strings.Contains(string(line), "^") {
			guardPos := strings.Index(string(line), "^")
			currentPosition = Position{guardPos, lineNo, "U"}

			visitedPositions = appendIfMissing(visitedPositions, pos2string(currentPosition))
			path = append(path, currentPosition)
			break
		}
	}

	// Move the guard
	move := true
	for move == true {

		switch moveDirection {
		case "U":
			// Check if the position above is blocked if not update current position to the position above
			if currentPosition.y-1 >= 0 {
				if guardMap[currentPosition.y-1][currentPosition.x] != '#' {
					currentPosition = Position{currentPosition.x, currentPosition.y - 1, "U"}
					path = append(path, currentPosition)

				} else {
					// Move right
					moveDirection = "R"
				}

			} else {
				// gone off the map
				move = false
			}

		case "D":
			// Move down
			if currentPosition.y+1 < len(guardMap) {
				if guardMap[currentPosition.y+1][currentPosition.x] != '#' {
					currentPosition = Position{currentPosition.x, currentPosition.y + 1, "D"}
					path = append(path, currentPosition)

				} else {
					// Move left
					moveDirection = "L"
				}
			} else {
				// gone off the map
				move = false
			}

		case "L":
			// Move left
			if currentPosition.x-1 >= 0 {
				if guardMap[currentPosition.y][currentPosition.x-1] != '#' {
					currentPosition = Position{currentPosition.x - 1, currentPosition.y, "L"}
					path = append(path, currentPosition)

				} else {
					// Move up
					moveDirection = "U"
				}
			} else {
				// gone off the map
				move = false
			}

		case "R":
			// Move right
			if currentPosition.x+1 < len(guardMap[currentPosition.y]) {
				if guardMap[currentPosition.y][currentPosition.x+1] != '#' {
					currentPosition = Position{currentPosition.x + 1, currentPosition.y, "R"}
					path = append(path, currentPosition)

				} else {
					// Move down
					moveDirection = "D"
				}
			} else {
				// gone off the map
				move = false
			}
		}

		visitedPositions = appendIfMissing(visitedPositions, pos2string(currentPosition))

	}

	totalPositions = len(visitedPositions)
	return totalPositions, path
}

func part1(inputfile string) (int, Path) {
	totalPositions := 0
	guardMap := readInput(inputfile)

	// This was used to solve part 1. Could replace with the path slice instead?
	visitedPositions := make([]string, 0)

	path := make(Path, 0)

	currentPosition := Position{0, 0, "U"}

	// Find the Staring position
	for lineNo, line := range guardMap {
		if strings.Contains(string(line), "^") {
			guardPos := strings.Index(string(line), "^")
			currentPosition = Position{guardPos, lineNo, "U"}

			visitedPositions = appendIfMissing(visitedPositions, pos2string(currentPosition))
			path = append(path, currentPosition)
			break
		}
	}

	// Move the guard
	move := true

	for move == true {
		move, currentPosition = doMove(currentPosition, guardMap)
		visitedPositions = appendIfMissing(visitedPositions, pos2string(currentPosition))

		if move {
			path = append(path, currentPosition)
		}
	}

	totalPositions = len(visitedPositions)
	return totalPositions, path
}

func isLoop(testPath Path) bool {

	pathMap := make(map[Position]int)
	//fmt.Println(testPath)

	for _, pos := range testPath {

		// If the position doesnt exist in the map then we havent found a loop
		//fmt.Println("Testiing pos: ", pos)
		if _, ok := pathMap[pos]; !ok {
			pathMap[pos] = 0

		} else {
			fmt.Printf("\nloop at %v\n", pos)
			//fmt.Println(pathMap)
			return true
		}
	}

	return false
}

func copyPath(path []Position) []Position {
	newSlice := make([]Position, len(path))
	copy(newSlice, path)
	return newSlice
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

func part21(path Path, inputFile string) int {

	loops := 0

	guardMap := readInput(inputFile)

	for i, currentPosition := range path {

		newPos := currentPosition
		move := true

		loopTest := copyPath(path)

		switch newPos.direction {
		case "U":
			fmt.Printf("Changing from %v to %s\n", newPos, "R")
			newPos.direction = "R"

		case "D":
			fmt.Printf("Changing from %v to %s\n", newPos, "L")
			newPos.direction = "L"

		case "L":
			fmt.Printf("Changing from %v to %s\n", newPos, "U")
			newPos.direction = "U"

		case "R":
			fmt.Printf("Changing from %v to %s\n", newPos, "D")
			newPos.direction = "D"

		}

		// Plot the new path then do the loop test

		loopTest = loopTest[:i]

		// There is an edge where the guard will oscillate between two positions. Need to detect and make
		// an additional right turn to break the loop
		infititeLoopCheck := map[Position]int{}

		for move == true {

			move, newPos = doMove(newPos, guardMap)

			infititeLoopCheck[newPos] = infititeLoopCheck[newPos] + 1

			if move {
				loopTest = append(loopTest, newPos)
			}

			if infititeLoopCheck[newPos] > 2 {
				//fmt.Println(newPos)

				//fmt.Println(infititeLoopCheck[newPos])
				switch newPos.direction {
				case "U":
					fmt.Printf("Inf loop fix %v to %s\n", newPos, "R")
					newPos.direction = "R"

				case "D":
					fmt.Printf("Inf loop fix %v to %s\n", newPos, "L")
					newPos.direction = "L"

				case "L":
					fmt.Printf("Inf loop fix %v to %s\n", newPos, "U")
					newPos.direction = "U"

				case "R":
					fmt.Printf("Inf loop fix %v to %s\n", newPos, "D")
					newPos.direction = "D"

				}
			}
		}

		if isLoop(loopTest) {
			loops++
		}

	}

	return loops
}

func printMap(guardMap map[int][]rune) {

	for i := 0; i < len(guardMap); i++ {
		fmt.Println(string(guardMap[i]))
	}

	fmt.Println()
}

func part2(path Path, inputFile string) int {

	loops := 0
	startPosition := Position{0, 0, "U"}

	guardMap := readInput(inputFile)

	// Find the Staring position
	for lineNo, line := range guardMap {
		if strings.Contains(string(line), "^") {
			guardPos := strings.Index(string(line), "^")
			startPosition = Position{guardPos, lineNo, "U"}
			break
		}
	}

	fmt.Println("Start Position: ", startPosition)
	loopPos := make([]Position, 0)

	for _, currentPosition := range path {
		tmpGuardMap := copyMap(guardMap)

		switch currentPosition.direction {

		case "U":
			if currentPosition.y-1 >= 0 {
				line := tmpGuardMap[currentPosition.y-1]

				if line[currentPosition.x] != '#' {
					line[currentPosition.x] = '#'
					tmpGuardMap[currentPosition.y-1] = line
				} else {
					// Move right
					if currentPosition.x+1 < len(tmpGuardMap[currentPosition.y]) {
						line := tmpGuardMap[currentPosition.y]
						line[currentPosition.x+1] = '#'
						tmpGuardMap[currentPosition.y] = line
					}
				}
			}

		case "D":
			if currentPosition.y+1 < len(tmpGuardMap) {
				line := tmpGuardMap[currentPosition.y+1]

				if line[currentPosition.x] != '#' {
					line[currentPosition.x] = '#'
					tmpGuardMap[currentPosition.y+1] = line
				} else {
					// Move left
					if currentPosition.x-1 >= 0 {
						line := tmpGuardMap[currentPosition.y]
						line[currentPosition.x-1] = '#'
						tmpGuardMap[currentPosition.y] = line
					}
				}
			}

		case "L":
			if currentPosition.x-1 >= 0 {
				line := tmpGuardMap[currentPosition.y]

				if line[currentPosition.x-1] != '#' {
					line[currentPosition.x-1] = '#'
					tmpGuardMap[currentPosition.y] = line
				} else {
					// Move up
					if currentPosition.y-1 >= 0 {
						line := tmpGuardMap[currentPosition.y-1]
						line[currentPosition.x] = '#'
						tmpGuardMap[currentPosition.y-1] = line
					}
				}
			}

		case "R":
			if currentPosition.x+1 < len(tmpGuardMap[currentPosition.y]) {
				line := tmpGuardMap[currentPosition.y]

				if line[currentPosition.x+1] != '#' {
					line[currentPosition.x+1] = '#'
					tmpGuardMap[currentPosition.y] = line
				} else {
					// Move down
					if currentPosition.y+1 < len(tmpGuardMap) {
						line := tmpGuardMap[currentPosition.y+1]
						line[currentPosition.x] = '#'
						tmpGuardMap[currentPosition.y+1] = line
					}
				}
			}
		}

		printMap(tmpGuardMap)
		newPos := startPosition
		testPath := make(Path, 0)
		//visitedPositions := make([]Position, 0)

		move := true
		for move == true {
			move, newPos = doMove(newPos, tmpGuardMap)

			if move {
				testPath = append(testPath, newPos)
			}

		}

		if isLoop(testPath) {
			loops++

		}

	}

	fmt.Println(loopPos)
	return loops
}

func isLoop2(testpath Path) bool {
	visited := make(map[Position]bool)
	for _, pos := range testpath {
		if visited[pos] {
			return true
		}
		visited[pos] = true
	}
	return false
}

func Contains(s []Position, e Position) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	inputFile := "sample_input.txt"
	p1Total, _ := part1(inputFile)

	fmt.Printf("Part 1 Total: %d\n\n", p1Total)
	//fmt.Printf("Part 2 Total: %v\n", part2(path, inputFile))
}

// 1908 too low
// 1910 not right
// 1933
