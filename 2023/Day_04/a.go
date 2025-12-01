package Day_04

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func A() {

	file, _ := os.Open("Day_04/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		text := scanner.Text()

		cardNumbers := strings.Split(text, ": ")[1]

		winningNumebersString := strings.Split(cardNumbers, " | ")[0]
		myNumbersString := strings.Split(cardNumbers, " | ")[1]

		myNumbers := strings.Split(myNumbersString, " ")
		winningNumbers := strings.Split(winningNumebersString, " ")

		myWinningNumbersCount := 0

		for _, myNumber := range myNumbers {
			for _, winningNumber := range winningNumbers {
				if winningNumber == myNumber && myNumber != "" {
					myWinningNumbersCount += 1
				}
			}
		}

		if myWinningNumbersCount > 0 {
			result += int(math.Pow(2, float64(myWinningNumbersCount-1)))
		}
	}

	fmt.Print(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
