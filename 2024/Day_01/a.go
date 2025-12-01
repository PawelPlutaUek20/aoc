package Day_01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func A() {
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

	slices.Sort(leftList)
	slices.Sort(rightList)

	listsDiffs := []int{}

	for i, leftListItem := range leftList {
		rightListItem := rightList[i]
		diff := leftListItem - rightListItem

		if diff < 0 {
			diff = diff * -1
		}

		listsDiffs = append(listsDiffs, diff)
	}

	diffsSum := 0
	for _, diff := range listsDiffs {
		diffsSum += diff
	}

	fmt.Println(diffsSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
