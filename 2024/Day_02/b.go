package Day_02

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func B() {
	file, _ := os.Open("Day_02/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeLevels := 0

	for scanner.Scan() {
		text := scanner.Text()
		levels := strings.Fields(text)

		for i := range levels {
			data := make([]string, len(levels)-1)
			copy(data[:i], levels[:i])
			copy(data[i:], levels[i+1:])

			isSafe := checkSafe(data)
			if isSafe {
				safeLevels++
				break
			}
		}
	}

	log.Println(safeLevels)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkSafe(levels []string) bool {
	allIncreasingOrDecresing := true
	var maxDiff float64 = 0
	var levelDir int

	for levelIndex := 0; levelIndex < len(levels)-1; levelIndex++ {
		levelA, _ := strconv.Atoi(levels[levelIndex])
		levelB, _ := strconv.Atoi(levels[levelIndex+1])

		currDiff := float64(levelA - levelB)
		maxDiff = math.Max(maxDiff, math.Abs(currDiff))

		var currDir int
		if currDiff > 0 {
			currDir = 1
		} else if currDiff < 0 {
			currDir = -1
		} else {
			currDir = 0
		}

		if currDir == 0 {
			allIncreasingOrDecresing = false
			break
		}

		if levelIndex == 0 {
			levelDir = currDir
			continue
		}

		if levelDir != currDir {
			allIncreasingOrDecresing = false
			break
		}
	}

	isSafe := allIncreasingOrDecresing && (maxDiff <= 3 && maxDiff >= 1)
	return isSafe
}
