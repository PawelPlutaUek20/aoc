package Day_10

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	north = iota
	east
	south
	west
)

type Position struct {
	x int
	y int
}

func A() {

	file, _ := os.Open("Day_10/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]string, 0)

	startPosition := Position{x: -1, y: -1}

	for scanner.Scan() {
		text := scanner.Text()

		x := strings.Index(text, "S")
		if x != -1 {
			startPosition = Position{x: x, y: len(lines)}
		}

		lines = append(lines, strings.Split(text, ""))
	}

	for i := 0; i < 4; i++ {
		path := make([]Position, 0)
		seen := make([]Position, 0)
		if walk(&lines, &path, &seen, i, startPosition) {
			fmt.Println(len(path) / 2)
			break
		}
	}

}

func walk(maze *[][]string, path *[]Position, seen *[]Position, dir int, curr Position) bool {
	dirs := [4]Position{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}

	if curr.x < 0 || curr.y < 0 || curr.y >= len(*maze) || curr.x >= len((*maze)[curr.y]) {
		return false
	}

	if (*maze)[curr.y][curr.x] == "S" && slices.Contains(*seen, curr) {
		*path = append(*path, curr)
		return true
	}

	if slices.Contains(*seen, curr) {
		return false
	}

	*path = append(*path, curr)
	*seen = append(*seen, curr)

	pipe := (*maze)[curr.y][curr.x]
	nextDir, isValidPipe := getNextDir(pipe, dir)
	if isValidPipe {
		nextOffest := dirs[nextDir]
		nextCurr := Position{x: curr.x + nextOffest.x, y: curr.y + nextOffest.y}
		if walk(maze, path, seen, nextDir, nextCurr) {
			return true
		}
	}

	return false
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position.
func getNextDir(pipe string, dir int) (int, bool) {
	switch pipe {
	case "S":
		return dir, true
	case "|":
		return dir, true
	case "-":
		return dir, true
	case "L":
		if dir == south {
			return east, true
		} else if dir == west {
			return north, true
		} else {
			return -1, false
		}
	case "J":
		if dir == south {
			return west, true
		} else if dir == east {
			return north, true
		} else {
			return -1, false
		}
	case "7":
		if dir == east {
			return south, true
		} else if dir == north {
			return west, true
		} else {
			return -1, false
		}
	case "F":
		if dir == north {
			return east, true
		} else if dir == west {
			return south, true
		} else {
			return -1, false
		}
	default:
		return -1, false
	}
}
