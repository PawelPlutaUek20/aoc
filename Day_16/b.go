package Day_16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func B() {
	file, _ := os.Open("Day_16/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, strings.Split(text, ""))
	}

	result := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			seen := []Seen{}
			initialPoint := Point{x: j, y: i}
			initialDirs := getInitialDirs(&lines, initialPoint.x, initialPoint.y)
			for _, initialDir := range initialDirs {
				// probably should cache the points length from point A to point B
				walk(&lines, &seen, initialPoint, initialDir)

				uniqPoints := make(map[Point]bool, 0)
				for _, point := range seen {
					uniqPoints[Point{x: point.x, y: point.y}] = true
				}

				if len(uniqPoints) > result {
					result = len(uniqPoints)
				}

				seen = []Seen{}
			}
		}
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type initialPosition struct {
	x   int
	y   int
	dir int
}

func getInitialDirs(lines *[][]string, x int, y int) []int {
	maxX := len((*lines)[0]) - 1
	maxY := len(*lines) - 1
	if x == 0 && y == 0 {
		return []int{south, east}
	} else if x == maxX && y == 0 {
		return []int{south, west}
	} else if x == 0 && y == maxY {
		return []int{north, east}
	} else if x == maxX && y == maxY {
		return []int{north, west}
	} else if x == 0 {
		return []int{east}
	} else if x == maxX {
		return []int{west}
	} else if y == 0 {
		return []int{south}
	} else if y == maxY {
		return []int{north}
	} else {
		return []int{}
	}
}
