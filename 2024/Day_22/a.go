package Day_22

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_22/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	data := make([]int, len(lines))
	for i, line := range lines {
		lineInt, _ := strconv.Atoi(line)
		data[i] = lineInt
	}

	result := 0
	for _, secretNumber := range data {
		result += evolveSecretNumberN(secretNumber, 2000)
	}
	fmt.Println(result)
}

func mix(value int, secretNumber int) int {
	return value ^ secretNumber
}

func prune(secretNumber int) int {
	return trueMod(secretNumber, 16777216)
}

func trueMod(x, n int) int {
	return ((x % n) + n) % n
}

func evolveSecretNumber(secretNumber int) int {
	nextSecretNumber := secretNumber

	// Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
	nextSecretNumber = prune(mix(nextSecretNumber*64, nextSecretNumber))

	// Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
	nextSecretNumber = prune(mix(nextSecretNumber/32, nextSecretNumber))

	// Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.
	nextSecretNumber = prune(mix(nextSecretNumber*2048, nextSecretNumber))

	return nextSecretNumber
}

func evolveSecretNumberN(secretNumber int, n int) int {
	result := secretNumber
	for i := 0; i < n; i++ {
		result = evolveSecretNumber(result)
	}
	return result
}
