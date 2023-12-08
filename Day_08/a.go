package Day_08

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type leftAndRight struct {
	left  string
	right string
}

func A() {

	file, _ := os.Open("Day_08/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r := regexp.MustCompile("([A-Z]+) = \\(([A-Z]+), ([A-Z]+)\\)")

	lines := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}

	instructions := strings.Split(lines[0], "")
	other := make(map[string]leftAndRight, 0)

	for i := 2; i < len(lines); i++ {
		dupa := lines[i]
		dupa2 := r.FindStringSubmatch(dupa)

		element := dupa2[1]
		left := dupa2[2]
		right := dupa2[3]

		other[element] = leftAndRight{left, right}
	}

	i := 0
	currElement := "AAA"

	for {
		if currElement == "ZZZ" {
			break
		}

		index := i % len(instructions)
		instruction := instructions[index]

		if instruction == "L" {
			currElement = other[currElement].left
		} else {
			currElement = other[currElement].right
		}

		i = i + 1
	}

	fmt.Print(i)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
