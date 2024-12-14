package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readinput(inputfile string) (left []int, right []int) {
	// Read the input file and return the each list of numbers as a sorted slice

	// Open the file
	file, err := os.Open(inputfile)

	if err != nil {
		fmt.Println("Error reading the file")
		os.Exit(1)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")

		leftNum, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("Error converting string to int")
			os.Exit(1)
		}

		rightNum, err := strconv.Atoi(line[len(line)-1])

		if err != nil {
			fmt.Println("Error converting string to int")
			os.Exit(1)
		}

		left = insertSorted(left, leftNum)
		right = insertSorted(right, rightNum)

	}
	return left, right
}

func insertSorted(slice []int, num int) []int {
	// Find the correct position to insert the number
	i := 0
	for i < len(slice) && slice[i] < num {
		i++
	}

	// Insert the number at the found position
	slice = append(slice[:i], append([]int{num}, slice[i:]...)...)
	return slice
}

func calcDistance(left []int, right []int) int {
	// Calculate the distance between the two lists

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		dist := left[i] - right[i]

		totalDistance += int(math.Abs(float64(dist)))
	}

	return totalDistance
}

func similarityScore(leftList []int, rightList []int) int {

	similarityScore := 0
	occurenceMap := make(map[int]int)

	for i := 0; i < len(leftList); i++ {
		for j := 0; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				occurenceMap[leftList[i]] += 1
			}
		}
	}

	for key, value := range occurenceMap {
		similarityScore += value * key
	}

	return similarityScore
}

func main() {
	// Read the input file
	left, right := readinput("input.txt")

	fmt.Printf("Total Distance: %d\n", calcDistance(left, right))
	fmt.Printf("Similarity Score: %d\n", similarityScore(left, right))
}
