package Day_05

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func B() {
	file, err := os.ReadFile("Day_05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	output := 0

	data := strings.Split(input, "\n\n")
	orderRules := strings.Fields(data[0])
	updates := strings.Fields(data[1])

	for _, update := range updates {
		updatePagesOriginal := strings.Split(update, ",")
		updatePages := slices.Clone(updatePagesOriginal)

		slices.SortFunc(updatePages, func(a, b string) int {
			for _, orderRule := range orderRules {
				rule := strings.Split(orderRule, "|")

				if a == rule[0] && b == rule[1] {
					return -1
				}

				if b == rule[0] && a == rule[1] {
					return 1
				}

			}

			return 0
		})

		if !slices.Equal(updatePages, updatePagesOriginal) {
			midIndex := len(updatePages) / 2
			midElement := updatePages[midIndex]

			midElementInt, _ := strconv.Atoi(midElement)
			output += midElementInt
		}

	}

	log.Println(output)
}
