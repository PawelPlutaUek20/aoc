package Day_16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func A() {
	file, _ := os.Open("Day_16/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, strings.Split(text, ""))
	}

	initialPoint := Point{x: 0, y: 0}
	seen := []Seen{}

	walk(&lines, &seen, initialPoint, east)

	result := make(map[Point]bool, 0)
	for _, point := range seen {
		result[Point{x: point.x, y: point.y}] = true
	}

	fmt.Println(len(result))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Point struct {
	x int
	y int
}

type Seen struct {
	x   int
	y   int
	dir int
}

const (
	north = iota
	east
	south
	west
)

func walk(maze *[][]string, seen *[]Seen, curr Point, dir int) bool {
	seeing := Seen{x: curr.x, y: curr.y, dir: dir}

	if slices.Contains(*seen, seeing) {
		return false
	}

	if curr.x < 0 || curr.y < 0 || curr.y >= len(*maze) || curr.x >= len((*maze)[curr.y]) {
		return false
	}

	dirs := [4]Point{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}

	*seen = append(*seen, seeing)

	tile := (*maze)[curr.y][curr.x]
	nextDirs := getNextDirs(tile, dir)

	for _, dir := range nextDirs {
		offset := dirs[dir]
		next := Point{x: curr.x + offset.x, y: curr.y + offset.y}
		walk(maze, seen, next, dir)
	}

	return false
}

func getNextDirs(tile string, dir int) []int {
	switch tile {
	case ".":
		return []int{dir}
	case "/":
		switch dir {
		case north:
			return []int{east}
		case east:
			return []int{north}
		case south:
			return []int{west}
		case west:
			return []int{south}
		default:
			panic(dir)
		}
	case "\\":
		switch dir {
		case north:
			return []int{west}
		case east:
			return []int{south}
		case south:
			return []int{east}
		case west:
			return []int{north}
		default:
			panic(dir)
		}
	case "|":
		switch dir {
		case north, south:
			return []int{dir}
		case east, west:
			return []int{north, south}
		default:
			panic(dir)
		}
	case "-":
		switch dir {
		case north, south:
			return []int{east, west}
		case east, west:
			return []int{dir}
		default:
			panic(dir)
		}
	default:
		panic(tile)
	}
}
