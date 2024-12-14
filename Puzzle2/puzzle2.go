package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func totalSafe(inputfile string, part2 bool) int {
	// Read the input file and return the each list of numbers as a sorted slice

	safeCount := 0

	// Open the file
	file, err := os.Open(inputfile)

	if err != nil {
		fmt.Println("Error reading the file")
		os.Exit(1)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		level := make([]int, 0)

		tmpLevel := strings.Split(scanner.Text(), " ")

		for i := 0; i < len(tmpLevel); i++ {
			levelItem, err := strconv.Atoi(tmpLevel[i])

			if err != nil {
				fmt.Println("Error converting string to int")
				os.Exit(1)
			}

			level = append(level, levelItem)
		}

		if isSafe(level) {
			safeCount++

		} else {
			if part2 {

				tmpLevel := copySlice(level)
				for i := 0; i < len(level); i++ {

					tmpLevel2 := removeAtIndex(tmpLevel, i)

					if isSafe(tmpLevel2) {
						safeCount++
						break
					}
				}
			}
		}

	}

	return safeCount
}

func removeAtIndex(slice []int, index int) []int {
	newSlice := copySlice(slice)
	return append(newSlice[:index], newSlice[index+1:]...)
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func isSafe(level []int) bool {

	currentItem := level[0]
	delta := 0

	asc := false
	desc := false

	if level[0] < level[1] {
		asc = true
	} else {
		desc = true
	}

	for i := 1; i < len(level); i++ {

		if currentItem <= level[i] {

			if desc {
				return false
			}

			delta = level[i] - currentItem

		}

		if currentItem >= level[i] {

			if asc {
				return false
			}

			delta = currentItem - level[i]

		}

		if delta == 0 || delta > 3 {
			return false
		}

		currentItem = level[i]
	}

	return true
}

func main() {

	// I dont actually need to run this twice. I could just return both answers in one pass.
	fmt.Printf("Total Safe Levels: %d\n", totalSafe("input.txt", false))
	fmt.Printf("Total Safe Dampened Levels: %d\n", totalSafe("input.txt", true))
}
