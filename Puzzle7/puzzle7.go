package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(inputFile string, part2 bool) int {

	result := 0
	// Open the file

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := scanner.Text()

		numbers := extractNumbers(data)

		answer := numbers[0]
		values := numbers[1:]

		if validResult(answer, values, part2) {

			result = result + answer
		}
	}

	return result
}

func sumNumbers(numbers []int) int {
	total := 1

	for _, num := range numbers {
		total = total * num
	}

	return total
}

func validResult(total int, values []int, part2 bool) bool {
	return tester(total, values, 0, 0, part2)
}

func tester(total int, values []int, index int, current int, part2 bool) bool {

	if index == len(values) {
		return current == total
	}

	// Try addition
	if tester(total, values, index+1, current+values[index], part2) {
		return true
	}

	// Try multiplication
	if tester(total, values, index+1, current*values[index], part2) {
		return true
	}

	// Do concatenation
	if part2 {

		// I needed some hints to figure this out!
		numDigits := int(math.Floor(math.Log10(float64(values[index]))) + 1)
		concatenated := current*int(math.Pow(10, float64(numDigits))) + values[index]

		if tester(total, values, index+1, concatenated, part2) {
			return true

		}
	}

	return false
}

func extractNumbers(equation string) []int {
	numbers := []int{}

	tmp1 := strings.Split(equation, ":")
	tmp2 := strings.Split(strings.TrimSpace(tmp1[1]), " ")

	num, err := strconv.Atoi(tmp1[0])

	if err != nil {
		fmt.Println("Error converting string to int")
		os.Exit(1)
	}

	numbers = append(numbers, num)

	tmpArr, err := convertToIntArray(tmp2)

	if err != nil {
		fmt.Println("Error converting string to int")
		os.Exit(1)
	}

	numbers = append(numbers, tmpArr...)

	return numbers

}

// Convert an array of strings to an array of integers
func convertToIntArray(strArray []string) ([]int, error) {
	intArray := []int{}
	err := error(nil)

	for _, str := range strArray {
		num, err := strconv.Atoi(str)

		if err != nil {
			return intArray, err
		}

		intArray = append(intArray, num)
	}

	return intArray, err
}

func main() {

	fmt.Printf("Part 1 Total: %d\nPart 2 Total: %d \n", solve("input.txt", false), solve("input.txt", true))
}
