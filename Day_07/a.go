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

type HandType uint8

const (
	highCard HandType = iota
	onePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func A() {
	file, _ := os.Open("Day_07/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cardStrengths := map[string]int{
		"2": 0,
		"3": 1,
		"4": 2,
		"5": 3,
		"6": 4,
		"7": 5,
		"8": 6,
		"9": 7,
		"T": 8,
		"J": 9,
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

		result := compareHandTypes(hand1, hand2, getHandType)

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

func getHandType(hand string) HandType {
	cards := strings.Split(hand, "")
	cardCount := make(map[string]int)

	for _, card := range cards {
		cardCount[card] += 1
	}

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

func compareHandStrengths(hand1 string, hand2 string, cardStrength map[string]int) int {

	cards2 := strings.Split(hand2, "")
	cards1 := strings.Split(hand1, "")

	strongerHand := 0
	for i, card := range cards1 {
		if cardStrength[card] > cardStrength[cards2[i]] {
			strongerHand = 1
			break
		} else if cardStrength[card] < cardStrength[cards2[i]] {
			strongerHand = -1
			break
		} else {
			// do nothing
		}
	}

	return strongerHand
}

func compareHandTypes(hand1 string, hand2 string, getHandType func(hand string) HandType) int {
	hand1Type := getHandType(hand1)
	hand2Type := getHandType(hand2)

	return int(hand1Type) - int(hand2Type)
}
