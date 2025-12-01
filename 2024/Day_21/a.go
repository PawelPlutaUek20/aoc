package Day_21

import (
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_21/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	doorCodes := strings.Split(strings.TrimSpace(input), "\n")

	sum := 0

	for _, doorCode := range doorCodes {
		bestSteps := getBestSteps(numericKeypad, numericKeypadStart, doorCode)
		bestPaths := CartesianProduct(bestSteps)

		for i := 0; i < 2; i++ {
			nextPaths := make([]string, 0)
			for _, path := range bestPaths {
				tempBestSteps := getBestSteps(directionalKeypad, directionalKeypadStart, path)
				bestPaths := CartesianProduct(tempBestSteps)
				nextPaths = append(nextPaths, bestPaths...)
			}

			min := slices.MinFunc(nextPaths, func(a, b string) int {
				return len(a) - len(b)
			})

			bestOfBest := make([]string, 0)
			for _, nextPath := range nextPaths {
				if len(nextPath) == len(min) {
					bestOfBest = append(bestOfBest, nextPath)
				}
			}

			bestPaths = bestOfBest
		}

		doorCodeNumber := getDoorCodeNumber(doorCode)
		bestPathLen := len(bestPaths[0])
		sum += doorCodeNumber * bestPathLen
	}

	log.Println(sum)
}

func getDoorCodeNumber(doorCode string) int {
	lineInt, _ := strconv.Atoi(strings.TrimSuffix(doorCode, "A"))
	return lineInt
}

func getBestSteps(keypad [][]string, start Point, target string) [][]string {
	result := make([][]string, 0)

	symbols := strings.Split(strings.TrimSpace(target), "")
	currStart := start

	for _, symbol := range symbols {
		path := bfs(keypad, currStart, symbol)
		currStart = getPositionOfTarget(keypad, symbol)
		result = append(result, path)
	}

	return result
}

type Point struct {
	row, col int
}

var numericKeypadStart Point = Point{3, 2}
var numericKeypad [][]string = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"", "0", "A"},
}

var directionalKeypadStart Point = Point{0, 2}
var directionalKeypad [][]string = [][]string{
	{"", "^", "A"},
	{"<", "v", ">"},
}

var directionsMap map[string]Point = map[string]Point{
	"^": {-1, 0},
	">": {0, 1},
	"v": {1, 0},
	"<": {0, -1},
}

type QueueItem struct {
	point Point
	path  string
}

func getPositionOfTarget(keypad [][]string, target string) Point {
	for row, cols := range keypad {
		for col, item := range cols {
			if item == target {
				return Point{row, col}
			}
		}
	}

	panic("could not find target in keypad")
}

func bfs(numericKeypad [][]string, start Point, target string) []string {
	result := make([]string, 0)
	best := math.MaxInt

	if numericKeypad[start.row][start.col] == target {
		result = append(result, "A")
		return result
	}

	q := make([]QueueItem, 0)
	q = append(q, QueueItem{start, ""})

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		pos := curr.point
		path := curr.path

		for arrow, direction := range directionsMap {
			nextPos := Point{pos.row + direction.row, pos.col + direction.col}
			if nextPos.row < 0 || nextPos.row >= len(numericKeypad) || nextPos.col < 0 || nextPos.col >= len(numericKeypad[0]) {
				continue
			}

			if numericKeypad[nextPos.row][nextPos.col] == "" {
				continue
			}

			nextPath := path + arrow
			if len(nextPath) > best {
				break
			}

			nextQueueItem := QueueItem{nextPos, nextPath}

			if numericKeypad[nextPos.row][nextPos.col] == target {
				best = len(nextPath)
				result = append(result, nextQueueItem.path+"A")
			}

			q = append(q, nextQueueItem)
		}
	}

	return result
}

func CartesianProduct(arr [][]string) []string {
	if len(arr) == 0 {
		return []string{}
	}

	result := arr[0]

	for i := 1; i < len(arr); i++ {
		var temp []string
		for _, prefix := range result {
			for _, item := range arr[i] {
				temp = append(temp, prefix+item)
			}
		}
		result = temp
	}

	return result
}
