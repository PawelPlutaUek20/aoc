package Day_18

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func B() {
	file, err := os.ReadFile("Day_18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]rune, 71)
	for row := range grid {
		grid[row] = make([]rune, 71)
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		coordinates := strings.Split(strings.TrimSpace(line), ",")

		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])

		grid[y][x] = '#'

		if i >= 1024 {
			steps := bfs(grid, Point{0, 0}, Point{70, 70})
			if steps == -1 {
				log.Printf("%d,%d\n", x, y)
				break
			}
		}
	}
}
