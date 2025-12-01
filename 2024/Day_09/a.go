package Day_09

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	ID int
}

func A() {
	file, err := os.ReadFile("Day_09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	diskMap := strings.Split(input, "")

	blocks := []*File{}
	id := 0

	for diskIdx, diskMap := range diskMap {
		if diskIdx%2 == 0 {
			// file
			blocksCount, _ := strconv.Atoi(diskMap)
			for i := 0; i < blocksCount; i++ {
				blocks = append(blocks, &File{ID: id})
			}
			id++
		} else {
			// free space
			blocksCount, _ := strconv.Atoi(diskMap)
			for i := 0; i < blocksCount; i++ {
				blocks = append(blocks, nil)
			}
		}
	}

	freeSpaceIndices := []int{}
	for index, block := range blocks {
		if block == nil {
			freeSpaceIndices = append(freeSpaceIndices, index)
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if len(freeSpaceIndices) == 0 {
			break
		}

		freeSpaceIndex := freeSpaceIndices[0]
		if freeSpaceIndex >= i {
			break
		}

		block := blocks[i]
		if block == nil {
			continue
		}

		blocks[i] = blocks[freeSpaceIndex]
		blocks[freeSpaceIndex] = block

		freeSpaceIndices = freeSpaceIndices[1:]
	}

	output := 0

	for idx, a := range blocks {
		if a == nil {
			continue
		}

		sNum := a.ID
		output += idx * sNum
	}

	log.Println(output)
}
