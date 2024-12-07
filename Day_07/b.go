package Day_07

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func B() {
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

		if checkCanBeSolvedB(result, operators, operators[0], 1) {
			output += result
		}
	}

	log.Println(output)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkCanBeSolvedB(result int, operators []int, currSum int, currIdx int) bool {
	if currIdx == len(operators) {
		return result == currSum
	}

	if checkCanBeSolvedB(result, operators, currSum+operators[currIdx], currIdx+1) {
		return true
	}

	if checkCanBeSolvedB(result, operators, currSum*operators[currIdx], currIdx+1) {
		return true
	}

	if checkCanBeSolvedB(result, operators, concat(currSum, operators[currIdx]), currIdx+1) {
		return true
	}

	return false
}

func concat(op1, op2 int) int {
	digits := int(math.Log10(float64(op2))) + 1
	out := op1*int(math.Pow10(digits)) + op2
	return out
}
