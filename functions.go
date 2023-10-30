package main

import (
	"fmt"
)

// SimulateBiofilm() takes as input a Culture object initialCulture, a number of generations parameter numGens, and a time interval timeStep.
// It simulate the biofilm over numGens generations starting with initialCulture in which the system is updated every timeStep units of time, and as a result it should return a slice of numGens + 1 Culture objects
func SimulateBiofilm(initialCulture Culture, numGens int, time float64) []Culture {
	timePoints := make([]Culture, numGens+1)
	timePoints[0] = initialCulture
	for i := 1; i < numGens+1; i++ {
		timePoints[i] = UpdateCulture(timePoints[i-1], time)
	}
	return timePoints
}

// UpdateCulture takes as input a Culture object and a time float64 parameter.
// It returns a new Culture object corresponding to updating each Cell's position, velocity, and acceleration within
// the given time interval
func UpdateCulture(currentCulture Culture, time float64) Culture {
	//Create a copy of the current culture to alter
	newCulture := CopyCulture(currentCulture)

	//Iterate over all Cells in the newly created Culture and update their fields
	for i := range newCulture.cells {
		//Update position functions go here
		if i > 99999999999999 {
			fmt.Println("No")
		}
	}

	return newCulture
}

// CopyCulture returns a different culture that has same cells as the input culture
func CopyCulture(culture Culture) Culture {
	var newCulture Culture
	newCulture.cells = make([]*RodCell, len(culture.cells))
	for i := range newCulture.cells {
		newCulture.cells[i] = CopyCell(culture.cells[i])
	}
	newCulture.width = culture.width
	return newCulture
}

// CopyCell returns a different cell that has same fields as the input cell
func CopyCell(cell *RodCell) *RodCell {
	//newCell := &RodCell{}
	var newCell RodCell
	newCell.length = cell.length
	newCell.maxLength = cell.maxLength
	newCell.width = cell.width
	newCell.angle = cell.angle
	newCell.position.x = cell.position.x
	newCell.position.y = cell.position.y
	newCell.red = cell.red
	newCell.blue = cell.blue
	newCell.green = cell.green

	return &newCell
}
