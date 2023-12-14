package Day_14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func A() {

	file, _ := os.Open("Day_14/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, strings.Split(text, ""))
	}

	for y, col := range lines {
		for x, row := range col {
			if row == "O" {
				tempY := y - 1
				for tempY >= 0 && lines[tempY][x] == "." {
					lines[tempY][x] = "O"
					lines[tempY+1][x] = "."
					tempY -= 1
				}
			}
		}
	}

	result := 0
	for i, col := range lines {
		for _, row := range col {
			result += (len(col) - i) * strings.Count(row, "O")
		}
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
