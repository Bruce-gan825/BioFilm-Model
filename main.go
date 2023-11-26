package main

import (
	"fmt"
	"gifhelper"
	"math/rand"
	"time"
)

func main() {
	//Create an initial culture
	var initialCulture Culture
	initialCulture.width = 1000

	//Create a collection of spherical cells
	var sta1, sta2, sta3, sta4 SphereCell
	sta1.red, sta1.green, sta1.blue = 20, 45, 100
	sta2.red, sta2.green, sta2.blue = 240, 190, 50
	sta3.red, sta3.green, sta3.blue = 166, 0, 200
	sta4.red, sta4.green, sta4.blue = 0, 85, 150

	sta1.position.x, sta1.position.y = 400, 400
	sta2.position.x, sta2.position.y = 600, 100
	sta3.position.x, sta3.position.y = 700, 550
	sta4.position.x, sta4.position.y = 350, 450

	//Radius of spherical cells
	sta1.radius, sta2.radius, sta3.radius, sta4.radius = 25, 10, 5, 11
	sta1.cellID = 1
	sta2.cellID = 2
	sta3.cellID = 3
	sta4.cellID = 4

	//Set initial movement of cells
	sta1.velocity.x, sta1.velocity.y = 0, 0
	sta2.velocity.x, sta2.velocity.y = -0.1, -0.1
	sta3.velocity.x, sta3.velocity.y = 0.1, 0.1
	sta4.velocity.x, sta4.velocity.y = 0.1, 0.1

	sta1.acceleration.x, sta1.acceleration.y = 0, 0
	sta2.acceleration.x, sta2.acceleration.y = 0, 0
	sta3.acceleration.x, sta3.acceleration.y = 0, 0
	sta4.acceleration.x, sta4.acceleration.y = 0, 0

	//Take points for each cell
	s1p := &sta1
	s2p := &sta2
	s3p := &sta3
	s4p := &sta4

	//Initialize culture
	initialCulture.cells = []*SphereCell{s1p, s2p, s3p, s4p}

	nutritionValue := 10

	nutritionShape := "circle"

	nutrition := MakeNutritionBoard(int(initialCulture.width), nutritionValue, nutritionShape)

	initialCulture.nutrition = nutrition

	//----initialCulture2 - just one cell in the middle
	var initialCulture2 Culture
	initialCulture2.width = 1000

	var cell SphereCell
	cell.cellID = 1
	cell.radius = 10
	cell.red, cell.green, cell.blue = 20, 45, 100
	cell.position.x, cell.position.y = 500, 500

	initialCulture2.cells = []*SphereCell{&cell}
	initialCulture2.nutrition = nutrition

	//--------randomly generate a culture of spherical cells------------------

	//Add an assortment of random cells
	//Create an initial culture
	var initialCulture3 Culture
	initialCulture3.width = 1000

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	initialCulture3.cells = make([]*SphereCell, 200)
	for i := 0; i < 200; i++ {
		var cell SphereCell
		cell.cellID = i + 1
		cell.radius = 10

		//generate random position within cultureWidth
		cell.position.x, cell.position.y = rand.Float64()*1000, rand.Float64()*1000
		//generate random velocity
		cell.velocity.x, cell.velocity.y = (-2 + rand.Float64()*4), (-2 + rand.Float64()*4)
		// Call the RandomDiffusion method to simulate the random diffusion of the cell
		cell.RandomDiffusion()
		// generate random rgb
		cell.red, cell.green, cell.blue = uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256))

		initialCulture3.cells[i] = &cell
	}
	initialCulture3.nutrition = nutrition

	//----------cell growth parameters------------------
	//growthRate is a constant that determines how much cells grow per time interval
	// 0.1 = 10% growth per time interval
	cellGrowthRate := 0.07
	//maxRadius is a constant that determines the maximum radius a cell can grow to before dividing
	cellMaxRadius := 20.0
	//cellGrowthNutritionThreshold is a constant that determines the minimum amount of nutrition a cell must have before it can grow
	cellGrowthNutritionThreshold := 1.6
	//--------------------------------------------------

	//Test Run BioFilm-Model simulation
	timePoints := SimulateBiofilm(initialCulture2, 400, 1, cellGrowthRate, cellMaxRadius, cellGrowthNutritionThreshold)

	fmt.Println("Simulation Complete")
	fmt.Println("Drawing cultures...")

	//Animate simulated cultures, create a GIF result
	images := AnimateSystem(timePoints, 1000, 1)
	fmt.Println("Images drawn!")
	fmt.Println("Generating an animated GIF...")
	gifhelper.ImagesToGIF(images, "output")
	fmt.Println("GIF Drawn!")
	fmt.Println("Simulation Complete!")

}

//--- functions used for ROD CELLS-----------------
//Test with rod-shaped cells
/*
	//Create a cell
	var bac1, bac2, bac3, bac4 RodCell
	bac1.red, bac1.green, bac1.blue = 185, 34, 22
	bac2.red, bac2.green, bac2.blue = 48, 0, 44
	bac3.red, bac3.green, bac3.blue = 0, 211, 202
	bac4.red, bac4.green, bac4.blue = 240, 240, 100

	bac1.position.x, bac1.position.y = 400, 400
	bac2.position.x, bac2.position.y = 600, 100
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

*/

//----------this code is used for testing InitializeCulture function ...FOR ROD CELLS -------

/*numCells := 50
cultureWidth := 1000.0
cellWidth := 20.0
maxCellLength := 50.0

initialCulture2 := InitializeCulture(numCells, cultureWidth, cellWidth, maxCellLength)*/

//-------------------------------

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
/*
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
*/
