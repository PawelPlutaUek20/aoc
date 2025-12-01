package Day_01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func B() {
	file, err := os.Open("Day_01/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	options := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	result := 0
	for scanner.Scan() {
		text := scanner.Text()

		firstDigitIndex := -1
		lastDigitIndex := -1

		firstDigit := 0
		lastDigit := 0

		for k, v := range options {

			f := strings.Index(text, k)
			l := strings.LastIndex(text, k)

			isNewFirst := (firstDigitIndex == -1 && f != -1) || (f != -1 && f < firstDigitIndex)
			isNewLast := (lastDigitIndex == -1 && l != -1) || (l != -1 && l > lastDigitIndex)

			if isNewFirst {
				firstDigitIndex = f
				firstDigit = v
			}

			if isNewLast {
				lastDigitIndex = l
				lastDigit = v
			}
		}

		result += firstDigit*10 + lastDigit
	}

	fmt.Print(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
