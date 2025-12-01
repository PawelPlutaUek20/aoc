package Day_02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func B() {

	file, _ := os.Open("Day_02/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		text := scanner.Text()
		red, green, blue := getFewestNumberOfCubes(text)

		result += red * green * blue
	}

	fmt.Print(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getFewestNumberOfCubes(text string) (int, int, int) {
	games := strings.Split(text, ": ")[1]
	setsOfCubes := strings.Split(games, "; ")

	r := regexp.MustCompile("([0-9]+) (red|green|blue)")

	mins := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range setsOfCubes {
		m := r.FindAllStringSubmatch(set, -1)

		for _, parsedSet := range m {
			number, _ := strconv.Atoi(parsedSet[1])
			color := parsedSet[2]

			if mins[color] < number {
				mins[color] = number
			}
		}
	}

	return mins["red"], mins["green"], mins["blue"]
}
