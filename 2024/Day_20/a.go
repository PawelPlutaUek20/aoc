package Day_20

import (
	"log"
	"os"
	"strings"
)

type Point struct {
	row, col int
}

var directions [4]Point = [4]Point{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

var cheatableWalls map[Point]bool = make(map[Point]bool)

func getMaze(lines []string) [][]rune {
	maze := make([][]rune, 0)
	for row, line := range lines {
		maze = append(maze, make([]rune, len(line)))
		for col, r := range line {
			maze[row][col] = r
		}
	}
	return maze
}

func A() {

	file, err := os.ReadFile("Day_20/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	maze := getMaze(lines)
	start := find(maze, 'S')
	end := find(maze, 'E')

	pathTime := getPath(maze, start, end)

	bestCheatTimes := make([]int, 0)

	for cheatableWall := range cheatableWalls {
		cheatMaze := getMaze(lines)
		cheatMaze[cheatableWall.row][cheatableWall.col] = '.'

		cheatPathTime := getPath(cheatMaze, start, end)

		if pathTime-cheatPathTime >= 100 {
			bestCheatTimes = append(bestCheatTimes, cheatPathTime)
		}
	}

	log.Println(len(bestCheatTimes))
}

type QueueItem struct {
	point Point
	len   int
}

func getPath(maze [][]rune, start Point, end Point) int {
	seen := make([][]bool, len(maze))
	for i := range maze {
		seen[i] = make([]bool, len(maze[i]))
	}

	stack := make([]QueueItem, 0)

	seen[start.row][start.col] = true
	stack = append(stack, QueueItem{start, 0})

	for len(stack) > 0 {
		poped := stack[0]
		stack = stack[1:]

		for _, direction := range directions {
			next := Point{poped.point.row + direction.row, poped.point.col + direction.col}
			if maze[next.row][next.col] == '#' {
				nextNext := Point{next.row + direction.row, next.col + direction.col}
				if nextNext.row >= 0 && nextNext.col >= 0 && nextNext.row < len(maze) && nextNext.col < len(maze[0]) {
					if maze[nextNext.row][nextNext.col] != '#' {
						cheatableWalls[next] = true
					}
				}
				continue
			}
			if seen[next.row][next.col] {
				continue
			}

			seen[next.row][next.col] = true
			stack = append(stack, QueueItem{next, poped.len + 1})

			if next == end {
				return poped.len + 1
			}
		}
	}

	return -1
}

func find(maze [][]rune, needle rune) Point {
	for row, cols := range maze {
		for col := range cols {
			if maze[row][col] == needle {
				return Point{row, col}
			}
		}
	}

	panic("could not find")
}
