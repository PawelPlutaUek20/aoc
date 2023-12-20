package Day_18

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	dir   string
	count int
	color string
}

type Point struct {
	x int
	y int
}

func A() {

	dirPointMap := map[string]Point{
		"U": {x: 0, y: -1},
		"R": {x: 1, y: 0},
		"D": {x: 0, y: 1},
		"L": {x: -1, y: 0},
	}

	file, _ := os.Open("Day_18/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("([A-Z]) ([0-9]+) (\\(.+\\))")

	lines := make([]Instruction, 0)
	for scanner.Scan() {
		text := scanner.Text()
		match := re.FindStringSubmatch(text)

		dir := match[1]
		count, _ := strconv.Atoi(match[2])
		color := match[3]

		lines = append(lines, Instruction{dir, count, color})
	}

	perimeter := 0
	currPoint := Point{x: 0, y: 0}
	xPts := []int{}
	yPts := []int{}

	for _, instruction := range lines {
		perimeter += instruction.count
		currDir := dirPointMap[instruction.dir]

		xNext := currPoint.x + currDir.x*instruction.count
		yNext := currPoint.y + currDir.y*instruction.count

		xPts = append(xPts, xNext)
		yPts = append(yPts, yNext)

		currPoint = Point{x: xNext, y: yNext}
	}

	result := 0

	for i := 0; i < len(xPts)-1; i++ {
		result += (xPts[i] + xPts[i+1]) * (yPts[i] - yPts[i+1])
	}

	fmt.Println(abs(result/2) + perimeter/2 + 1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
