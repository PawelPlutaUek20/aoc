package Day_24

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("Day_24/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	sections := strings.Split(strings.TrimSpace(input), "\n\n")

	line1 := strings.Split(strings.TrimSpace(sections[0]), "\n")
	line2 := strings.Split(strings.TrimSpace(sections[1]), "\n")

	wires := make(map[string]int)
	for _, line := range line1 {
		data := strings.Split(strings.TrimSpace(line), ": ")
		wire := data[0]
		bit, _ := strconv.Atoi(data[1])
		wires[wire] = bit
	}

	q := make([][5]string, 0)
	for _, line := range line2 {
		parts := strings.Split(strings.TrimSpace(line), " ")
		q = append(q, [5]string{parts[0], parts[1], parts[2], parts[3], parts[4]})
	}

	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		wire1, gate, wire2, resultWire := item[0], item[1], item[2], item[4]

		w1, existsW1 := wires[wire1]
		w2, existsW2 := wires[wire2]

		if existsW1 && existsW2 {
			wires[resultWire] = getResultWireValue(w1, gate, w2)
		} else {
			q = append(q, item)
		}
	}

	result := 0
	for wire, bit := range wires {
		if wire[0] == 'z' && bit == 1 {
			significantBit, _ := strconv.Atoi(strings.TrimPrefix(wire, "z"))
			result += 1 << significantBit
		}
	}
	fmt.Println(result)
}

func getResultWireValue(wire1 int, gate string, wire2 int) int {
	switch gate {
	case "OR":
		return wire1 | wire2
	case "XOR":
		return wire1 ^ wire2
	case "AND":
		return wire1 & wire2
	default:
		panic("unknown gate")
	}
}
