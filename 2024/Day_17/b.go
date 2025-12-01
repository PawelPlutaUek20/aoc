package Day_17

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type QItem struct {
	val         int
	instruction int
}

func B() {
	file, err := os.ReadFile("Day_17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(input, "\n\n")

	re := regexp.MustCompile("[0-9]+")
	program := lines[1]
	instructions := re.FindAllString(program, -1)

	q := make([]QItem, 0)
	q = append(q, QItem{0, 1})

	result := 0

out:
	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		seeking := instructions[len(instructions)-item.instruction]
		found := findNext(instructions, item.val, seeking)
		for _, newItem := range found {
			if item.instruction == len(instructions) {
				result = newItem
				break out
			}
			q = append(q, QItem{newItem << 3, item.instruction + 1})
		}
	}

	log.Println(result)
}

func newRegister(a int) register {
	return register{
		a: a,
		b: 0,
		c: 0,

		ptr: 0,

		output: make([]string, 0),
	}
}

func findNext(instructions []string, curr int, seeking string) []int {
	result := make([]int, 0)

	for a := curr; a < curr+8; a++ {
		r := newRegister(a)

		for r.ptr < len(instructions) {
			opcode, _ := strconv.Atoi(instructions[r.ptr])
			operand, _ := strconv.Atoi(instructions[r.ptr+1])

			r.run(opcode, operand)
		}

		if r.output[0] == seeking {
			result = append(result, a)
		}
	}

	return result
}
