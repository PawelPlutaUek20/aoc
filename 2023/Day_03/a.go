package Day_03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func A() {
	file, _ := os.Open("Day_03/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	result := 0

	for y, line := range lines {
		number := ""
		for x, char := range line {
			isDigit := unicode.IsDigit(char)
			nextIsDigit := x+1 < len(line) && unicode.IsDigit(rune(lines[y][x+1]))

			if isDigit && nextIsDigit {
				number = number + string(char)
			} else if isDigit && !nextIsDigit {
				number = number + string(char)

				isPartNumber := false
				f, l := x-len(number)+1, x

				// check top
				if y-1 >= 0 {
					s := lines[y-1][f : l+1]
					for _, r := range s {
						if r != '.' {
							isPartNumber = true
						}
					}
				}

				// check bottom
				if y+1 < len(lines) {
					s := lines[y+1][f : l+1]
					for _, r := range s {
						if r != '.' {
							isPartNumber = true
						}
					}
				}

				// check top_left
				if y-1 >= 0 && f-1 >= 0 {
					r := rune(lines[y-1][f-1])
					if r != '.' {
						isPartNumber = true
					}
				}
				// check left
				if f-1 >= 0 {
					r := rune(lines[y][f-1])
					if r != '.' {
						isPartNumber = true
					}
				}
				// check bottom_left
				if y+1 < len(lines) && f-1 >= 0 {
					r := rune(lines[y+1][f-1])
					if r != '.' {
						isPartNumber = true
					}
				}
				// check top_right
				if y-1 >= 0 && l+1 < len(line) {
					r := rune(lines[y-1][l+1])
					if r != '.' {
						isPartNumber = true
					}
				}
				// check right
				if l+1 < len(line) {
					r := rune(lines[y][l+1])
					if r != '.' {
						isPartNumber = true
					}
				}
				// check bottom_right
				if y+1 < len(lines) && l+1 < len(line) {
					r := rune(lines[y+1][l+1])
					if r != '.' {
						isPartNumber = true
					}
				}

				if isPartNumber {
					partNumber, _ := strconv.Atoi(number)
					result += partNumber
				}
				number = ""
			} else {
				number = ""
			}
		}
	}

	fmt.Print(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
