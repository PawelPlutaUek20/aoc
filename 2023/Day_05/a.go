package Day_05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func A() {
	file, _ := os.Open("Day_05/input.txt")
	defer file.Close()

	const (
		readSeeds = iota
		mappingFns
	)

	scanner := bufio.NewScanner(file)

	accSeeds := make([]int, 0)
	currSeeds := make([]int, 0)

	currMode := readSeeds

	r := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		text := scanner.Text()

		if currMode == readSeeds {
			accSeeds = mapStringToInt(r.FindAllString(text, -1))
			currSeeds = slices.Clone(accSeeds)
			currMode = mappingFns
		} else {
			data := mapStringToInt(r.FindAllString(text, -1))

			if len(data) == 0 {
				accSeeds = slices.Clone(currSeeds)
				continue
			}

			dst := data[0]
			src := data[1]
			length := data[2]

			for i, seed := range accSeeds {
				if seed >= src && seed < src+length {
					currSeeds[i] = dst + seed - src
				}
			}
		}
	}

	result := slices.Min(currSeeds)
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func mapStringToInt(arr []string) []int {
	arr2 := make([]int, 0)

	for _, v := range arr {
		num, _ := strconv.Atoi(v)
		arr2 = append(arr2, num)
	}

	return arr2
}
