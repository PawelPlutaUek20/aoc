package Day_04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func B() {
	file, _ := os.Open("Day_04/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	scratchCards := make(map[int]int)

	for scanner.Scan() {
		text := scanner.Text()

		card := strings.Split(text, ":")[0]
		numbers := strings.Split(text, ":")[1]

		winningNumbersString := strings.Split(numbers, "|")[0]
		myNumbersString := strings.Split(numbers, "|")[1]

		r := regexp.MustCompile("[0-9]+")
		cardNumberString := r.FindAllString(card, -1)[0]
		winningNumbersStringSlice := r.FindAllString(winningNumbersString, -1)
		myNumbersStringSlice := r.FindAllString(myNumbersString, -1)

		cardNumber, _ := strconv.Atoi(cardNumberString)

		numberofMatchingNumbers := 0
		for _, myNumber := range myNumbersStringSlice {
			if slices.Contains(winningNumbersStringSlice, myNumber) {
				numberofMatchingNumbers += 1
			}
		}

		scratchCards[cardNumber] += 1
		for n := 1; n <= numberofMatchingNumbers; n++ {
			scratchCards[cardNumber+n] += scratchCards[cardNumber]
		}

		result += scratchCards[cardNumber]
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
