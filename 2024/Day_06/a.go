package Day_06

import (
	"log"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

var directions = [4]Position{
	{x: 0, y: -1},
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: -1, y: 0},
}

const WALL = '#'

func A() {
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

	visitedPositions := make(map[Position]bool)
	currPos, currDir := getInitialDirAndPos(grid)

	visitedPositions[currPos] = true

	for {
		nextPos := nextPosition(currPos, currDir)
		if !isInBounds(grid, nextPos) {
			break
		}

		if grid[nextPos.y][nextPos.x] == '#' {
			currDir = (currDir + 1) % 4
			continue
		}

		visitedPositions[nextPos] = true
		currPos = nextPos
	}

	log.Println(len(visitedPositions))
}

func getInitialDirAndPos(grid [][]byte) (Position, int) {
	for y, row := range grid {
		for x, col := range row {
			if col == '^' {
				pos := Position{x: x, y: y}
				dir := 0
				return pos, dir
			}
		}
	}

	panic("could not find initial dir and pos")
}

func nextPosition(currPos Position, dir int) Position {
	dirDelta := directions[dir]

	return Position{
		x: currPos.x + dirDelta.x,
		y: currPos.y + dirDelta.y,
	}
}

func isInBounds(grid [][]byte, pos Position) bool {
	maxY := len(grid)
	maxX := len(grid[0])
	minY := 0
	minX := 0

	return pos.x >= minX && pos.x < maxX && pos.y >= minY && pos.y < maxY
}
