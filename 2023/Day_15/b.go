package Day_15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

const (
	dash   = "-"
	equals = "="
)

func B() {

	file, _ := os.Open("Day_15/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("([a-z]+)([-=])([1-9]*)")

	lines := ""
	for scanner.Scan() {
		text := scanner.Text()
		lines += text
	}

	steps := strings.Split(lines, ",")
	boxes := make(map[int][]lens, 0)

	for _, step := range steps {
		match := re.FindStringSubmatch(step)
		label := match[1]
		operation := match[2]
		focalLength, _ := strconv.Atoi(match[3])

		box := hashAlgorithm(label, 0)
		idx := slices.IndexFunc(boxes[box], func(lens lens) bool { return lens.label == label })

		if operation == equals {
			if boxes[box] == nil {
				boxes[box] = []lens{}
			}
			if idx > -1 {
				slices.Replace(boxes[box], idx, idx+1, lens{label, focalLength})
			} else {
				boxes[box] = append(boxes[box], lens{label, focalLength})
			}
		} else if operation == dash {
			if idx > -1 {
				boxes[box] = slices.Delete(boxes[box], idx, idx+1)
			}
		}
	}

	result := 0
	for boxNr, box := range boxes {
		for i, lens := range box {
			result += (boxNr + 1) * (i + 1) * lens.focalLength
		}
	}
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
