package Day_01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func B() {
	file, _ := os.Open("Day_01/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		text := scanner.Text()

		listItems := strings.Fields(text)
		leftListItem, _ := strconv.Atoi(listItems[0])
		rightListItem, _ := strconv.Atoi(listItems[1])

		leftList = append(leftList, leftListItem)
		rightList = append(rightList, rightListItem)
	}

	appearancesByLocationID := make(map[int]int)
	for _, rightListLocationID := range rightList {
		appearancesByLocationID[rightListLocationID] += 1
	}

	similaritiesSum := 0
	for _, leftListLocationID := range leftList {
		similaritiesSum += leftListLocationID * appearancesByLocationID[leftListLocationID]
	}

	fmt.Println(similaritiesSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
