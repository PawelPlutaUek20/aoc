package Day_11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const (
	galaxy     = "#"
	emptySpace = "."
)

type Point struct {
	x int
	y int
}

func A() {

	file, _ := os.Open("Day_11/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	rowsWithoutGalaxies := getRowsWithoutGalaxies(lines)
	colsWithoutGalaxies := getRowsWithoutGalaxies(transpose(lines))
	galaxiesPositions := getGalaxiesPositions(lines)

	pairs := make([][2]Point, 0)
	for i := 0; i < len(galaxiesPositions)-1; i++ {
		for j := i; j < len(galaxiesPositions)-1; j++ {
			pairs = append(pairs, [2]Point{galaxiesPositions[i], galaxiesPositions[j+1]})
		}
	}

	result := 0
	for _, pair := range pairs {
		point1 := pair[0]
		point2 := pair[1]

		lX := min(point1.x, point2.x)
		hX := max(point1.x, point2.x)

		lY := min(point1.y, point2.y)
		hY := max(point1.y, point2.y)

		distance := hX - lX + hY - lY

		for i := lX + 1; i < hX; i++ {
			if slices.Contains(colsWithoutGalaxies, i) {
				distance += 1
			}
		}

		for i := lY + 1; i < hY; i++ {
			if slices.Contains(rowsWithoutGalaxies, i) {
				distance += 1
			}
		}

		result += distance
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getGalaxiesPositions(input [][]string) []Point {
	result := make([]Point, 0)
	for y, col := range input {
		for x, row := range col {
			if row == galaxy {
				result = append(result, Point{x, y})
			}
		}
	}
	return result
}

func getRowsWithoutGalaxies(input [][]string) []int {
	rowsWithoutGalaxies := make([]int, 0)
	for i, col := range input {
		if !slices.Contains(col, galaxy) {
			rowsWithoutGalaxies = append(rowsWithoutGalaxies, i)
		}
	}
	return rowsWithoutGalaxies
}

func transpose(a [][]string) [][]string {
	newArr := make([][]string, len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
