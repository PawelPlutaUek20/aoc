package Day_11

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Fields(input)

	for i := 0; i < 25; i++ {
		newLines := make([]string, 0, len(lines[0]))
		for _, line := range lines {
			if line == "0" {
				newLines = append(newLines, "1")
			} else if len(line)%2 == 0 {
				leftStone := line[0 : len(line)/2]
				rightStone := line[len(line)/2:]

				leftStoneInt, _ := strconv.Atoi(leftStone)
				rightStoneInt, _ := strconv.Atoi(rightStone)

				newLines = append(newLines, strconv.Itoa(leftStoneInt))
				newLines = append(newLines, strconv.Itoa(rightStoneInt))
			} else {
				stoneInt, _ := strconv.Atoi(line)
				newLines = append(newLines, strconv.Itoa(stoneInt*2024))
			}
		}
		lines = newLines
	}

	log.Println(len(lines))
}
