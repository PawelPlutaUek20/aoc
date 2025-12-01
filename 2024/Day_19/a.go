package Day_19

import (
	"log"
	"os"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_19/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	availableTowelPatters := strings.Split(strings.TrimSpace(lines[0]), ", ")
	designs := strings.Split(strings.TrimSpace(lines[1]), "\n")

	sum := validDesigns(availableTowelPatters, designs)
	log.Println(sum)
}

func validDesigns(patterns []string, designs []string) int {
	cache := make(map[string]bool)

	var isValid func(design string) bool

	isValid = func(design string) bool {
		if design == "" {
			return true
		}

		cached, isCached := cache[design]
		if isCached {
			return cached
		}

		for _, pattern := range patterns {
			after, found := strings.CutPrefix(design, pattern)
			if found && isValid(after) {
				return true
			}
		}

		cache[design] = false
		return false
	}

	sum := 0
	for _, design := range designs {
		if isValid(design) {
			sum++
		}
	}

	return sum
}
