package main

import (
	"fmt"
	"math"
)

// MakeNutritionBoard takes as input the width of the culture and the nutrition value
// Returns a 2D slice of ints representing the nutrition board of the culture
func MakeNutritionBoard(NBwidth int, nutritionValue int, nutritionShape string, NBfromFile [][]int, dontSpread bool) [][]int {

	nutritionBoard := MakeSquareBoard(NBwidth)
	// if input is given from file, use the input board as nutritionBoard
	// note: len(NBfromFile) should be 0 when not reading from input
	if len(NBfromFile) > 0 {
		nutritionBoard.UpdateBoardFromFile(NBfromFile, dontSpread)

		// if the nutrition shape is circle, add nutrition to the circle
	} else if nutritionShape == "circle" {
		nutritionBoard.AddToCircle(nutritionValue)

		// add nutrition to the whole board
	} else {
		nutritionBoard.AddToWholeBoard(nutritionValue)
	}

	return nutritionBoard
}

// UpdateBoardFromFile takes the board from file as input
// returns nutrition board appropriate to its size and settings
func (nb NutritionBoard) UpdateBoardFromFile(NBfromFile [][]int, dontSpread bool) {
	size := len(nb)
	size2 := len(NBfromFile)

	if size == len(NBfromFile) || dontSpread { //if input board is the same size as the nutBoard
		fmt.Println("Making nutrition board with size", size)
		for x := 0; x < size2; x++ {
			for y := 0; y < size2; y++ {
				nb[x][y] = NBfromFile[x][y]
			}
		}

		//if input board is smaller than the set nutrition board, it will try to spread the values as far apart as possible
	} else if size > size2 {
		fmt.Println("Spread values from input board of size", size2, "to nutrition board of size", size, "...")
		nb.SpreadValues(NBfromFile)
	} else { //if input board is larger than the nutBoard
		fmt.Println("Fit values from input board of size", size2, "into nutrition board of size", size, "...")
		nb.FitValues(NBfromFile)
	}
}

// FitValues is ran when the size of the Nutrition Board is smaller than the given board input
// Function updates the Nutrition Board by trying to squeeze the values into Nutrition Board
func (nb NutritionBoard) FitValues(NBfromFile [][]int) {

	//interval determines the step size to pull value from file matrix
	interval := int(math.Round(float64(len(NBfromFile)-1) / float64(len(nb)-1)))
	fmt.Println(interval)

	for i := range nb {
		for j := range nb[i] {
			row := i * interval
			if row > len(NBfromFile)-1 {
				row = len(NBfromFile) - 1
			}
			col := j * interval
			if col > len(NBfromFile)-1 {
				col = len(NBfromFile) - 1
			}

			nb[i][j] = NBfromFile[row][col]
		}

	}

}

// SpreadValues is ran when the size of the Nutrition Board is larger than the given input (unless parameter dontSpread is true)
// Function updates the Nutrition Board by spreading the values as far apart as possible
func (nb NutritionBoard) SpreadValues(NBfromFile [][]int) {

	//interval determines the step size the value will be placed in the NutritionBoard
	interval := int(math.Round(float64(len(nb)-1) / float64(len(NBfromFile)-1)))

	for i := range NBfromFile {
		for j, val := range NBfromFile[i] {
			row := i * interval
			if row > len(nb)-1 {
				row = len(nb) - 1
			}
			col := j * interval
			if col > len(nb)-1 {
				col = len(nb) - 1
			}
			nb[row][col] = val
		}
	}
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
