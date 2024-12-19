package Day_19

import (
	"log"
	"os"
	"strings"
)

func B() {
	file, err := os.ReadFile("Day_19/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	availableTowelPatters := strings.Split(strings.TrimSpace(lines[0]), ", ")
	designs := strings.Split(strings.TrimSpace(lines[1]), "\n")

	sum := possibleDesigns(availableTowelPatters, designs)
	log.Println(sum)
}

func possibleDesigns(patterns []string, designs []string) int {
	cache := make(map[string]int)

	var possibilities func(design string) int

	possibilities = func(design string) int {
		if design == "" {
			return 1
		}

		cached, isCached := cache[design]
		if isCached {
			return cached
		}

		sum := 0
		for _, pattern := range patterns {
			after, found := strings.CutPrefix(design, pattern)
			if found {
				sum += possibilities(after)
			}
		}

		cache[design] = sum
		return sum
	}

	sum := 0
	for _, design := range designs {
		sum += possibilities(design)
	}

	return sum
}
