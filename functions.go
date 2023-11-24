package main

//THETA IS IN RADIANS (0 - 2 Pi)

import (
	"math"
	"math/rand"
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

	//growthRate is a constant that determines how much cells grow per time interval
	// 0.1 = 10% growth per time interval
	growthRate := 0.05
	//maxRadius is a constant that determines the maximum radius a cell can grow to before dividing
	maxRadius := 20.0
	//cellGrowthNutritionThreshold is a constant that determines the minimum amount of nutrition a cell must have before it can grow
	cellGrowthNutritionThreshold := 1.6

	//Update particles
	for i := range newCulture.particles {
		newCulture.particles[i].position = UpdateParticle(newCulture.particles[i], time)
	}

	//Iterate over all Cells in the newly created Culture and update their fields
	for i := range newCulture.cells {
		//Update position functions go here
		newCulture.cells[i].acceleration = UpdateAcceleration(currentCulture, newCulture.cells[i])
		newCulture.cells[i].velocity = UpdateVelocity(newCulture.cells[i], time)
		newCulture.cells[i].position = UpdatePosition(newCulture.cells[i], time)

		//Update cellNutrition according to nutritions on board and cell's position
		//Also update the nutrition level on board after it's consumed by the cell
		newCulture.cells[i].cellNutrition = ConsumeNutrients(newCulture.nutrition, newCulture.cells[i])

		//Every cell should release signals
		newParticles := newCulture.cells[i].ReleaseSignals(20, 4)
		newCulture.particles = append(newCulture.particles, newParticles...)
		//grow cells if cell's nutrition level is greater than threshold
		if newCulture.cells[i].cellNutrition >= cellGrowthNutritionThreshold {
			newCulture.cells[i].radius = GrowCellSpherical(newCulture.cells[i], growthRate)
			newCulture.cells[i].cellNutrition -= cellGrowthNutritionThreshold //spend energy to grow
		}
		//divide cells if radius is greater than maxRadius
		if newCulture.cells[i].radius >= maxRadius {
			child1, child2 := DivideCellSpherical(newCulture.cells[i])
			newCulture.cells[i] = child1                        //replace original cell with child1
			newCulture.cells = append(newCulture.cells, child2) //append child2 to culture
		}
	}

	for i := range newCulture.cells {
		newCulture.cells[i].ReceiveSignals(newCulture.particles)
	}

	//Apply simple collision function for the newCulture
	CheckSphereCollision(newCulture)

	return newCulture

}

// ConsumeNutrients takes as input a nutrition board and a SphereCell object
// It returns the updated cellNutrition of the SphereCell object after consuming nutrients
// And updates the nutrition board accordingly
func ConsumeNutrients(nutritionBoard [][]int, s *SphereCell) float64 {
	//Note: the cell is treated as a square (width 2*radius) for the purposes of nutrient consumption

	cellNutrient := s.cellNutrition //current cellNutrition

	//Iterate over the nutrition board and check if cell is in contact with any nutrition
	//If so, cellNutrition increases by 1 and nutrition level decreases by 1
	xstart := int(s.position.x - s.radius)
	xend := int(s.position.x + s.radius)
	ystart := int(s.position.y - s.radius)
	yend := int(s.position.y + s.radius)

	for i := xstart; i < xend; i++ {
		for j := ystart; j < yend; j++ {
			if inBounds(i, j, len(nutritionBoard)) && nutritionBoard[i][j] > 0 {
				cellNutrient++
				nutritionBoard[i][j]--
			}
		}
	}
	return cellNutrient

}

// inBounds checks if a given x and y coordinate is within the bounds of a square board of width width
// It returns true if the x and y coordinate is within the bounds of the board, and false if otherwise
func inBounds(x, y, width int) bool {
	return x >= 0 && x < width && y >= 0 && y < width
}

