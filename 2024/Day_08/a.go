package Day_08

import (
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func A() {
	file, err := os.ReadFile("Day_08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	antinodes := map[Point]bool{}
	antennas := map[rune][]Point{}

	for row, line := range lines {
		for col, antenna := range line {
			if antenna != '.' {
				if _, exists := antennas[antenna]; !exists {
					antennas[antenna] = []Point{}
				}

				antennas[antenna] = append(antennas[antenna], Point{Y: row, X: col})
			}
		}
	}

	for _, positions := range antennas {
		antennasAntinodes := FindAllAntinodes(positions)
		for _, antinode := range antennasAntinodes {
			if antinode.X >= 0 && antinode.Y >= 0 && antinode.X < len(lines[0]) && antinode.Y < len(lines) {
				antinodes[antinode] = true
			}
		}
	}

	log.Println(len(antinodes))
}

func FindAllAntinodes(antennas []Point) []Point {
	antinodes := []Point{}
	antennasPairs := [][2]Point{}

	for i := 0; i < len(antennas)-1; i++ {
		antenna1 := antennas[i]
		for j := i + 1; j < len(antennas); j++ {
			antenna2 := antennas[j]
			newPair := [2]Point{antenna1, antenna2}
			antennasPairs = append(antennasPairs, newPair)
		}
	}

	for _, antennaPair := range antennasPairs {
		point1 := antennaPair[0]
		point2 := antennaPair[1]

		xDelta := point2.X - point1.X
		yDelta := point2.Y - point1.Y

		antinodes = append(
			antinodes,
			Point{X: point1.X - xDelta, Y: point1.Y - yDelta},
			Point{X: point2.X + xDelta, Y: point2.Y + yDelta},
		)
	}

	return antinodes
}
