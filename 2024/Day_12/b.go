package Day_12

import (
	"log"
	"os"
	"strings"
)

func B() {
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
	price := calcPriceB(grid, plantRegions)

	log.Println(price)
}

func calcPriceB(grid [][]rune, plantRegions []PlantRegion) int {
	price := 0

	for _, plantRegion := range plantRegions {
		area := len(plantRegion.region)
		corners := 0
		for _, point := range plantRegion.region {
			for i, direction := range directions {
				nextDirection := directions[(i+1)%4]

				if !isPlantRegion(grid, plantRegion.plant, point, direction) && !isPlantRegion(grid, plantRegion.plant, point, nextDirection) {
					corners++
					continue
				}

				if isPlantRegion(grid, plantRegion.plant, point, direction) && isPlantRegion(grid, plantRegion.plant, point, nextDirection) && !isPlantRegion(grid, plantRegion.plant, point, Point{X: direction.X + nextDirection.X, Y: direction.Y + nextDirection.Y}) {
					corners++
				}
			}
		}

		price += area * corners
	}

	return price
}

func isPlantRegion(grid [][]rune, plantType rune, point Point, direction Point) bool {
	nextPoint := Point{Y: point.Y + direction.Y, X: point.X + direction.X}
	inBounds := nextPoint.X >= 0 && nextPoint.X < len(grid[0]) && nextPoint.Y >= 0 && nextPoint.Y < len(grid)
	if !inBounds {
		return false
	}
	if grid[nextPoint.Y][nextPoint.X] != plantType {
		return false
	}

	return true
}
