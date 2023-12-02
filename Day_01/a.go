package Day_01

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func A() {
	file, err := os.Open("Day_01/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		text := scanner.Text()

		isFirstDigitAssigned := false
		firstDigit := 0
		lastDigit := 0

		for _, c := range text {
			if c >= '0' && c <= '9' {
				if !isFirstDigitAssigned {
					isFirstDigitAssigned = true
					firstDigit = int(c - '0')
					lastDigit = int(c - '0')
				} else {
					lastDigit = int(c - '0')
				}
			}
		}

		result += firstDigit*10 + lastDigit
	}

	fmt.Print(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
