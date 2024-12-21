package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(inputFile string) (orderRules map[int][]int, pages [][]int) {
	// Read the file line by line and return the data as a map
	orderRules = make(map[int][]int)
	pages = make([][]int, 0)

	readRules := true

	// Open the file

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readRules = false
		}

		if readRules {
			// Split the line
			page, order := splitLine(line)
			orderRules[page] = append(orderRules[page], order)
		} else {

			if line != "" {
				strArr := strings.Split(line, ",")
				pages = append(pages, convertToIntArr(strArr))
			}
		}

	}

	return orderRules, pages

}

func convertToIntArr(strArr []string) []int {
	// Convert the string array to an int array
	intArr := make([]int, 0)

	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		intArr = append(intArr, num)
	}

	return intArr
}

func splitLine(line string) (int, int) {
	// Split the line into the two parts

	// Split the line
	parts := strings.Split(line, "|")

	if len(parts) != 2 {
		log.Fatalf("Error splitting the line: %v", line)
	}

	page, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	order, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	return page, order
}

func contains(arr []int, num int) bool {
	// Check if the number is in the array
	for _, n := range arr {
		if n == num {
			return true
		}
	}

	return false
}

func getMiddleNumber(arr []int) int {
	// Get the middle number in the array
	if len(arr) == 0 {
		return 0
	}

	if len(arr)%2 == 0 {
		return arr[len(arr)/2]
	}

	return arr[len(arr)/2]
}

func removeAndAppend(arr []int, num int) []int {
	// Find the index of the number to be removed
	index := -1
	for i, v := range arr {
		if v == num {
			index = i
			break
		}
	}

	// If the number is not found, return the original array
	if index == -1 {
		return arr
	}

	// Remove the number from the array
	arr = append(arr[:index], arr[index+1:]...)

	// Append the number to the end of the array
	arr = append(arr, num)

	return arr
}

func makeValid(arr []int, orderRules map[int][]int) []int {

	pagesSeen := make([]int, 0)

	for _, page := range arr {

		rules := orderRules[page]
		pagesSeen = append(pagesSeen, page)

		for _, rule := range rules {
			if contains(pagesSeen, rule) {

				// Move the invalid rule number to the end of the array
				newArr := removeAndAppend(arr, rule)

				// recursively call makeValid with new array
				return makeValid(newArr, orderRules)
			}
		}
	}

	// After the recursion unwinds the input arr will be a valid array which falls all the way through to here
	// so simply return the now valid input array
	return arr
}

func solution(inputfile string) (int, int) {
	// Read the input file
	orderRules, pages := readInput(inputfile)
	validUpdate := true
	part1total := 0
	part2total := 0

	for _, update := range pages {
		validUpdate = true
		pagesSeen := make([]int, 0)

		for _, page := range update {
			pagesSeen = append(pagesSeen, page)

			// Get the order rules for the page
			rules := orderRules[page]

			for _, rule := range rules {
				// Check if the page is already in the pages seen. If it is, then the update is invalid
				// since the page would have been seen before the current one which is forbidden by the rules
				if contains(pagesSeen, rule) {
					validUpdate = false
				}
			}
		}

		if validUpdate {
			mid := getMiddleNumber(update)
			part1total = part1total + mid

		} else {
			newVaild := makeValid(update, orderRules)
			mid := getMiddleNumber(newVaild)
			part2total = part2total + mid

		}

	}

	return part1total, part2total
}

func main() {
	p1, p2 := solution("input.txt")

	fmt.Printf("Part 1 Total: %d\nPart 2 Total: %d\n", p1, p2)
}