// DivideCellSpherical takes as input a SphereCell object and returns two SphereCell objects
// children cells inherit same parameters of parent cell except for position and cellID (for second child)
// where the children cells are placed in random positions in the vicinity of the parent cell
func DivideCellSpherical(s *SphereCell) (*SphereCell, *SphereCell) {
	theta := rand.Float64() * 2 * math.Pi
	child1 := &SphereCell{}
	child2 := &SphereCell{}
	child1.position.x = (s.position.x + s.radius*math.Cos(theta)/2)
	child2.position.x = (s.position.x + s.radius*math.Cos(theta-math.Pi)/2)
	child1.position.y = (s.position.y + s.radius*math.Sin(theta)/2)
	child2.position.y = (s.position.y + s.radius*math.Sin(theta-math.Pi)/2)
	child1.radius, child2.radius = s.radius/2, s.radius/2
	child1.red, child1.green, child1.blue = s.red, s.green, s.blue
	child2.red, child2.green, child2.blue = s.red, s.green, s.blue
	child1.velocity.x, child1.velocity.y = s.velocity.x, s.velocity.y
	child2.velocity.x, child2.velocity.y = s.velocity.x, s.velocity.y
	child1.acceleration.x, child1.acceleration.y = s.acceleration.x, s.acceleration.y
	child2.acceleration.x, child2.acceleration.y = s.acceleration.x, s.acceleration.y
	child1.cellID = s.cellID
	child2.cellID = rand.Intn(1000000)
	return child1, child2
}

// GrowCellSpherical takes as input a SphereCell object and a growthRate float64 parameter
// It returns the updated radius of the SphereCell object after growing by growthRate
func GrowCellSpherical(s *SphereCell, growthRate float64) float64 {
	return s.radius + growthRate*s.radius
}

// UpdateAcceleration takes as input a Culture object and a particular cell within the Culture
// It returns the net acceleration due to the net forces that a Cell experiences at a given point in time.
func UpdateAcceleration(currentCulture Culture, s *SphereCell) OrderedPair {
	var accel OrderedPair
	//ADD NET FORCE CALCULATION HERE
	accel.x = 0
	accel.y = 0
	return accel
}

// UpdateVelocity takes as input a Cell object and a float time.
// It returns the updated velocity of that Cell estimated over time seconds as a result of Newtonian physics
func UpdateVelocity(s *SphereCell, time float64) OrderedPair {
	var vel OrderedPair
	vel.x = s.velocity.x + s.acceleration.x*time
	vel.y = s.velocity.y + s.acceleration.y*time
	return vel
}

// UpdatePosition takes as input a Cell object and a float time.
// It returns the updated position (in components) of that Cell estimated over time seconds as a result of Newtonian physics
func UpdatePosition(s *SphereCell, time float64) OrderedPair {
	var pos OrderedPair
	pos.x = s.position.x + s.velocity.x*time + 0.5*s.acceleration.x*time*time
	pos.y = s.position.y + s.velocity.y*time + 0.5*s.acceleration.y*time*time
	return pos
}

// CopyCulture returns a different culture that has same cells as the input culture
func CopyCulture(culture Culture) Culture {
	var newCulture Culture
	//newCulture.cells = make([]*RodCell, len(culture.cells))
	newCulture.cells = make([]*SphereCell, len(culture.cells))
	for i := range newCulture.cells {
		//newCulture.cells[i] = CopyCell(culture.cells[i])
		newCulture.cells[i] = CopySphereCell(culture.cells[i])
	}
	//copy nutrition board
	newCulture.nutrition = CopyNutritionBoard(culture.nutrition)

	newCulture.width = culture.width
	newCulture.particles = make([]*SignalParticle, len(culture.particles))
	for i := range newCulture.particles {
		newCulture.particles[i] = CopyParticle(culture.particles[i])
	}
	return newCulture
}

// CopyNutritionBoard returns a new nutrition board that has same values as the input nutrition board
func CopyNutritionBoard(nutritionBoard [][]int) [][]int {
	newNutritionBoard := make([][]int, len(nutritionBoard))
	for i := range newNutritionBoard {
		newNutritionBoard[i] = make([]int, len(nutritionBoard[i]))
		for j := range newNutritionBoard[i] {
			newNutritionBoard[i][j] = nutritionBoard[i][j]
		}
	}
	return newNutritionBoard
}

