package Day_15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func A() {

	file, _ := os.Open("Day_15/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := ""
	for scanner.Scan() {
		text := scanner.Text()
		lines += text
	}

	result := 0
	steps := strings.Split(lines, ",")

	for _, step := range steps {
		result += hashAlgorithm(step, 0)
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func hashAlgorithm(str string, curr int) int {
	char := []rune(str)

	for _, r := range char {
		ascii := int(r)
		curr = (((curr + ascii) * 17) % 256)
	}

	return curr
}
