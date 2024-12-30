package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Read input file. It will be one long line
func readInput(inputFile string) string {

	// Open the file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	data := ""

	for scanner.Scan() {
		data = scanner.Text()
	}

	return data
}

type SparseArray struct {
	data map[int]int
}

func NewSparseArray() *SparseArray {
	return &SparseArray{
		data: make(map[int]int),
	}
}

func (sa *SparseArray) Set(index int, value int) {
	if value != 0 {
		sa.data[index] = value
	} else {
		delete(sa.data, index)
	}
}

func (sa *SparseArray) Get(index int) int {
	if value, exists := sa.data[index]; exists {
		return value
	}
	return 0
}

func (sa *SparseArray) Last() (int, int) {
	if len(sa.data) == 0 {
		return -1, 0 // Return -1 if the sparse array is empty
	}

	maxIndex := -1

	for index := range sa.data {
		if index > maxIndex {
			maxIndex = index
		}
	}

	return maxIndex, sa.data[maxIndex]
}

func (sa *SparseArray) Copy() *SparseArray {
	newSparseArray := NewSparseArray()
	for index, value := range sa.data {
		newSparseArray.Set(index, value)
	}
	return newSparseArray
}

func (sa *SparseArray) MoveBlock(srcStart, srcEnd, destStart int) {
	// Copy the block to the destination
	for i := srcStart; i <= srcEnd; i++ {
		value := sa.Get(i)
		sa.Set(destStart+(i-srcStart), value)
	}

	// Clear the source block
	for i := srcStart; i <= srcEnd; i++ {
		sa.Set(i, 0)
	}
}

func (sa *SparseArray) FindLowestGap(size int) int {
	// Find the lowest index gap that can accommodate the set of values
	maxIndex, _ := sa.Last()
	for i := 0; i <= maxIndex; i++ {
		gap := true
		for j := 0; j < size; j++ {
			if sa.Get(i+j) != 0 {
				gap = false
				break
			}
		}
		if gap {
			return i
		}
	}
	return -1
}

func (sa *SparseArray) InsertAtLowestGap(values []int) {
	gapIndex := sa.FindLowestGap(len(values))
	for i, value := range values {
		sa.Set(gapIndex+i, value)
	}
}

func createDiskLayout(input string) SparseArray {
	readSize := false
	fileID := 1
	diskLayout := NewSparseArray()

	diskPointer := 0

	for _, char := range input {

		if !readSize {

			sizeVal, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatalf("Error converting char to integer: %v", err)
			}

			for i := 0; i < sizeVal; i++ {

				diskLayout.Set(diskPointer, fileID)
				diskPointer++
			}

			fileID++

			readSize = true

		} else {

			space, err := strconv.Atoi(string(char))

			if err != nil {
				log.Fatalf("Error converting char to integer: %v", err)
			}

			diskPointer += space
			readSize = false
		}

	}

	return *diskLayout
}

func defragDisk(diskLayout SparseArray) SparseArray {
	defraged := NewSparseArray()

	diskPointer := 0
	blankCount := 0

	lastIndex, _ := diskLayout.Last()

	for i := lastIndex; i >= 0; i-- {

		// Find the next blank slot
		diskLoc := diskLayout.Get(diskPointer)

		for diskLoc != 0 {
			defraged.Set(diskPointer, diskLoc)

			diskPointer++
			diskLoc = diskLayout.Get(diskPointer)
		}

		blankCount++

		if diskLayout.Get(i) != 0 {
			defraged.Set(diskPointer, diskLayout.Get(i))
			diskPointer++
		}

		testVal := (lastIndex - blankCount)
		if diskPointer >= testVal {
			defraged.Set(diskPointer, diskLayout.Get(i))

			// Trim the sparse array - Stupid! Need to figure out how to stop
			// the loop overshooting. Some messed up logic somewhere. But I am losing my
			// mind looking at it. So this bodge will do for now
			trim := diskPointer - (lastIndex - blankCount) - 1
			dfLast, _ := defraged.Last()

			if trim > 0 {
				for i := 0; i < trim; i++ {
					defraged.Set(dfLast-i, 0)
				}

			}

			break
		}

	}

	return *defraged
}

func defragDiskPart2(diskLayout SparseArray) SparseArray {

	defraged := diskLayout.Copy()

	lastIndex, lastFile := diskLayout.Last()

	fileSize := -1

	for i := lastIndex; i >= 0; i-- {

		currentFile := diskLayout.Get(i)
		// Get the first file to move
		if diskLayout.Get(i) != 0 && currentFile == lastFile {
			fileSize++

		} else {

			fileSize++
			if lastFile > 0 {
				// Move the file

				gapIndex := defraged.FindLowestGap(fileSize)

				if gapIndex > 0 {
					if gapIndex < i {
						defraged.MoveBlock(i+1, i+fileSize, gapIndex)
					}
				}

			}

			fileSize = 0
		}
		lastFile = currentFile

	}

	return *defraged
}

func printDiskLayout(diskLayout SparseArray) {
	lastIndex, currentFile := diskLayout.Last()
	fileCount := 0

	for i := 0; i < lastIndex+1; i++ {

		if diskLayout.Get(i) == 0 {
			fmt.Print(".")
		} else {
			if diskLayout.Get(i) != currentFile {
				fileCount++
			}

			fmt.Print(diskLayout.Get(i) - 1)
			currentFile = diskLayout.Get(i)
		}
	}
	fmt.Println()

}

func calculateChecksum(diskLayout SparseArray) int {
	checksum := 0

	lastIndex, _ := diskLayout.Last()

	for i := 0; i <= lastIndex; i++ {

		if diskLayout.Get(i) > 0 {
			checksum += i * (diskLayout.Get(i) - 1)
		} else {
			checksum += i * (diskLayout.Get(i))
		}
	}

	return checksum
}

func printSparseArray(sa SparseArray) {
	lastIndex, _ := sa.Last()

	for i := 0; i <= lastIndex; i++ {
		fmt.Printf("%d: %d\n", i, sa.Get(i))
	}
}

func main() {

	// I made this SO much harder than it needed to be
	// I thought sparse arrays would be clever. But I think they made it harder.
	data := readInput("input.txt")

	diskLayout := createDiskLayout(data)

	defragged := defragDisk(diskLayout)
	defragged2 := defragDiskPart2(diskLayout)

	fmt.Printf("Part 1 - Defragged disk checksum: %d\n", calculateChecksum(defragged))
	fmt.Printf("Part 2 - Defragged disk checksum: %d\n", calculateChecksum(defragged2))

}