// CopyParticle returns a new particle that has same values as the input particle
func CopyParticle(p *SignalParticle) *SignalParticle {
	var newParticle SignalParticle
	newParticle.position.x = p.position.x
	newParticle.position.y = p.position.y
	newParticle.velocity.x = p.velocity.x
	newParticle.velocity.y = p.velocity.y
	return &newParticle
}

// CheckSphereCollision is a function that takes as input a culture of spherical cells
// It iterates over every pair of cells within the culture and performs collision detection
func CheckSphereCollision(newCulture Culture) {
	for i := 0; i < len(newCulture.cells); i++ {
		for j := 0; j < len(newCulture.cells); j++ {
			//Check if cells to collide are not the original cell
			if newCulture.cells[i].cellID != newCulture.cells[j].cellID {

				//Check if cells overlap
				if CheckSphereOverlap(newCulture.cells[i], newCulture.cells[j]) {
					//If cells do overlap, shove both cells
					ShoveOverlapCells(newCulture.cells[i], newCulture.cells[j])
					//newCulture.cells[i].red, newCulture.cells[i].green, newCulture.cells[i].blue = 255, 255, 255
					//newCulture.cells[j].red, newCulture.cells[j].green, newCulture.cells[j].blue = 255, 255, 255
				}
			}
		}
	}
}

// CheckSphereOverlap is a function that takes as input two SphereCell objects.
// It returns true if the two cells are determined to be overlapping, and false if otherwise.
func CheckSphereOverlap(s1, s2 *SphereCell) bool {
	//Mathematically, if the distance between the two cells exceeds 2*radius of a cell
	//Two spherical cells would be in contact
	if Distance(s1.position, s2.position) < (s1.radius + s2.radius) {
		return true
	}
	return false
}

// GetOverlap is a function that takes as input two SphereCell objects.
// It returns the amount that two overlapping cells are overlapping by as a float value.
func GetOverlap(s1, s2 *SphereCell) float64 {
	return Distance(s1.position, s2.position) - s1.radius - s2.radius
}

// Distance takes two OrderedPairs representing the position of two cells in 2D space.
// It returns the distance between these two points as a float value
func Distance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// ShoveOverlapCells is a function that takes as input two SphereCell objects
// If the two cells are overlapping, the cells will be shoved apart at an angle
// directly opposite to the point of contact.
func ShoveOverlapCells(s1, s2 *SphereCell) {
	//IMPORTANT, order of updating position matters
	//Must first update the "original" cell s1
	overlap := 0.5 * GetOverlap(s1, s2)
	separation := Distance(s1.position, s2.position)
	s1.position.x -= overlap * (s1.position.x - s2.position.x) / separation
	s1.position.y -= overlap * (s1.position.y - s2.position.y) / separation

	s2.position.x += overlap * (s1.position.x - s2.position.x) / separation
	s2.position.y += overlap * (s1.position.y - s2.position.y) / separation

}

// CopySphereCell returns a different cell that has same fields as the input cell
func CopySphereCell(cell *SphereCell) *SphereCell {
	//newCell := &RodCell{}
	var newCell SphereCell
	newCell.position.x = cell.position.x
	newCell.position.y = cell.position.y
	newCell.velocity.x = cell.velocity.x
	newCell.velocity.y = cell.velocity.y
	newCell.acceleration.x = cell.acceleration.x
	newCell.acceleration.y = cell.acceleration.y

	newCell.cellID = cell.cellID
	newCell.radius = cell.radius
	newCell.red = cell.red
	newCell.blue = cell.blue
	newCell.green = cell.green

	return &newCell
}

//-- this function is not called anywhere in the code ---
// GetMidPoint takes as input two OrderedPairs
// It returns the middle point on a Cartesian plane that is between the two OrderedPairs
// func GetMidPoint(pointOne, pointTwo OrderedPair) OrderedPair {
// 	var midPoint OrderedPair
// 	midPoint.x = (pointOne.x + pointTwo.x) / 2
// 	midPoint.y = (pointOne.y + pointTwo.y) / 2
// 	return midPoint
// }

