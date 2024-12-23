package Day_23

import (
	"fmt"
	"os"
	"strings"
	"slices"
	"sort"
)

func A() {
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

	connectedSets := make(map[[3]string]bool)

	for computer1, connections1 := range sets {
		for _, computer2 := range connections1 {
			connections2 := sets[computer2]
			for _, computer3 := range connections2 {
				contains := slices.Contains(connections1, computer3)
				if computer1 != computer3 && contains {
					connStr := []string{computer1, computer2, computer3}
					sort.Strings(connStr)
					sortedConnStr := [3]string{connStr[0], connStr[1], connStr[2]}
					connectedSets[sortedConnStr] = true
				}
			}
		}
	}

	sum := 0
	for connectionSet := range connectedSets {
		if connectionSet[0][0] == 't' || connectionSet[1][0] == 't' || connectionSet[2][0] == 't' {
			sum++
		}
	}

	fmt.Println(sum)
}
