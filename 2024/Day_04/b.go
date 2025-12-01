package Day_04

import (
	"log"
	"os"
	"strings"
)

func B() {
	file, err := os.ReadFile("Day_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)
	lines := strings.Fields(input)

	total := 0

	for i, col := range lines {
		for j, row := range col {

			if row != 'A' {
				continue
			}

			if i-1 < 0 || j-1 < 0 || i+1 >= len(lines) || j+1 >= len(lines[0]) {
				continue
			}

			if lines[i-1][j-1] == 'M' && lines[i-1][j+1] == 'S' && lines[i+1][j+1] == 'S' && lines[i+1][j-1] == 'M' {
				total++
			} else if lines[i-1][j-1] == 'M' && lines[i-1][j+1] == 'M' && lines[i+1][j+1] == 'S' && lines[i+1][j-1] == 'S' {
				total++
			} else if lines[i-1][j-1] == 'S' && lines[i-1][j+1] == 'S' && lines[i+1][j+1] == 'M' && lines[i+1][j-1] == 'M' {
				total++
			} else if lines[i-1][j-1] == 'S' && lines[i-1][j+1] == 'M' && lines[i+1][j+1] == 'M' && lines[i+1][j-1] == 'S' {
				total++
			}
		}
	}

	log.Println(total)
}