//--- this function is not called anywhere in the code ---
// Shove() is a method of a SphereCell that takes as input another overlapping SphereCell s2
// It updates the position of the SphereCell by pushing the cell away from the overlapping cell
// func (s *SphereCell) Shove(s2 *SphereCell) {
// 	//Get overlap between the two cells
// 	overlap := GetOverlap(s, s2)
// 	separation := Distance(s.position, s2.position)
// 	s.position.x -= overlap * (s.position.x - s2.position.x) / separation
// 	s.position.y -= overlap * (s.position.y - s2.position.y) / separation
// }

// ReleaseSignals generates numParticles of SignalParticle with velocities evenly distributed in all directions.
// Each particle is positioned outiside from the cell, moving away at a speed of particleSpeed.
func (cell SphereCell) ReleaseSignals(particleSpeed float64, numParticles int) []*SignalParticle {
	particles := make([]*SignalParticle, numParticles)

	for i := 0; i < numParticles; i++ {
		angle := 2 * math.Pi * float64(i) / float64(numParticles)
		velocityX := math.Cos(angle) * particleSpeed
		velocityY := math.Sin(angle) * particleSpeed

		particles[i].velocity = OrderedPair{velocityX, velocityY}
		particles[i].position = OrderedPair{
			cell.position.x + (cell.radius+1)*math.Cos(angle),
			cell.position.y + (cell.radius+1)*math.Sin(angle),
		}
	}
	return particles
}

// ReceiveSignals checks if a cell should receive any signal in the current time step
// if so, it will move towrad the driection of the signals
func (cell *SphereCell) ReceiveSignals(particles []*SignalParticle) {
// if so, it will move towrad the direction of the signals
func (cell SphereCell) ReceiveSignals(particles []SignalParticle) {
	for i := range particles {
		if cell.CloseTo(particles[i].position) {
			cell.MoveToward(particles[i].position)
			particles = append(particles[:i], particles[i+1:]...)
		}
	}
}

// CloseTo detemines if a position is close to a cell. If the position is in the square tangent to the cell,
// it is considered as close.
func (cell SphereCell) CloseTo(position OrderedPair) bool {
	return position.x >= cell.position.x-cell.radius && position.x < cell.position.x+cell.radius &&
		position.y >= cell.position.y-cell.radius && position.y < cell.position.y+cell.radius
}

// MoveToward is a method of cell that makes cell move towards the input position
func (cell *SphereCell) MoveToward(position OrderedPair) {
	// Calculate the direction vector
	directionX := position.x - cell.position.x
	directionY := position.y - cell.position.y

	// Normalize the direction vector to get the unit direction
	magnitude := math.Sqrt(directionX*directionX + directionY*directionY)
	if magnitude == 0 {
		return // Avoid division by zero; the cell is already at the position
	}

	unitDirectionX := directionX / magnitude
	unitDirectionY := directionY / magnitude

	// Set the cell's velocity in the direction of the target position
	// Assuming you want to set the velocity equal to the unit direction vector
	cell.velocity = OrderedPair{unitDirectionX, unitDirectionY}
}

// // UpdateParticle takes as input a Particle object and a float time.
// It returns the updated position (in components) of that Particle estimated over time seconds as a result of Newtonian physics
func UpdateParticle(p *SignalParticle, time float64) OrderedPair {
	var pos OrderedPair
	pos.x = p.position.x + p.velocity.x*time
	pos.y = p.position.y + p.velocity.y*time
	return pos
}

// RandomDiffusion simulate the random diffusion as Brownian motion.
// It updates the velocity of one SphereCell by changing the direction randomly but the magnitude of the celocity remains the same.
func (cell *SphereCell) RandomDiffusion() {

	// Generate a random angle in radians
	// angle is uniformly distributed between 0 and 2Pi (0 to 360 degrees)
	angle := rand.Float64() * 2 * math.Pi

	// Calculate the original speed of the cell
	originalSpeed := math.Sqrt(cell.velocity.x*cell.velocity.x + cell.velocity.y*cell.velocity.y)

	// Update the velocity based on angle and the current speed
	cell.velocity.x = math.Cos(angle) * originalSpeed
	cell.velocity.y = math.Sin(angle) * originalSpeed
}
