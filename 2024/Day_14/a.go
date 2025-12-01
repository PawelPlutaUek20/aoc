package Day_14

import (
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	X, Y float64
}

type Velocity struct {
	X, Y float64
}

type Robot struct {
	Position Position
	Velocity Velocity
}

type Space struct {
	Width  float64
	Height float64
}

func mod(a, b float64) float64 {
	return math.Mod(math.Mod(a, b)+b, b)
}

func A() {

	file, err := os.ReadFile("Day_14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	space := Space{101, 103}
	seconds := 100
	robots := parseLines(lines)

	for i := range robots {
		for second := 0; second < seconds; second++ {
			robot := &robots[i]
			robot.Position.X = mod(robot.Position.X+robot.Velocity.X, space.Width)
			robot.Position.Y = mod(robot.Position.Y+robot.Velocity.Y, space.Height)
		}
	}

	factor := getSafetyFactor(space, robots)
	log.Println(factor)
}

func parseLines(lines []string) []Robot {
	result := make([]Robot, 0, len(lines))

	re := regexp.MustCompile("-?[0-9]+")

	for _, line := range lines {
		data := re.FindAllString(line, -1)

		positionX, _ := strconv.Atoi(data[0])
		positionY, _ := strconv.Atoi(data[1])
		velocityX, _ := strconv.Atoi(data[2])
		velocityY, _ := strconv.Atoi(data[3])

		robot := Robot{
			Position: Position{float64(positionX), float64(positionY)},
			Velocity: Velocity{float64(velocityX), float64(velocityY)},
		}

		result = append(result, robot)
	}

	return result
}

func getSafetyFactor(space Space, robots []Robot) int {
	quadrant1 := 0
	quadrant2 := 0
	quadrant3 := 0
	quadrant4 := 0

	middleWidth := (math.Floor(space.Width / 2))
	middlwHeight := (math.Floor(space.Height / 2))

	for _, robot := range robots {
		if robot.Position.X < middleWidth && robot.Position.Y < middlwHeight {
			quadrant1++
		} else if robot.Position.X > middleWidth && robot.Position.Y < middlwHeight {
			quadrant2++
		} else if robot.Position.X < middleWidth && robot.Position.Y > middlwHeight {
			quadrant3++
		} else if robot.Position.X > middleWidth && robot.Position.Y > middlwHeight {
			quadrant4++
		}
	}

	return quadrant1 * quadrant2 * quadrant3 * quadrant4
}
