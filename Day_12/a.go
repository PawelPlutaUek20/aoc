package Day_12

import (
	"log"
	"os"
	"strings"
)

func A() {
	file, err := os.ReadFile("Day_12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Fields(input)

	grid := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, letter := range line {
			row = append(row, letter)
		}
		grid = append(grid, row)
	}

	plantRegions := walk(grid)
	price := calcPrice(grid, plantRegions)

	log.Println(price)
}

type Point struct {
	Y, X int
}

type PlantRegion struct {
	plant  rune
	region []Point
}

var directions [4]Point = [4]Point{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func walk(grid [][]rune) []PlantRegion {
	visited := make(map[Point]bool)
	plantRegions := make([]PlantRegion, 0)

	for y, row := range grid {
		for x, plant := range row {
			point := Point{y, x}
			if visited[point] {
				continue
			}

			new := dfs(grid, point)

			plantRegion := PlantRegion{}
			plantRegion.plant = plant
			plantRegion.region = []Point{}

			for point := range new {
				visited[point] = true
				plantRegion.region = append(plantRegion.region, point)
			}

			plantRegions = append(plantRegions, plantRegion)
		}
	}

	return plantRegions
}

func dfs(grid [][]rune, startingPoint Point) map[Point]bool {
	seen := make(map[Point]bool)
	stack := make([]Point, 0)

	seen[startingPoint] = true
	stack = append(stack, startingPoint)

	for len(stack) > 0 {
		poped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, direction := range directions {
			nextPoint := Point{Y: poped.Y + direction.Y, X: poped.X + direction.X}
			_, visited := seen[nextPoint]
			if visited {
				continue
			}

			inBounds := nextPoint.X >= 0 && nextPoint.X < len(grid[0]) && nextPoint.Y >= 0 && nextPoint.Y < len(grid)
			if !inBounds {
				continue
			}

			sameType := grid[poped.Y][poped.X] == grid[nextPoint.Y][nextPoint.X]
			if !sameType {
				continue
			}

			seen[nextPoint] = true
			stack = append(stack, nextPoint)
		}
	}

	return seen
}

func calcPrice(grid [][]rune, plantRegions []PlantRegion) int {
	price := 0

	for _, plantRegion := range plantRegions {
		area := len(plantRegion.region)

		perimeter := 0
		for _, point := range plantRegion.region {
			for _, direction := range directions {
				nextPoint := Point{Y: point.Y + direction.Y, X: point.X + direction.X}
				inBounds := nextPoint.X >= 0 && nextPoint.X < len(grid[0]) && nextPoint.Y >= 0 && nextPoint.Y < len(grid)
				if !inBounds {
					perimeter++
					continue
				}
				if grid[nextPoint.Y][nextPoint.X] != plantRegion.plant {
					perimeter++
				}
			}
		}

		price += perimeter * area
	}

	return price
}
