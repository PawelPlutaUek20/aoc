package Day_14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func B() {

	file, _ := os.Open("Day_14/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, strings.Split(text, ""))
	}

	for i := 0; i < 1000; i++ {
		tempLines := slices.Clone(lines)
		for i := 0; i < 4; i++ {
			tilt(&lines, i)
			tilt(&tempLines, i)
		}
	}

	result := 0
	for i, col := range lines {
		fmt.Println(col)
		for _, row := range col {
			result += (len(col) - i) * strings.Count(row, "O")
		}
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Position struct {
	x int
	y int
}

func tilt(lines *[][]string, dir int) {
	switch dir {
	case 0:
		tiltNorth(lines)
	case 1:
		tiltWest(lines)
	case 2:
		tiltSouth(lines)
	case 3:
		tiltEast(lines)
	}
}

func tiltNorth(lines *[][]string) {
	for y, col := range *lines {
		for x, row := range col {
			if row == "O" {
				tempY := y - 1
				for tempY >= 0 && (*lines)[tempY][x] == "." {
					(*lines)[tempY][x] = "O"
					(*lines)[tempY+1][x] = "."
					tempY -= 1
				}
			}
		}
	}
}

func tiltWest(lines *[][]string) {
	for y, col := range *lines {
		for x, row := range col {
			if row == "O" {
				tempX := x - 1
				for tempX >= 0 && (*lines)[y][tempX] == "." {
					(*lines)[y][tempX] = "O"
					(*lines)[y][tempX+1] = "."
					tempX -= 1
				}
			}
		}
	}
}

func tiltSouth(lines *[][]string) {
	for y := (len(*lines) - 1); y >= 0; y-- {
		for x := 0; x < len((*lines)[0]); x++ {
			row := (*lines)[y][x]
			if row == "O" {
				tempY := y + 1
				for tempY >= 0 && tempY < len(*lines) && (*lines)[tempY][x] == "." {
					(*lines)[tempY][x] = "O"
					(*lines)[tempY-1][x] = "."
					tempY += 1
				}
			}
		}
	}
}

func tiltEast(lines *[][]string) {
	for y := 0; y < len(*lines); y++ {
		for x := len((*lines)[0]) - 1; x >= 0; x-- {
			row := (*lines)[y][x]
			if row == "O" {
				tempX := x + 1
				for tempX >= 0 && tempX < len((*lines)[0]) && (*lines)[y][tempX] == "." {
					(*lines)[y][tempX] = "O"
					(*lines)[y][tempX-1] = "."
					tempX += 1
				}
			}
		}
	}
}
