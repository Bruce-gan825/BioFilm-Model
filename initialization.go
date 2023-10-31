package main

import (
	"math/rand"
)

// InitializeCulture takes as imput the number of cells, culture width, cell width and cell maxlength
// Returns a Culture object with numCells having random colors, length (< maxlength), angle, and position (within culture width)
// Culture width, cellWidth and and cellMaxLength are set values (not random)
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
