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
	bac2.position.x, bac2.position.y = 500, 500
	bac3.position.x, bac3.position.y = 450, 550
	bac4.position.x, bac4.position.y = 350, 450

	bac1.width, bac1.length = 20, 50
	bac2.width, bac2.length = 10, 40
	bac3.width, bac3.length = 20, 50
	bac4.width, bac4.length = 15, 45

	//Take pointers for each cell
	b1p := &bac1
	b2p := &bac2
	b3p := &bac3
	b4p := &bac4

	//Initialize culture
	initialCulture.cells = []*RodCell{b1p, b2p, b3p, b4p}

	//Test Run BioFilm-Model simulation
	timePoints := SimulateBiofilm(initialCulture, 1000, 5)
	fmt.Println("Simulation Complete")
	fmt.Println("Drawing cultures...")

	//Animate simulated cultures, create a GIF result
	images := AnimateSystem(timePoints, 1000, 5)
	fmt.Println("Images drawn!")
	fmt.Println("Generating an animated GIF...")
	gifhelper.ImagesToGIF(images, "output")
	fmt.Println("GIF Drawn!")
	fmt.Println("Simulation Complete!")
}
