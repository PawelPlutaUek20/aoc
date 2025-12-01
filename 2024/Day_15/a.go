package Day_15

import (
	"log"
	"os"
	"strings"
)

const ROBOT = '@'
const BOX = 'O'
const WALL = '#'
const PATH = '.'

type Point struct {
	X, Y int
}

var directions = [4]Point{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

var directionByMove = map[string]Point{
	"^": directions[0],
	">": directions[1],
	"v": directions[2],
	"<": directions[3],
}

func A() {
	file, err := os.ReadFile("Day_15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	grid := parseGrid(lines[0])
	moves := parseMoves(lines[1])

	startingPosition, startingGrid := findStartingPosition(grid)

	result := walk(startingGrid, moves, startingPosition)
	boxes := findAllBoxes(result)

	sum := 0
	for _, box := range boxes {
		sum += 100*box.Y + box.X
	}

	log.Println(sum)
}

func findAllBoxes(grid [][]rune) []Point {
	result := make([]Point, 0)

	for y, row := range grid {
		for x, item := range row {
			if item == BOX {
				result = append(result, Point{x, y})
			}
		}
	}

	return result
}

func walk(grid [][]rune, moves []string, startingPosition Point) [][]rune {
	result := grid
	currPosition := startingPosition

	for _, move := range moves {
		currDirection := directionByMove[move]
		nextPosition := Point{currPosition.X + currDirection.X, currPosition.Y + currDirection.Y}
		if result[nextPosition.Y][nextPosition.X] == WALL {
			continue
		} else if result[nextPosition.Y][nextPosition.X] == PATH {
			currPosition = nextPosition
		} else {
			tempPosition := nextPosition
			for result[tempPosition.Y][tempPosition.X] != WALL {
				if result[tempPosition.Y][tempPosition.X] == BOX {
					tempPosition = Point{tempPosition.X + currDirection.X, tempPosition.Y + currDirection.Y}
					continue
				}
				if result[tempPosition.Y][tempPosition.X] == PATH {
					result[tempPosition.Y][tempPosition.X] = BOX
					result[nextPosition.Y][nextPosition.X] = PATH
					currPosition = nextPosition
					break
				}
			}
		}
	}

	return result
}

func findStartingPosition(grid [][]rune) (Point, [][]rune) {
	result := Point{}

	for y, rows := range grid {
		for x, item := range rows {
			if item == ROBOT {
				result.X = x
				result.Y = y

				grid[y][x] = PATH
			}
		}
	}

	return result, grid
}

func parseMoves(moves string) []string {
	ignoreNewlines := strings.ReplaceAll(moves, "\n", "")
	parsed := strings.Split(strings.TrimSpace(ignoreNewlines), "")
	return parsed
}

func parseGrid(grid string) [][]rune {
	parsed := make([][]rune, 0)

	rows := strings.Split(strings.TrimSpace(grid), "\n")
	for i, row := range rows {
		parsed = append(parsed, make([]rune, len(row)))
		for j, item := range row {
			parsed[i][j] = item
		}
	}

	return parsed
}
