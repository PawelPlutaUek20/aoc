package Day_11

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func B() {
	file, err := os.ReadFile("Day_11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Fields(input)

	sum := 0

	for _, line := range lines {
		sum += run(line, 75)
	}

	log.Println(sum)
}

type stone struct {
	number string
	step   int
}

var cache = map[stone]int{}

func run(number string, step int) int {
	stone := stone{number, step}

	cached, exists := cache[stone]
	if exists {
		return cached
	}

	if step == 0 {
		return 1
	}

	var res int
	if number == "0" {
		res = run("1", step-1)
	} else if len(number)%2 == 0 {
		leftStone, _ := strconv.Atoi(number[0 : len(number)/2])
		rightStone, _ := strconv.Atoi(number[len(number)/2:])
		res = run(strconv.Itoa(leftStone), step-1) + run(strconv.Itoa(rightStone), step-1)
	} else {
		stoneNum, _ := strconv.Atoi(number)
		res = run(strconv.Itoa(stoneNum*2024), step-1)
	}

	cache[stone] = res
	return res
}
