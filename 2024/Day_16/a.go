package Day_16

import (
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Fields(strings.TrimSpace(input))

	grid := make([][]string, len(lines))
	for i, lines := range lines {
		grid[i] = strings.Split(lines, "")
	}

	start := findStartingPosition(grid)
	end := findEndingPosition(grid)

	bestPath := findBestPath(grid, start, end)

	steps := len(bestPath) - 1
	turns := countTurns(bestPath)

	initDir := Point{0, 1}
	firstDir := Point{bestPath[1].row - bestPath[0].row, bestPath[1].col - bestPath[0].col}
	if firstDir != initDir {
		turns++
	}

	log.Println(turns*1000 + steps)
}

type Point struct {
	row, col int
}

var directions [4]Point = [4]Point{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func getId(maze [][]string, row int, col int) int {
	return row*len(maze[0]) + col
}

func getPoint(maze [][]string, id int) Point {
	return Point{row: id / len(maze[0]), col: id % len(maze[0])}
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i, s := range seen {
		if s == false && dists[i] < math.MaxInt {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	dist := math.MaxInt

	for i, s := range seen {
		if s == true {
			continue
		}

		if dists[i] < dist {
			dist = dists[i]
			idx = i
		}
	}

	return idx
}

func findBestPath(maze [][]string, start Point, end Point) []Point {
	startId := getId(maze, start.row, start.col)
	endId := getId(maze, end.row, end.col)

	seen := make([]bool, getId(maze, len(maze)-1, len(maze[0])-1))
	prev := make([]int, getId(maze, len(maze)-1, len(maze[0])-1))

	for i := range prev {
		prev[i] = -1
	}

	dists := make([]int, getId(maze, len(maze)-1, len(maze[0])-1))
	for i := range dists {
		dists[i] = math.MaxInt
	}

	dists[startId] = 0

	for hasUnvisited(seen, dists) {
		currId := getLowestUnvisited(seen, dists)
		seen[currId] = true

		curr := getPoint(maze, currId)

		for _, direction := range directions {
			next := Point{row: curr.row + direction.row, col: curr.col + direction.col}
			if maze[next.row][next.col] == "#" {
				continue
			}

			nextId := getId(maze, next.row, next.col)
			if seen[nextId] {
				continue
			}

			prevId := prev[currId]
			newDist := dists[currId] + calcDist(maze, prevId, currId, nextId)
			if newDist < dists[nextId] {
				dists[nextId] = newDist
				prev[nextId] = currId
			}
		}
	}

	path := []Point{}
	curr := endId

	for prev[curr] != -1 {
		path = append(path, getPoint(maze, curr))
		curr = prev[curr]
	}

	path = append(path, start)
	slices.Reverse(path)

	return path
}

func calcDist(maze [][]string, prevId int, currId int, nextId int) int {
	if prevId == -1 {
		currPoint, nextPoint := getPoint(maze, currId), getPoint(maze, nextId)

		initDir := Point{0, 1}
		nextDir := Point{nextPoint.row - currPoint.row, nextPoint.col - currPoint.col}

		if nextDir != initDir {
			return 1000
		} else {
			return 1
		}
	}

	points := []Point{getPoint(maze, prevId), getPoint(maze, currId), getPoint(maze, nextId)}
	return 1000*countTurns(points) + 1
}

func countTurns(path []Point) int {
	turns := 0

	for i := 1; i < len(path)-1; i++ {
		currDir := Point{path[i].row - path[i-1].row, path[i].col - path[i-1].col}
		nextDir := Point{path[i+1].row - path[i].row, path[i+1].col - path[i].col}

		if currDir != nextDir {
			turns++
		}
	}

	return turns
}

func findStartingPosition(grid [][]string) Point {
	result := Point{}

	for row, rows := range grid {
		for col, item := range rows {
			if item == "S" {
				result.col = col
				result.row = row

			}
		}
	}

	return result
}

func findEndingPosition(grid [][]string) Point {
	result := Point{}

	for row, rows := range grid {
		for col, item := range rows {
			if item == "E" {
				result.col = col
				result.row = row

			}
		}
	}

	return result
}
