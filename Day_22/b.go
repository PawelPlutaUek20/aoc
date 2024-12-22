package Day_22

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const iterations = 2000

func B() {
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

	priceChangesMapSum := make(map[[4]int]int)

	for _, secretNumber := range data {
		prices := getPrices(secretNumber)
		priceChanges := getPricesChangesMap(prices)
		for priceChange, price := range priceChanges {
			priceChangesMapSum[priceChange] += price
		}
	}

	var bestPrice int

	for _, price := range priceChangesMapSum {
		if price > bestPrice {
			bestPrice = price
		}
	}

	fmt.Println(bestPrice)
}

func getPrice(secretNumber int) int {
	return secretNumber % 10
}

func getPrices(initialSecretNumber int) []int {
	result := make([]int, 0)

	curr := initialSecretNumber
	result = append(result, getPrice(curr))

	for i := 0; i < iterations; i++ {
		next := evolveSecretNumber(curr)
		result = append(result, getPrice(next))
		curr = next
	}

	return result
}

func getPricesChangesMap(prices []int) map[[4]int]int {
	result := make(map[[4]int]int)

	if len(prices) < 2 {
		return result
	}

	for i := 4; i < len(prices); i++ {
		priceChanges := [4]int{prices[i-3] - prices[i-4], prices[i-2] - prices[i-3], prices[i-1] - prices[i-2], prices[i] - prices[i-1]}
		_, exists := result[priceChanges]
		if !exists {
			result[priceChanges] = prices[i]
		}
	}

	return result
}
