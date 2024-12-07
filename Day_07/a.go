package Day_07

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func A() {
	file, _ := os.Open("Day_07/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	output := 0

	for scanner.Scan() {
		text := scanner.Text()

		equation := strings.Split(text, ":")
		result, _ := strconv.Atoi(equation[0])

		operators := []int{}
		for _, op := range strings.Fields(equation[1]) {
			operator, _ := strconv.Atoi(op)
			operators = append(operators, operator)
		}

		if checkCanBeSolved(result, operators, operators[0], 1) {
			output += result
		}
	}

	log.Println(output)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkCanBeSolved(result int, operators []int, currSum int, currIdx int) bool {
	if currIdx == len(operators) {
		return result == currSum
	}

	if currIdx >= len(operators) {
		return false
	}

	if checkCanBeSolved(result, operators, currSum+operators[currIdx], currIdx+1) {
		return true
	}

	if checkCanBeSolved(result, operators, currSum*operators[currIdx], currIdx+1) {
		return true
	}

	return false
}
