package Day_06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func A() {

	file, _ := os.Open("Day_06/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	results := make(map[int]int)

	r := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	times := mapStringToInt(r.FindAllString(lines[0], -1))
	distances := mapStringToInt(r.FindAllString(lines[1], -1))

	for index, time := range times {
		for i := 0; i < time; i++ {
			if (time-i)*i > distances[index] {
				results[index] += 1
			}
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

func mapStringToInt(arr []string) []int {
	arr2 := make([]int, 0)

	for _, v := range arr {
		num, _ := strconv.Atoi(v)
		arr2 = append(arr2, num)
	}

	return arr2
}
