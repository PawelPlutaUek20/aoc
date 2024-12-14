package Day_14

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type SafetyFactor struct {
	Second int
	Factor int
}

func B() {
	file, err := os.ReadFile("Day_14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	space := Space{101, 103}

	seconds := 10000
	robots := parseLines(lines)

	bestSafetyFactor := SafetyFactor{
		Second: 0,
		Factor: getSafetyFactor(space, robots),
	}

	for second := 1; second < seconds; second++ {
		for i := range robots {
			robot := &robots[i]
			robot.Position.X = mod(robot.Position.X+robot.Velocity.X, space.Width)
			robot.Position.Y = mod(robot.Position.Y+robot.Velocity.Y, space.Height)
		}

		factor := getSafetyFactor(space, robots)
		if factor < bestSafetyFactor.Factor {
			bestSafetyFactor.Factor = factor
			bestSafetyFactor.Second = second
		}
	}

	robots = parseLines(lines)

	for second := 1; second <= bestSafetyFactor.Second; second++ {
		for i := range robots {
			robot := &robots[i]
			robot.Position.X = mod(robot.Position.X+robot.Velocity.X, space.Width)
			robot.Position.Y = mod(robot.Position.Y+robot.Velocity.Y, space.Height)
		}

		if second == bestSafetyFactor.Second {
			space.Print(robots)
		}
	}

	log.Println(bestSafetyFactor.Second)
}

func (space *Space) Print(robots []Robot) {
	height := int(space.Height)
	width := int(space.Width)

	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]string, width)
		for j := 0; j < width; j++ {
			grid[i][j] = "."
		}
	}

	for _, robot := range robots {
		grid[int(robot.Position.Y)][int(robot.Position.X)] = "#"
	}

	for _, row := range grid {
		fmt.Println(row)
	}
}
