package Day_18

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	row, col int
}

type PointWithStep struct {
	point Point
	step  int
}

var directions [4]Point = [4]Point{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func A() {
	file, err := os.ReadFile("Day_18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]rune, 71)
	for row := range grid {
		grid[row] = make([]rune, 71)
	}

	for i := 0; i < 1024; i++ {
		line := lines[i]

		coordinates := strings.Split(strings.TrimSpace(line), ",")

		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])

		grid[y][x] = '#'
	}

	steps := bfs(grid, Point{0, 0}, Point{70, 70})
	log.Println(steps)
}

func bfs(grid [][]rune, start Point, end Point) int {
	q := []PointWithStep{}

	seen := make([][]bool, len(grid))
	for i, row := range grid {
		seen[i] = make([]bool, len(row))
	}

	seen[start.row][start.col] = true
	q = append(q, PointWithStep{start, 0})

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		currPoint := curr.point
		for _, direction := range directions {
			nextPoint := Point{row: currPoint.row + direction.row, col: currPoint.col + direction.col}

			if nextPoint.col < 0 || nextPoint.col >= len(grid[0]) || nextPoint.row < 0 || nextPoint.row >= len(grid) {
				continue
			}

			if grid[nextPoint.row][nextPoint.col] == '#' {
				continue
			}

			if nextPoint == end {
				return curr.step + 1
			}

			if !seen[nextPoint.row][nextPoint.col] {
				seen[nextPoint.row][nextPoint.col] = true

				next := PointWithStep{point: nextPoint, step: curr.step + 1}
				q = append(q, next)
			}
		}
	}

	return -1
}
