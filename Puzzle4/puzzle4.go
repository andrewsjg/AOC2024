package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(inputFile string) map[int]string {
	// Read the file line by line and return the data as a map
	data := make(map[int]string)

	// Open the file

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		data[i] = scanner.Text()
		i++
	}

	return data

}

func part1(inputfile string) int {
	// Read the input file
	data := readInput(inputfile)

	count := 0

	for i := 0; i < len(data); i++ {

		line := data[i]
		runes := []rune(line)

		for j := 0; j < len(runes); j++ {
			// Check for "XMAS"
			if j+3 < len(runes) && string(runes[j:j+4]) == "XMAS" {
				count++

			}

			// Check for "SAMX"
			if j+3 < len(runes) && string(runes[j:j+4]) == "SAMX" {
				count++

			}

		}

		for cIdx, char := range runes {
			if char == 'X' {

				// Find left up diagonal
				if i > 2 && cIdx > 2 {
					l1 := []rune(data[i-1])
					l2 := []rune(data[i-2])
					l3 := []rune(data[i-3])

					if cIdx-1 < len(l1) && cIdx-2 < len(l2) && cIdx-3 < len(l3) {
						if l1[cIdx-1] == 'M' && l2[cIdx-2] == 'A' && l3[cIdx-3] == 'S' {

							count++
						}
					}
				}

				// Find right up diagonal
				if i > 2 && cIdx < len(runes)-3 {
					l1 := []rune(data[i-1])
					l2 := []rune(data[i-2])
					l3 := []rune(data[i-3])

					if cIdx+1 < len(l1) && cIdx+2 < len(l2) && cIdx+3 < len(l3) {

						if l1[cIdx+1] == 'M' && l2[cIdx+2] == 'A' && l3[cIdx+3] == 'S' {

							count++
						}
					}
				}

				// find left down diagonal
				if i < len(data)-3 && cIdx > 2 {
					l1 := []rune(data[i+1])
					l2 := []rune(data[i+2])
					l3 := []rune(data[i+3])

					if cIdx-1 < len(l1) && cIdx-2 < len(l2) && cIdx-3 < len(l3) {
						if l1[cIdx-1] == 'M' && l2[cIdx-2] == 'A' && l3[cIdx-3] == 'S' {

							count++
						}
					}
				}

				// find right down diagonal
				if i < len(data)-3 && cIdx < len(runes)-3 {
					l1 := []rune(data[i+1])
					l2 := []rune(data[i+2])
					l3 := []rune(data[i+3])

					if cIdx+1 < len(l1) && cIdx+2 < len(l2) && cIdx+3 < len(l3) {
						if l1[cIdx+1] == 'M' && l2[cIdx+2] == 'A' && l3[cIdx+3] == 'S' {

							count++
						}
					}
				}

				// Find below
				l1 := []rune(data[i+1])
				if cIdx < len(l1) && cIdx >= 0 {
					if l1[cIdx] == 'M' {
						l2 := []rune(data[i+2])
						if cIdx < len(l2) && cIdx >= 0 {
							if l2[cIdx] == 'A' {
								l3 := []rune(data[i+3])
								if cIdx < len(l3) && cIdx >= 0 {
									if l3[cIdx] == 'S' {

										count++
									}
								}
							}
						}
					}
				}

				// Find above
				l1 = []rune(data[i-1])
				if cIdx < len(l1) && cIdx >= 0 {
					if l1[cIdx] == 'M' {
						l2 := []rune(data[i-2])
						if cIdx < len(l2) && cIdx >= 0 {
							if l2[cIdx] == 'A' {
								l3 := []rune(data[i-3])
								if cIdx < len(l3) && cIdx >= 0 {
									if l3[cIdx] == 'S' {

										count++
									}
								}
							}
						}
					}
				}
			}

		}

	}

	return count
}

func part2(inputfile string) int {
	// Read the input file
	data := readInput(inputfile)
	count := 0

	// M == 77, S == 83 M+S == 160

	// Find 'A' in the data
	for i := 0; i < len(data); i++ {
		line := data[i]
		runes := []rune(line)

		for j := 0; j < len(runes); j++ {
			if runes[j] == 'A' {

				UL := ' '
				UR := ' '
				DL := ' '
				DR := ' '

				// Get the letter up and to the left
				if i > 0 && j > 0 {
					l1 := []rune(data[i-1])
					UL = l1[j-1]
				}

				// Get the letter down and to the right
				if i < len(data)-1 && j < len(runes)-1 {
					l1 := []rune(data[i+1])
					DR = l1[j+1]

				}

				if UL+DR == 160 {
					// Get the letter up and to the right
					if i > 0 && j < len(runes)-1 {

						l1 := []rune(data[i-1])
						UR = l1[j+1]
					}

					// Get the letter down and to the left
					if i < len(data)-1 && j > 0 {
						l1 := []rune(data[i+1])
						DL = l1[j-1]

					}

					if UR+DL == 160 {
						count++
					}

				}

			}

		}
	}

	return count
}

func main() {
	fmt.Printf("Part 1 Total: %d\n", part1("input.txt"))
	fmt.Printf("Part 2 Total: %d\n", part2("input.txt"))
}
