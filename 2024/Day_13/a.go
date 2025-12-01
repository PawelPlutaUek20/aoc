package Day_13

import (
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	sum := float64(0)
	for _, line := range lines {
		machineButtonBehavior := strings.Split(strings.TrimSpace(line), "\n")

		buttonA := parseData(machineButtonBehavior[0])
		buttonB := parseData(machineButtonBehavior[1])
		prize := parseData(machineButtonBehavior[2])

		a, b := solveEquation(buttonA, buttonB, prize)
		if math.Trunc(a) == a && math.Trunc(b) == b {
			sum += 3*a + b
		}
	}

	log.Println(sum)
}

type ButtonBehavior struct {
	X, Y float64
}

func parseData(behavior string) ButtonBehavior {
	r := regexp.MustCompile("[0-9]+")
	data := r.FindAllString(behavior, -1)

	x, _ := strconv.Atoi(data[0])
	y, _ := strconv.Atoi(data[1])

	return ButtonBehavior{float64(x), float64(y)}
}

// buttonA.Y + buttonB.Y = result.Y
// buttonA.X + buttonB.X = result.X
func solveEquation(buttonA, buttonB, result ButtonBehavior) (float64, float64) {
	det := buttonA.Y*buttonB.X - buttonA.X*buttonB.Y

	x := (result.Y*buttonB.X - result.X*buttonB.Y) / det
	y := (buttonA.Y*result.X - buttonA.X*result.Y) / det

	return x, y
}
