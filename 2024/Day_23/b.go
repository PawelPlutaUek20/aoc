package Day_23

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func B() {
	file, err := os.ReadFile("Day_23/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sets := make(map[string][]string)

	for _, conn := range lines {
		computers := strings.Split(conn, "-")
		computer1 := computers[0]
		computer2 := computers[1]

		sets[computer1] = append(sets[computer1], computer2)
		sets[computer2] = append(sets[computer2], computer1)
	}

	
	maxClique := make(map[string]bool)

	for set := range sets {
		clique := findClique(set, sets)
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}

	result := make([]string, 0)
	for computer := range maxClique {
		result = append(result, computer)
	}
	slices.Sort(result)

	fmt.Println(strings.Join(result, ","))
}

func findClique(start string, sets map[string][]string) map[string]bool {
	cliques := make(map[string]bool)

	q := make([]string, 0)
	q = append(q, start)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if cliques[curr] {
			continue
		}

		isClique := true
		possibilities := sets[curr]

		for clique := range cliques {
			if !slices.Contains(possibilities, clique) {
				isClique = false
			}
		}

		if !isClique {
			continue
		}

		cliques[curr] = true
		for _, possibility := range possibilities {
			q = append(q, possibility)
		}
	}

	return cliques
}
