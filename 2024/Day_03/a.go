package Day_03

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func A() {
	file, err := os.ReadFile("Day_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)

	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		sum += num1 * num2
	}

	log.Println(sum)
}
