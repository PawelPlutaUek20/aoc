package Day_09

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type FilesBlock struct {
	ID     int
	Len    int
	IsFree bool
}

func B() {
	file, err := os.ReadFile("Day_09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	diskMap := strings.Split(input, "")

	blocks := []FilesBlock{}
	id := 0

	for diskIdx, diskMap := range diskMap {
		if diskIdx%2 == 0 {
			// file
			blocksCount, _ := strconv.Atoi(diskMap)
			newBlock := FilesBlock{ID: id, IsFree: false, Len: blocksCount}
			blocks = append(blocks, newBlock)
			id++
		} else {
			// free space
			blocksCount, _ := strconv.Atoi(diskMap)
			newBlock := FilesBlock{IsFree: true, Len: blocksCount}
			blocks = append(blocks, newBlock)
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		blockA := blocks[i]
		if blockA.IsFree {
			continue
		}

		for j := 0; j < i; j++ {
			blockB := blocks[j]
			if !blockB.IsFree {
				continue
			}

			if blockB.Len < blockA.Len {
				continue
			}

			if blockB.Len == blockA.Len {
				blocks[i] = blockB
				blocks[j] = blockA
				break
			}

			if blockB.Len > blockA.Len {
				newBlocks := []FilesBlock{}
				newBlocks = append(newBlocks, blocks[0:j]...)
				newBlocks = append(newBlocks, blockA)
				newBlocks = append(newBlocks, FilesBlock{IsFree: true, Len: blockB.Len - blockA.Len})
				newBlocks = append(newBlocks, blocks[j+1:i]...)
				newBlocks = append(newBlocks, FilesBlock{IsFree: true, Len: blockA.Len})
				newBlocks = append(newBlocks, blocks[i+1:]...)

				blocks = newBlocks
				break
			}
		}
	}

	output := 0

	files := []*File{}
	for _, block := range blocks {
		if block.IsFree {
			for i := 0; i < block.Len; i++ {
				files = append(files, nil)
			}
		} else {
			for i := 0; i < block.Len; i++ {
				files = append(files, &File{ID: block.ID})
			}
		}
	}

	for idx, file := range files {
		if file == nil {
			continue
		}

		output += idx * file.ID
	}

	log.Println(output)
}
