package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getMatches(content []byte, pattern string) []string {

	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}

	// Find all matches
	matches := re.FindAllString(string(content), -1)

	return matches
}

func readInput(inputFile string) []byte {
	// Read the file data
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return data
}

func processInstructions(instructions []string) int {
	// Process the matches
	total := 0
	for _, instruction := range instructions {
		if strings.Contains(instruction, "mul") {
			total = total + doMultiplication(instruction)
		}
	}
	return total
}

func doMultiplication(instruction string) int {
	// Extract the numbers from the instruction
	nums := strings.Split(instruction, ",")
	num1 := getNumber(nums[0])
	num2 := getNumber(nums[1])

	return num1 * num2
}

func getNumber(num string) int {
	// Extract the number from the string
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}

	match := re.FindString(num)
	numInt := 0
	if match != "" {
		numInt = convertToInt(match)
	}

	return numInt
}

func convertToInt(num string) int {
	// Convert the string to an integer
	numInt, err := strconv.Atoi(num)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	return numInt
}

func processInput(inputFile string) []byte {
	// Read the file data
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	newData := removeBetween(string(data), "don't()", "do()")

	for {
		newData = strings.Replace(newData, "don't()do()", "do()", -1)
		newData = removeBetween(newData, "don't()", "do()")

		if !strings.Contains(newData, "don't()do()") {
			break
		}
	}

	newData = removeAfter(newData, "don't()")

	return []byte(newData)
}

func removeBetween(input, start, end string) string {
	// Define the regular expression pattern with the (?s) flag to ensure the pattern matches across multiple lines
	pattern := fmt.Sprintf(`(?s)%s.*?%s`, regexp.QuoteMeta(start), regexp.QuoteMeta(end))

	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}

	// Replace the matched content with an empty string
	result := re.ReplaceAllString(input, start+end)

	return result
}

func removeAfter(input, substring string) string {
	// Find the index of the substring
	index := strings.Index(input, substring)
	if index == -1 {
		// Substring not found, return the original string
		return input
	}

	// Slice the string up to the position of the substring
	return input[:index+len(substring)]
}

// This is either a genius or stupid way to solve this!
func main() {

	// Part 1
	data := readInput("input.txt")
	mults := getMatches(data, `mul\(\d+,\d+\)`)

	fmt.Printf("Total of multiplications: %d\n", processInstructions(mults))

	// Part 2
	data = processInput("input.txt")
	mults = getMatches(data, `mul\(\d+,\d+\)`)
	fmt.Printf("Total of conditional multiplications: %d\n", processInstructions(mults))

}
