package Day_06

import (
	"log"
	"os"
	"strings"
)

func B() {
	file, err := os.ReadFile("Day_06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(input, "\n")

	grid := [][]byte{}
	for _, line := range lines {
		if line != "" {
			grid = append(grid, []byte(line))
		}
	}

	output := 0

	initialPos, _ := getInitialDirAndPos(grid)
	visited, _ := walk(grid, initialPos, 0)

	for visitedPos := range visited {
		grid[visitedPos.y][visitedPos.x] = WALL

		_, isLoop := walk(grid, initialPos, 0)
		if isLoop {
			output++
		}

		grid[visitedPos.y][visitedPos.x] = '.'
	}

	log.Println(output)
}

func walk(grid [][]byte, initialPos Position, initialDir int) (map[Position]int, bool) {
	visitedPositions := make(map[Position]int)
	isLoop := false

	currPos, currDir := initialPos, initialDir
	visitedPositions[currPos] = currDir

	for {
		nextPos := nextPosition(currPos, currDir)

		if !isInBounds(grid, nextPos) {
			break
		}

		visitedDir, wasVisited := visitedPositions[nextPos]
		if wasVisited && visitedDir == currDir {
			isLoop = true
			break
		}

		if grid[nextPos.y][nextPos.x] == '#' {
			currDir = (currDir + 1) % 4
			continue
		}

		currPos = nextPos
		visitedPositions[nextPos] = currDir
	}

	return visitedPositions, isLoop
}
