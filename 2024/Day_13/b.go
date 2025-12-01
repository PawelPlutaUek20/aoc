package Day_13

import (
	"log"
	"math"
	"os"
	"strings"
)

func B() {
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
		prize.X += 10000000000000
		prize.Y += 10000000000000

		a, b := solveEquation(buttonA, buttonB, prize)
		if math.Trunc(a) == a && math.Trunc(b) == b {
			sum += 3*a + b
		}
	}

	log.Printf("%f", sum)
}
