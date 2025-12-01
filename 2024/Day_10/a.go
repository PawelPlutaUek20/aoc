package Day_10

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X, Y int
}

var dirs = [4]Position{
	{Y: -1, X: 0},
	{Y: 0, X: 1},
	{Y: 1, X: 0},
	{Y: 0, X: -1},
}

func A() {
	file, err := os.ReadFile("Day_10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]int, 0, len(lines))
	for _, line := range lines {
		heights := []int{}
		for _, stringHeight := range strings.Split(line, "") {
			height, _ := strconv.Atoi(stringHeight)
			heights = append(heights, height)
		}
		grid = append(grid, heights)
	}

	trailHeads := []Position{}

	for y, row := range grid {
		for x, col := range row {
			if col == 0 {
				trailHeads = append(trailHeads, Position{X: x, Y: y})
			}
		}
	}

	score := walk(grid, trailHeads)

	log.Println(score)
}

func walk(grid [][]int, trailHeads []Position) int {

	positions := map[Position]map[Position]bool{}
	for _, trailHead := range trailHeads {
		positions[trailHead] = map[Position]bool{}
		newPositions := walkWithState(grid, trailHead, trailHead, 0)
		for position := range newPositions {
			positions[trailHead][position] = true
		}
	}

	sum := 0
	for _, targetPositions := range positions {
		sum += len(targetPositions)
	}

	return sum
}

func walkWithState(grid [][]int, position Position, trailHead Position, index int) map[Position]bool {
	if position.X < 0 || position.X >= len(grid[0]) || position.Y < 0 || position.Y >= len(grid) {
		return map[Position]bool{}
	}

	if index != grid[position.Y][position.X] {
		return map[Position]bool{}
	}

	if index == grid[position.Y][position.X] && index == 9 {
		return map[Position]bool{position: true}
	}

	positions := map[Position]bool{}
	for _, dir := range dirs {
		nextPos := Position{X: position.X + dir.X, Y: position.Y + dir.Y}
		pathEnds := walkWithState(grid, nextPos, trailHead, index+1)
		for pathEnd := range pathEnds {
			positions[pathEnd] = true
		}
	}

	return positions
}
