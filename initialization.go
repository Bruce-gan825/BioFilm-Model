package main

import "math"

// InitializeCulture takes as imput the number of cells, culture width, cell width and cell maxlength
// Returns a Culture object with numCells having random colors, length (< maxlength), angle, and position (within culture width)
// Culture width, cellWidth and and cellMaxLength are set values (not random)
/*
func InitializeCulture(numCells int, cultureWidth, cellWidth, cellMaxLength float64) Culture {
	var cture Culture
	cture.width = cultureWidth
	cture.cells = make([]*RodCell, numCells)

	for i := 0; i < numCells; i++ {
		var cell RodCell
		cell.width = cellWidth
		cell.maxLength = cellMaxLength
		// generate random angle between 0 - 3.14 rad (180 degrees)
		cell.angle = rand.Float64() * 3.14
		// generate random length less than cellMaxLength
		cell.length = rand.Float64() * cellMaxLength
		//generate random position within cultureWidth
		cell.position.x, cell.position.y = rand.Float64()*cultureWidth, rand.Float64()*cultureWidth
		// generate random rgb
		cell.red, cell.green, cell.blue = uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256))

		cture.cells[i] = &cell

	}
	return cture
} //needs further update to have cells not overlay on top of each other

*/

// MakeNutritionBoard takes as input the width of the culture and the nutrition value
// Returns a 2D slice of ints representing the nutrition board of the culture
func MakeNutritionBoard(width int, nutritionValue int, nutritionShape string) [][]int {

	nutritionBoard := MakeSquareBoard(width)

	if nutritionShape == "circle" {
		nutritionBoard.AddToCircle(nutritionValue)
	} else {
		nutritionBoard.AddToWholeBoard(nutritionValue)
	}

	return nutritionBoard
}

// MakeSquareBoard takes as input the width of the culture
// Returns a 2D slice of ints
func MakeSquareBoard(width int) NutritionBoard {
	nutritionBoard := make(NutritionBoard, width)

	for i := range nutritionBoard {
		nutritionBoard[i] = make([]int, width)
	}

	return nutritionBoard
}

// AddToWholeBoard takes as input the value to add to the whole board
// Returns a 2D slice of ints
func (nb NutritionBoard) AddToWholeBoard(valueToAdd int) {
	size := len(nb)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			nb[x][y] += valueToAdd
		}
	}
}

// AddToCircle takes as input the value to add to the circle
// Returns a 2D slice of ints in the shape of a circle
func (nb NutritionBoard) AddToCircle(valueToAdd int) {
	size := len(nb)
	center := size / 2
	radius := (float64(size) / 2) * 0.9

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			distance := math.Sqrt(math.Pow(float64(x-center), 2) + math.Pow(float64(y-center), 2))

			// Check if the pixel is within the circle's rough approximation
			if distance < radius {
				nb[x][y] += valueToAdd
			}
		}
	}
}
