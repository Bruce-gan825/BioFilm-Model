package main

import (
	"fmt"
	"gifhelper"
	"math/rand"
	"time"
)

func main() {

	//======================= set nutrition board parameters ==============================
	// import text file?
	importNutritionBoardFromFile := true
	//                          put FILE NAME HERE       (and put file in NutritionBoardInputs folder)
	filename := "NutritionBoardInputs/nutritionBoard2.txt" //can update this to be a command line argument

	var NBfromFile [][]int

	if importNutritionBoardFromFile { //if import is true
		NBfromFile = ReadNutritionBoardFromFile(filename) // NBfromFile is a 2D slice of ints
	}

	NBwidth := 1000            // if NBwidth = culture width, nutrition placed every pixel of the board
	nutritionValue := 10       // value of nutrition in each pixel
	nutritionShape := "circle" //make circle nutrition board?
	dontSpread := false        // if input board size < wanted nutrition board size, spread values?

	// MakeNutritionBoard order of prioritization:
	//1. given input Board from file -> use input Board as Nutrition Board
	//2. nutritionShape = circle -> make circular nutrition board
	//3. add nutrition value to every pixel
	nutrition := MakeNutritionBoard(NBwidth, nutritionValue, nutritionShape, NBfromFile, dontSpread)
	//================================================================================

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

	//Initialize biofilm
	initialCulture.biofilms = make([]*Biofilm, 4)

	for i := range initialCulture.biofilms {
		initialCulture.biofilms[i] = &Biofilm{} // Allocate a new Biofilm
	}
	//Initialize culture
	initialCulture.biofilms[0].cells = []*SphereCell{s1p}
	initialCulture.biofilms[1].cells = []*SphereCell{s2p}
	initialCulture.biofilms[2].cells = []*SphereCell{s3p}
	initialCulture.biofilms[3].cells = []*SphereCell{s4p}

	initialCulture.nutrition = nutrition

	//fmt.Println(initialCulture.nutrition)

	//----initialCulture2 - just one cell in the middle--------------
	var initialCulture2 Culture
	initialCulture2.width = 500

	var cell SphereCell
	cell.cellID = 1
	cell.radius = 4
	cell.red, cell.green, cell.blue = 20, 45, 100
	cell.position.x, cell.position.y = 250, 250

	initialCulture2.biofilms = make([]*Biofilm, 1)
	initialCulture2.biofilms[0] = &Biofilm{}
	initialCulture2.biofilms[0].cells = []*SphereCell{&cell}
	initialCulture2.nutrition = nutrition

	//--------randomly generate a culture of spherical cells------------------

	//Add an assortment of random cells
	//Create an initial culture
	var initialCulture3 Culture
	initialCulture3.width = 1000

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	initialCulture3.biofilms = make([]*Biofilm, 1)
	initialCulture3.biofilms[0] = &Biofilm{}
	initialCulture3.biofilms[0].cells = make([]*SphereCell, 200)
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

		initialCulture3.biofilms[0].cells[i] = &cell
	}
	initialCulture3.nutrition = nutrition

	//two cells that have same y position to test quorum sensing
	// var initialCulture4 Culture
	// initialCulture4.width = 1000
	// initialCulture4.cells = make([]*SphereCell, 2)
	// var cell1, cell2 SphereCell
	// cell1.red, cell1.green, cell1.blue = 180, 0, 180
	// cell2.red, cell2.green, cell2.blue = 20, 200, 20
	// cell1.position.x, cell1.position.y = 250, 500
	// cell2.position.x, cell2.position.y = 750, 500
	// cell1.radius, cell2.radius = 25, 25
	// cell1.cellID, cell2.cellID = 1, 2
	// initialCulture4.cells = []*SphereCell{&cell1, &cell2}
	// initialCulture4.nutrition = nutrition

	//55 cells in the same biofilm, for biofilm division test
	var initialCulture5 Culture
	initialCulture5.width = 1000
	initialCulture5.biofilms = make([]*Biofilm, 1)
	initialCulture5.biofilms[0] = &Biofilm{}
	initialCulture5.biofilms[0].cells = make([]*SphereCell, 55)
	for i := 0; i < 5; i++ {
		for j := 0; j < 11; j++ {
			var cell SphereCell
			cell.position.x = float64(400 + (j * 20))
			cell.position.y = float64(450 + (i * 20))
			cell.radius = 10
			cell.red, cell.green, cell.blue = 100, 200, 0
			initialCulture5.biofilms[0].cells[i*11+j] = &cell
		}
	}

	//----------cell growth parameters------------------
	//growthRate is a constant that determines how much cells grow per time interval
	// 0.1 = 10% growth per time interval
	cellGrowthRate := 0.17
	//maxRadius is a constant that determines the maximum radius a cell can grow to before dividing
	cellMaxRadius := 20.0
	//cellGrowthNutritionThreshold is a constant that determines the minimum amount of nutrition a cell must have before it can grow
	cellGrowthNutritionThreshold := 0.001

	//example parameters
	//width 1000 and cell radius ~10 - growthRate 0.07, maxRadius 20, threshold 1.6
	//width 30 and cell radius 1 - growthRate 0.17, maxRadius 1.5, threshold 0.001 : single cell
	//width 50 (nutrition file2 30x30 gradient) radius 4 - growthRate 0.17, maxRadius 2.3, threshold 0.001
	//--------------------------------------------------

	//Test Run BioFilm-Model simulation
	timePoints := SimulateBiofilm(initialCulture5, 50, 1, cellGrowthRate, cellMaxRadius, cellGrowthNutritionThreshold)

	//fmt.Println(timePoints[len(timePoints)-1].nutrition)
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
