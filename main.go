package main

import (
	"fmt"
	"gifhelper"
)

func main() {
	//Create an initial culture
	var initialCulture Culture
	initialCulture.width = 1000

	//Create a cell
	var bac1, bac2, bac3, bac4 RodCell
	bac1.red, bac1.green, bac1.blue = 185, 34, 22
	bac2.red, bac2.green, bac2.blue = 48, 0, 44
	bac3.red, bac3.green, bac3.blue = 0, 211, 202
	bac4.red, bac4.green, bac4.blue = 240, 240, 100

	bac1.position.x, bac1.position.y = 400, 400
	bac2.position.x, bac2.position.y = 600, 600
	bac3.position.x, bac3.position.y = 700, 550
	bac4.position.x, bac4.position.y = 350, 450

	//Angle of cell in radians
	bac1.angle = 0.4
	bac2.angle = 0.2
	bac3.angle = 0.52
	bac4.angle = 0.4

	bac1.width, bac1.length = 20, 50
	bac2.width, bac2.length = 10, 40
	bac3.width, bac3.length = 20, 50
	bac4.width, bac4.length = 15, 45

	//set maxLength
	bac1.maxLength = 120
	bac2.maxLength = 110
	bac3.maxLength = 120
	bac4.maxLength = 115

	//Test elongate
	bac1.Elongate(400)
	bac5, bac6 := bac1.Divide()

	//Take pointers for each cell
	//b1p := &bac1
	b2p := &bac2
	b3p := &bac3
	b4p := &bac4

	//Initialize culture
	initialCulture.cells = []*RodCell{b2p, b3p, b4p, bac5, bac6}

	//----------this code is used for testing InitializeCulture function-------

	/*numCells := 50
	cultureWidth := 1000.0
	cellWidth := 20.0
	maxCellLength := 50.0

	initialCulture2 := InitializeCulture(numCells, cultureWidth, cellWidth, maxCellLength)*/

	//-------------------------------

	//Test Run BioFilm-Model simulation
	/*timePoints := SimulateBiofilm(initialCulture, 500, 5)
	fmt.Println("Simulation Complete")
	fmt.Println("Drawing cultures...")

	//Animate simulated cultures, create a GIF result
	images := AnimateSystem(timePoints, 1000, 5)
	fmt.Println("Images drawn!")
	fmt.Println("Generating an animated GIF...")
	gifhelper.ImagesToGIF(images, "output")
	fmt.Println("GIF Drawn!")
	fmt.Println("Simulation Complete!")*/

	//Test Elongate
	/*timePoints := SimulateBiofilm(initialCulture, 0, 5)
	fmt.Println("Simulation Complete")
	fmt.Println("Drawing cultures...")
	timePoints[0].cells[0].Elongate(50)
	images := AnimateSystem(timePoints, 1000, 1)
	fmt.Println("Images drawn!")
	fmt.Println("Generating an animated GIF...")
	gifhelper.ImagesToGIF(images, "Elongate50")
	fmt.Println("GIF Drawn!")
	fmt.Println("Simulation Complete!")*/

	//Test Divide
	timePoints := SimulateBiofilm(initialCulture, 0, 5)
	fmt.Println("Simulation Complete")
	fmt.Println("Drawing cultures...")
	timePoints[0].cells[0].Elongate(50)
	children := make([]*RodCell, 2)
	children[0], children[1] = timePoints[0].cells[0].Divide()
	timePoints[0].cells = append(timePoints[0].cells[1:], children...)
	images := AnimateSystem(timePoints, 1000, 1)
	fmt.Println("Images drawn!")
	fmt.Println("Generating an animated GIF...")
	gifhelper.ImagesToGIF(images, "Divide")
	fmt.Println("GIF Drawn!")
	fmt.Println("Simulation Complete!")
}
