package Day_02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func A() {

	file, _ := os.Open("Day_02/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	conditions := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for scanner.Scan() {
		text := scanner.Text()

		gameId := getGameId(text)
		isValidGame := getIsValidGame(text, conditions)

		if isValidGame {
			result += gameId
		}
	}

	fmt.Print(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getIsValidGame(text string, conditions map[string]int) bool {
	games := strings.Split(text, ": ")[1]
	setsOfCubes := strings.Split(games, "; ")

	r := regexp.MustCompile("([0-9]+) (red|green|blue)")

	isValidGame := true
	for _, set := range setsOfCubes {
		m := r.FindAllStringSubmatch(set, -1)

		for _, parsedSet := range m {
			number, _ := strconv.Atoi(parsedSet[1])
			color := parsedSet[2]

			if number > conditions[color] {
				isValidGame = false
			}
		}
	}

	return isValidGame
}

func getGameId(text string) int {
	r := regexp.MustCompile("Game ([0-9]+):")
	m := r.FindStringSubmatch(text)[1]

	gameId, _ := strconv.Atoi(m)

	return gameId
}
