package Day_06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func B() {

	file, _ := os.Open("Day_06/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	results := make(map[int]int)

	r := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	times := r.FindAllString(lines[0], -1)
	distances := r.FindAllString(lines[1], -1)

	timeStr := ""
	for _, t := range times {
		timeStr += t
	}

	distanceStr := ""
	for _, d := range distances {
		distanceStr += d
	}

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)

	for i := 0; i < time; i++ {
		if (time-i)*i > distance {
			results[0] += 1
		}
	}

	result := 1
	for _, res := range results {
		result *= res
	}

	fmt.Print(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
