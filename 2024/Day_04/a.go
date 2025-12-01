package Day_04

import (
	"log"
	"os"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)
	lines := strings.Fields(input)

	total := 0

	dirs := []string{}
	dirs = append(dirs, lines...)
	dirs = append(dirs, transpose(lines)...)
	dirs = append(dirs, diagonals(lines)...)

	for _, l := range dirs {
		count := strings.Count(l, "XMAS")
		total += count

		countBackwards := strings.Count(l, "SAMX")
		total += countBackwards
	}

	log.Println(total)
}

func transpose(lines []string) []string {
	xl := len(lines[0])

	result := make([]string, xl)
	for _, str := range lines {
		for j, b := range str {
			result[j] += string(b)
		}
	}

	return result
}

func diagonals(lines []string) []string {
	result := []string{}

	for i := 0; i < len(lines[0]); i++ {
		x := i
		y := 0
		str := ""

		for x >= 0 {
			str += string(lines[y][x])
			x--
			y++
		}
		result = append(result, str)
	}

	for i := 1; i < len(lines[0]); i++ {
		x := i
		y := len(lines[0]) - 1
		str := ""

		for x < len(lines[0]) {
			str += string(lines[y][x])
			x++
			y--
		}
		result = append(result, str)
	}

	for i := 0; i < len(lines[0]); i++ {
		x := i
		y := len(lines[0]) - 1
		str := ""

		for x >= 0 {
			str += string(lines[y][x])
			x--
			y--
		}
		result = append(result, str)
	}

	for i := 1; i < len(lines[0]); i++ {
		x := i
		y := 0
		str := ""

		for x < len(lines[0]) {
			str += string(lines[y][x])
			x++
			y++
		}
		result = append(result, str)
	}

	return result
}
