package Day_09

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func A() {
	file, _ := os.Open("Day_09/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r := regexp.MustCompile("-?[0-9]+")
	result := 0

	for scanner.Scan() {
		text := scanner.Text()

		sequence := make([]int, 0)

		for _, str := range r.FindAllString(text, -1) {
			num, _ := strconv.Atoi(str)
			sequence = append(sequence, num)
		}

		result += extrapolate(sequence)
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func extrapolate(sequence []int) int {
	if !(slices.ContainsFunc(sequence, func(a int) bool { return a != 0 })) {
		return 0
	}

	differences := findDifferences(sequence)
	return sequence[len(sequence)-1] + extrapolate(differences)
}

func findDifferences(sequence []int) []int {
	differences := make([]int, 0)

	for index := 0; index < len(sequence)-1; index++ {
		differences = append(differences, sequence[index+1]-sequence[index])
	}

	return differences
}
