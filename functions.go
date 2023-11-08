package main

import (
	"fmt"
	"math"
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

// CheckSphereOverlap is a function that takes as input two SphereCell objects.
// It returns true if the two cells are determined to be overlapping, and false if otherwise.
func CheckSphereOverlap(s1, s2 SphereCell) bool {
	//Mathematically, if the distance between the two cells exceeds 2*radius of a cell
	//Two spherical cells would be in contact
	if Distance(s1.position, s2.position) > (s1.radius + s2.radius) {
		return true
	}
	return false
}

// Distance takes two OrderedPairs representing the position of two cells in 2D space.
// It returns the distance between these two points as a float value
func Distance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
