package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// RULES:
/*

- If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
- If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
- If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone

*/

func getStones(inputFile string) []string {
	stones := []string{}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		stones = strings.Split(scanner.Text(), " ")
	}

	return stones
}

// Naive approch. Works for a low number of stones.
func applyRules(stones []string) []string {
	newStones := []string{}

	for _, stone := range stones {

		if stone == "0" {
			newStones = append(newStones, "1")

		} else if len(stone)%2 == 0 {
			leftStone := stone[:len(stone)/2]
			rightStone := stone[len(stone)/2:]

			// StrConv will convert 0000 etc to 0 while leaving all other numbers the same.
			strLeftStone, err := strconv.Atoi(leftStone)
			if err != nil {
				log.Fatalf("Error converting int to string: %v", err)
			}

			strRightStone, err := strconv.Atoi(rightStone)
			if err != nil {
				log.Fatalf("Error converting int to string: %v", err)
			}

			leftStone = strconv.Itoa(strLeftStone)
			rightStone = strconv.Itoa(strRightStone)

			//fmt.Printf("Left stone: %s, Right Stone: %s\n", leftStone, rightStone)
			newStones = append(newStones, leftStone)
			newStones = append(newStones, rightStone)

		} else {

			intStone, err := strconv.Atoi(stone)
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}

			newStones = append(newStones, fmt.Sprintf("%d", 2024*intStone))
		}
	}

	return newStones
}

// Optimized approach. Simply count the occurences of each stone and apply the rules to the count.
func applyRules_v2(stoneMap map[string]int) map[string]int {

	newStoneMap := make(map[string]int)

	for stone, count := range stoneMap {

		if stone == "0" {
			newStoneMap["1"] += count

		} else if len(stone)%2 == 0 {
			leftStone := stone[:len(stone)/2]
			rightStone := stone[len(stone)/2:]

			// StrConv will convert 0000 etc to 0 while leaving all other numbers the same.
			strLeftStone, err := strconv.Atoi(leftStone)
			if err != nil {
				log.Fatalf("Error converting int to string: %v", err)
			}

			strRightStone, err := strconv.Atoi(rightStone)
			if err != nil {
				log.Fatalf("Error converting int to string: %v", err)
			}

			leftStone = strconv.Itoa(strLeftStone)
			rightStone = strconv.Itoa(strRightStone)

			newStoneMap[leftStone] += count
			newStoneMap[rightStone] += count

		} else {

			intStone, err := strconv.Atoi(stone)
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}

			newStoneMap[fmt.Sprintf("%d", 2024*intStone)] += count

		}
	}

	return newStoneMap
}

// Naive approach. Works for a low number of stones.
func part1(inputFile string, blinks int) int {
	stones := getStones(inputFile)

	for i := 1; i <= blinks; i++ {
		stones = applyRules(stones)
	}

	return len(stones)
}

func part2(inputFile string, blinks int) int {
	stones := getStones(inputFile)

	stoneMap := make(map[string]int)

	for _, stone := range stones {
		stoneMap[stone]++
	}

	for i := 1; i <= blinks; i++ {
		stoneMap = applyRules_v2(stoneMap)
	}

	stoneCount := 0

	for _, sCount := range stoneMap {

		stoneCount += sCount
	}

	return stoneCount
}

func main() {

	numBlinks := 75

	input := "input.txt"
	//fmt.Printf("Stones count after %d blinks = %d\n", numBlinks, part1(input, numBlinks))
	fmt.Printf("Stones count after %d blinks = %d\n", numBlinks, part2(input, numBlinks))

}
