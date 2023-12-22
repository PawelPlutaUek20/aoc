package Day_21

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func A() {

	file, _ := os.Open("Day_21/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	startPosition := Point{x: -1, y: -1}

	for scanner.Scan() {
		text := scanner.Text()

		sIndex := strings.Index(text, "S")
		if sIndex > -1 {
			startPosition = Point{x: sIndex, y: len(lines)}
		}

		lines = append(lines, strings.Split(text, ""))
	}

	result := bfs(&lines, GardenPlot{x: startPosition.x, y: startPosition.y, steps: 0})
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type GardenPlot struct {
	x     int
	y     int
	steps int
}

type Point struct {
	x int
	y int
}

func bfs(grid *[][]string, head GardenPlot) int {
	dirs := []Point{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}

	result := 0
	seen := []Point{Point{x: head.x, y: head.y}}
	queue := []GardenPlot{head}

	for len(queue) > 0 {
		curr, newQueue := dequeue(queue)
		queue = newQueue

		if curr.steps%2 == 0 {
			result += 1
		}

		if curr.steps == 64 {
			continue
		}

		for _, dir := range dirs {
			nextDir := Point{x: curr.x + dir.x, y: curr.y + dir.y}
			if nextDir.x < 0 || nextDir.y < 0 || nextDir.y >= len(*grid) || nextDir.x >= len((*grid)[nextDir.y]) {
				continue
			}

			if (*grid)[nextDir.y][nextDir.x] == "#" {
				continue
			}

			if slices.Contains(seen, nextDir) {
				continue
			}

			seen = append(seen, nextDir)
			queue = append(queue, GardenPlot{x: nextDir.x, y: nextDir.y, steps: curr.steps + 1})
		}
	}

	return result
}

func dequeue[K interface{}](queue []K) (K, []K) {
	element := queue[0]
	if len(queue) == 1 {
		return element, []K{}
	} else {
		return element, queue[1:]
	}
}
