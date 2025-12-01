package Day_25

import (
	"fmt"
	"os"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_25/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(file, "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	keys := make([][]string, 0)
	locks := make([][]string, 0)

	for _, line := range lines {
		maybeLockOrKey := strings.Split(line, "\n")
		if isLock(maybeLockOrKey) {
			locks = append(locks, maybeLockOrKey)
		} else if isKey(maybeLockOrKey) {
			keys = append(keys, maybeLockOrKey)
		} else {
			panic("not a lock and not a key")
		}
	}

	fit := 0
	for _, lock := range locks {
		for _, key := range keys {
			fits := checkFit(lock, key)

			if fits {
				fit++
			}
		}
	}

	fmt.Println(fit)
}

func isLock(schematic []string) bool {
	return schematic[0] == "#####" && schematic[len(schematic)-1] == "....."
}

func isKey(schematic []string) bool {
	return schematic[0] == "....." && schematic[len(schematic)-1] == "#####"
}

func getColumnHeights(schematic []string) []int {
	result := make([]int, 5)

	for _, columns := range schematic[1 : len(schematic)-1] {
		for column, item := range columns {
			if item == '#' {
				result[column]++
			}
		}
	}

	return result
}

func checkFit(lock []string, key []string) bool {
	keyHeights := getColumnHeights(key)
	lockHeights := getColumnHeights(lock)

	fits := true

	for j, keyHeight := range keyHeights {
		lockHeight := lockHeights[j]
		if keyHeight+lockHeight > 5 {
			fits = false
		}
	}

	return fits
}