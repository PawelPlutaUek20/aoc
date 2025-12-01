package Day_15

import (
	"log"
	"os"
	"slices"
	"strings"
)

func B() {
	file, err := os.ReadFile("Day_15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	grid := parseGridB(lines[0])
	moves := parseMoves(lines[1])
	startingPosition, startingGrid := findStartingPosition(grid)

	result := walkB(startingGrid, moves, startingPosition)
	boxes := findAllBoxesB(result)

	sum := 0
	for _, box := range boxes {
		sum += 100*box.Y + box.X
	}

	log.Println(sum)
}

func findAllBoxesB(grid [][]rune) []Point {
	result := make([]Point, 0)

	for y, row := range grid {
		for x, item := range row {
			if item == '[' {
				result = append(result, Point{x, y})
			}
		}
	}

	return result
}

func walkB(grid [][]rune, moves []string, startingPosition Point) [][]rune {
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
			group := dfs(result, currPosition, currDirection)
			if group != nil {
				result[nextPosition.Y][nextPosition.X] = PATH
				for _, box := range group {
					result[box.pos.Y][box.pos.X] = PATH
				}
				for _, box := range group {
					result[box.pos.Y+currDirection.Y][box.pos.X+currDirection.X] = box.chr
				}
				currPosition = nextPosition
			}
		}
	}

	return result
}

type Box struct {
	chr rune
	pos Point
}

func dfs(grid [][]rune, startingPoint Point, direction Point) []Box {
	visited := make([]Point, 0)
	stack := make([]Point, 0)
	stack = append(stack, startingPoint)

	for len(stack) > 0 {
		poped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		popedItem := grid[poped.Y][poped.X]

		directions := []Point{direction}
		if popedItem == ']' {
			directions = append(directions, Point{-1, 0})
		} else if popedItem == '[' {
			directions = append(directions, Point{1, 0})
		}

		for _, direction := range directions {
			next := Point{poped.X + direction.X, poped.Y + direction.Y}
			nextItem := grid[next.Y][next.X]
			if slices.Contains(visited, next) {
				continue
			}

			if nextItem == WALL {
				return nil
			}

			if nextItem == PATH {
				continue
			}

			visited = append(visited, next)
			stack = append(stack, next)
		}
	}

	boxes := make([]Box, len(visited))
	for i, box := range visited {
		boxes[i] = Box{chr: grid[box.Y][box.X], pos: box}
	}

	return boxes
}

func parseGridB(grid string) [][]rune {
	parsed := make([][]rune, 0)

	rows := strings.Split(strings.TrimSpace(grid), "\n")
	for i, row := range rows {
		parsed = append(parsed, make([]rune, 2*len(row)))
		for j, item := range row {
			if item == WALL || item == PATH {
				parsed[i][2*j] = item
				parsed[i][2*j+1] = item
			} else if item == BOX {
				parsed[i][2*j] = '['
				parsed[i][2*j+1] = ']'
			} else if item == ROBOT {
				parsed[i][2*j] = ROBOT
				parsed[i][2*j+1] = PATH
			}
		}
	}

	return parsed
}
