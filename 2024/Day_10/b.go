package Day_10

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func B() {
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

	score := walkB(grid, trailHeads)

	log.Println(score)
}

func walkB(grid [][]int, trailHeads []Position) int {
	sum := 0
	for _, trailHead := range trailHeads {
		rating := walkWithStateB(grid, trailHead, trailHead, 0)
		sum += rating
	}

	return sum
}

func walkWithStateB(grid [][]int, position Position, trailHead Position, index int) int {
	if position.X < 0 || position.X >= len(grid[0]) || position.Y < 0 || position.Y >= len(grid) {
		return 0
	}

	if index != grid[position.Y][position.X] {
		return 0
	}

	if index == grid[position.Y][position.X] && index == 9 {
		return 1
	}

	rating := 0
	for _, dir := range dirs {
		newPos := Position{X: position.X + dir.X, Y: position.Y + dir.Y}
		rating += walkWithStateB(grid, newPos, trailHead, index+1)
	}

	return rating
}
