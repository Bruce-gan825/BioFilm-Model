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
	growthRate := 0.03
	maxRadius := 20.0

	//Iterate over all Cells in the newly created Culture and update their fields
	for i := range newCulture.cells {
		//Update position functions go here
		newCulture.cells[i].acceleration = UpdateAcceleration(currentCulture, newCulture.cells[i])
		newCulture.cells[i].velocity = UpdateVelocity(newCulture.cells[i], time)
		newCulture.cells[i].position = UpdatePosition(newCulture.cells[i], time)

		//grow cells
		newCulture.cells[i].radius = GrowCellSpherical(newCulture.cells[i], growthRate)

		if newCulture.cells[i].radius >= maxRadius {
			child1, child2 := DivideCellSpherical(newCulture.cells[i])
			newCulture.cells[i] = child1
			newCulture.cells = append(newCulture.cells, child2)
		}
	}

	// for i := range newCulture.cells {
	// 	if newCulture.cells[i].radius >= maxRadius {
	// 		child1, child2 := DivideCellSpherical(newCulture.cells[i])
	// 		newCulture.cells[i] = child1
	// 		newCulture.cells = append(newCulture.cells, child2)
	// 	}

	// }

	//Apply simple collision function for the newCulture
	CheckSphereCollision(newCulture)

	return newCulture

}

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

// GetRectPoints takes as input the origin (centre point) of a RodCell and a rotation angle Theta
// It returns the top left, top right, bottom left, bottom right points of the body rectangle
func GetRectPoints(center OrderedPair, width, length, theta float64) []OrderedPair {
	vertices := make([]OrderedPair, 4)

	//Top Left
	vertices[0].x = center.x - (length/2)*math.Cos(theta) - (width/2)*math.Sin(theta)
	vertices[0].y = center.y - (length/2)*math.Sin(theta) + (width/2)*math.Cos(theta)
	//Top Right
	vertices[1].x = center.x + (length/2)*math.Cos(theta) - (width/2)*math.Sin(theta)
	vertices[1].y = center.y + (length/2)*math.Sin(theta) + (width/2)*math.Cos(theta)
	//Bottom Left
	vertices[2].x = center.x - (length/2)*math.Cos(theta) + (width/2)*math.Sin(theta)
	vertices[2].y = center.y - (length/2)*math.Sin(theta) - (width/2)*math.Cos(theta)
	//Bottom Right
	vertices[3].x = center.x + (length/2)*math.Cos(theta) + (width/2)*math.Sin(theta)
	vertices[3].y = center.y + (length/2)*math.Sin(theta) - (width/2)*math.Cos(theta)

	return vertices
}

// GetMidPoint takes as input two OrderedPairs
// It returns the middle point on a Cartesian plane that is between the two OrderedPairs
func GetMidPoint(pointOne, pointTwo OrderedPair) OrderedPair {
	var midPoint OrderedPair
	midPoint.x = (pointOne.x + pointTwo.x) / 2
	midPoint.y = (pointOne.y + pointTwo.y) / 2
	return midPoint
}

// Elongate() is a method of Rodcell that takes as input a length, and the cell will elongate that length
func (c *RodCell) Elongate(length float64) {
	cellLength := c.length + length
	c.length = cellLength
}

// Divide() is a method of Rodcell as it returns two childen of the parent cell after division
// two children should have have the sizes as their paretn and same other fields
func (c *RodCell) Divide() (*RodCell, *RodCell) {
	child1 := &RodCell{}
	child2 := &RodCell{}
	child1.position.x = c.position.x - ((c.width/2)+((c.maxLength-c.width)/2))*math.Cos(c.angle)
	child2.position.x = c.position.x + ((c.width/2)+((c.maxLength-c.width)/2))*math.Cos(c.angle)
	child1.position.y = c.position.y - ((c.width / 2) + ((c.maxLength-c.width)/2)*math.Sin(c.angle))
	child2.position.y = c.position.y + ((c.width / 2) + ((c.maxLength-c.width)/2)*math.Sin(c.angle))
	child1.maxLength, child2.maxLength = c.maxLength, c.maxLength
	child1.red, child1.green, child1.blue = c.red, c.green, c.green
	child2.red, child2.green, child2.blue = c.red, c.green, c.green
	child1.length = (c.length - c.width) / 2
	child2.length = (c.length - c.width) / 2
	child1.width, child2.width = c.width, c.width
	child1.angle, child2.angle = c.angle, c.angle
	return child1, child2

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
