package Day_16

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func B() {
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

	result := dijkstra(grid, start, end)
	fmt.Println("result", result)
}

type PointWithDir struct {
	point Point
	dir int
}

type QueueItem struct {
	item PointWithDir
	cost int
}

func dijkstra(maze [][]string, start Point, end Point) int {
	startItem := PointWithDir{start, 1}

	dist := make(map[PointWithDir]int)
	prev := make(map[PointWithDir][]PointWithDir)
	ends := make([]PointWithDir, 0)

	q := make([]QueueItem, 0)
	q = append(q, QueueItem{startItem, 0})

	dist[startItem] = 0
	best := math.MaxInt

	for len(q) > 0 {
		lowestUnvisited := getLowestUnvisitedB(q)

		item := q[lowestUnvisited].item
		cost := q[lowestUnvisited].cost
		q = append(q[:lowestUnvisited], q[lowestUnvisited+1:]...)

		currBestCost, hasBestCost := dist[item]
		if hasBestCost && cost > currBestCost {
			continue
		}

		if item.point == end {
			if cost > best {
				break
			}
			best = cost
			ends = append(ends, item)
		}

		nexts := []QueueItem{
			{PointWithDir{Point{item.point.row + directions[item.dir].row, item.point.col + directions[item.dir].col}, item.dir}, cost + 1},
			{PointWithDir{item.point, (item.dir + 1) % 4}, cost + 1000},
			{PointWithDir{item.point, (item.dir + 3) % 4}, cost + 1000},
		}

		for _, next := range nexts {
			if maze[next.item.point.row][next.item.point.col] == "#" {
				continue
			}
			currBestCost, hasBestCost := dist[next.item]

			if hasBestCost && next.cost > currBestCost {
				continue
			}
			if (hasBestCost && next.cost == currBestCost) {
				prev[next.item] = append(prev[next.item], item)
			}
			if !hasBestCost || (hasBestCost && next.cost < currBestCost) {
				prev[next.item] = []PointWithDir{item}
				dist[next.item] = next.cost
			}
			q = append(q, next)
		}
	}

	q2 := ends
	seen := make(map[PointWithDir]bool)
	for _, end := range q2 {
		seen[end] = true
	}

	for len(q2) > 0 {
		key := q2[0]
		q2 = q2[1:]

		for _, last := range prev[key] {
			if seen[last] {
				continue
			}
			seen[last] = true
			q2 = append(q2, last)
		}
	}

	result := make(map[Point]bool)
	for pointWithDir := range seen {
		result[pointWithDir.point] = true
	}

	return len(result)
}

func getLowestUnvisitedB(q []QueueItem) int {
	idx := -1
	minCost := math.MaxInt

	for i, item := range q {
		if item.cost < minCost {
			idx = i
			minCost = item.cost
		}
	}

	return idx
}