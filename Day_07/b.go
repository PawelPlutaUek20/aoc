package Day_07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func B() {
	file, _ := os.Open("Day_07/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cardStrengths := map[string]int{
		"J": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
		"9": 8,
		"T": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	slices.SortFunc(lines, func(a string, b string) int {
		hand1 := strings.Split(a, " ")[0]
		hand2 := strings.Split(b, " ")[0]

		getHandTypeFunc := getHandTypeB(cardStrengths)
		result := compareHandTypes(hand1, hand2, getHandTypeFunc)

		if result == 0 {
			return compareHandStrengths(a, b, cardStrengths)
		} else {
			return result
		}
	})

	result := 0
	for rank, line := range lines {
		bid, _ := strconv.Atoi(strings.Split(line, " ")[1])
		result += bid * (rank + 1)
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getHandTypeB(s map[string]int) func(hand string) HandType {
	return func(hand string) HandType {

		cards := strings.Split(hand, "")

		cardCount := make(map[string]int)
		for _, card := range cards {
			cardCount[card] += 1
		}

		uniqCards := make([]string, 0)
		for card, _ := range cardCount {
			uniqCards = append(uniqCards, card)
		}

		slices.SortFunc(uniqCards, func(a string, b string) int {
			return cardCount[b] - cardCount[a]
		})

		jokerCardCount := cardCount["J"]

		if jokerCardCount == 5 {
			cardCount["A"] = 5
		} else if uniqCards[0] == "J" {
			cardCount[uniqCards[1]] += jokerCardCount
		} else {
			cardCount[uniqCards[0]] += jokerCardCount
		}

		delete(cardCount, "J")

		counts := make([]int, 0)
		for _, count := range cardCount {
			counts = append(counts, count)
		}

		slices.Sort(counts)

		if counts[0] == 5 {
			return FiveOfAKind
		} else if counts[1] == 4 {
			return FourOfAKind
		} else if counts[0] == 2 && counts[1] == 3 {
			return FullHouse
		} else if counts[0] == 1 && counts[1] == 1 && counts[2] == 3 {
			return ThreeOfAKind
		} else if counts[0] == 1 && counts[1] == 2 && counts[2] == 2 {
			return TwoPair
		} else if counts[0] == 1 && counts[1] == 1 && counts[2] == 1 && counts[3] == 2 {
			return onePair
		} else {
			return highCard
		}
	}
}
