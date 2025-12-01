package Day_17

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type register struct {
	a int
	b int
	c int

	output []string

	ptr int
}

func A() {
	file, err := os.ReadFile("Day_17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("[0-9]+")
	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(input, "\n\n")

	registers := strings.Split(lines[0], "\n")
	program := lines[1]

	registerA, _ := strconv.Atoi(re.FindString(registers[0]))
	registerB, _ := strconv.Atoi(re.FindString(registers[1]))
	registerC, _ := strconv.Atoi(re.FindString(registers[2]))

	r := register{
		a: registerA,
		b: registerB,
		c: registerC,

		output: make([]string, 0),

		ptr: 0,
	}

	instructions := re.FindAllString(program, -1)

	for r.ptr < len(instructions) {
		opcode, _ := strconv.Atoi(instructions[r.ptr])
		operand, _ := strconv.Atoi(instructions[r.ptr+1])

		r.run(opcode, operand)
	}

	log.Println(strings.Join(r.output, ","))
}

func (r *register) combo(operand int) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return r.a
	case 5:
		return r.b
	case 6:
		return r.c
	default:
		panic("invalid opcode")
	}
}

func (r *register) run(opcode int, operand int) {
	switch opcode {
	case 0:
		r.a = r.a >> r.combo(operand)
	case 1:
		r.b = r.b ^ operand
	case 2:
		r.b = r.combo(operand) % 8
	case 3:
		if r.a != 0 {
			r.ptr = operand
			return
		}
	case 4:
		r.b = r.b ^ r.c
	case 5:
		r.output = append(r.output, strconv.Itoa(r.combo(operand)%8))
	case 6:
		r.b = r.a >> r.combo(operand)
	case 7:
		r.c = r.a >> r.combo(operand)
	}

	r.ptr += 2
}
