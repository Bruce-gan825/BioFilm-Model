package main

import (
	"math"
)

// Biofilm contains a slice of SphereCell pointers, representing the cells that are part of this particular biofilm
type Biofilm struct {
	cells []*SphereCell
}

// Vortex describes the movement of each biofilm around the center of culture
// cells in the same biofilm rotates together with the same angular velocity (seen a biofilm as a rigid body)
func (c *Culture) Vortex(angularVelocity float64) {
	cultureCenter := c.FindCultureCenter()

	for _, biofilm := range c.biofilms {
		biofilmCenter := biofilm.FindBiofilmCenter()

		// Calculate the components of the distance between the culture center and the biofilm center
		deltaX := biofilmCenter.x - cultureCenter.x
		deltaY := biofilmCenter.y - cultureCenter.y

		// Calculate the angle and distance from the culture center to the biofilm center
		// calculate the angle in radians between the positive x-axis of a plane and the point given by the coordinates (deltaX, deltaY)
		biofilmAngle := math.Atan2(deltaY, deltaX)
		biofilmDistance := math.Sqrt(deltaX*deltaX + deltaY*deltaY)

		// Rotate the biofilm center around the culture center
		newBiofilmAngle := biofilmAngle + angularVelocity
		newBiofilmCenter := OrderedPair{
			x: cultureCenter.x + biofilmDistance*math.Cos(newBiofilmAngle),
			y: cultureCenter.y + biofilmDistance*math.Sin(newBiofilmAngle),
		}

		// Update the position of each cell in the same biofilm
		for _, cell := range biofilm.cells {
			// Calculate the relative position of the cell to the original biofilm center
			relativeX := cell.position.x - biofilmCenter.x
			relativeY := cell.position.y - biofilmCenter.y

			// Update the cell's position, maintaining its relative position to the biofilm center
			cell.position.x = newBiofilmCenter.x + relativeX
			cell.position.y = newBiofilmCenter.y + relativeY
		}
	}
}

// IsInBiofilm divided all the cells into several biofilms
func (c *Culture) IsInBiofilm(cells []*SphereCell, biofilms []*Biofilm, threshold float64) {
	for _, cell := range cells {
		for _, biofilm := range biofilms {
			if IsInProximity(cell, biofilm, threshold) {
				biofilm.cells = append(biofilm.cells, cell)
				break // ensuring a cell can only belong to one biofilm
			}
		}
	}
}

// IsInProximity checks if the cell is within a certain distance (threshold) from the center of the biofilm
// it returns either true (is in proximity) or false (not in proximity).
func IsInProximity(cell *SphereCell, biofilm *Biofilm, threshold float64) bool {

	// find the center of the biofilm
	biofilmCenter := biofilm.FindBiofilmCenter()

	distance := Distance(cell.position, biofilmCenter)
	return distance <= threshold
}

func (c *Culture) FindCultureCenter() OrderedPair {
	var sumX, sumY float64
	var cultureCenter OrderedPair

	numCells := len(c.cells)

	if numCells == 0 {
		return OrderedPair{0, 0}
	}

	for _, cell := range c.cells {
		sumX += cell.position.x
		sumY += cell.position.y
	}

	// Averaging the sum to find the centroid
	cultureCenter.x = sumX / float64(numCells)
	cultureCenter.y = sumY / float64(numCells)

	return cultureCenter
}

// FindioFilmCenter returns the coordinates of the biofilm center
func (b *Biofilm) FindBiofilmCenter() OrderedPair {
	var sumX, sumY float64
	var biofilmCenter OrderedPair

	numCells := len(b.cells)

	if numCells == 0 {
		return OrderedPair{0, 0}
	}

	for _, cell := range b.cells {
		sumX += cell.position.x
		sumY += cell.position.y
	}

	// Calculating the average position
	biofilmCenter.x = sumX / float64(numCells)
	biofilmCenter.y = sumY / float64(numCells)

	return biofilmCenter
}
