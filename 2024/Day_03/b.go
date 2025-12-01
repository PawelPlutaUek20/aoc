package Day_03

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func B() {
	file, err := os.ReadFile("Day_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)

	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)|don't\(\)|do\(\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	sum := 0
	disabled := false

	for _, match := range matches {
		if match[0] == "do()" {
			disabled = false
		} else if match[0] == "don't()" {
			disabled = true
		} else {
			if !disabled {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])

				sum += num1 * num2
			}
		}

	}

	log.Println(sum)
}
