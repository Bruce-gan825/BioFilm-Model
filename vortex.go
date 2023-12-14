package main

import (
	"math"
)

// Biofilm contains a slice of SphereCell pointers, representing the cells that are part of this particular biofilm
type Biofilm struct {
	cells []*SphereCell
}

// Vortex describes the movement of each biofilm around it respective biofilm center
func (c *Culture) Vortex(angularVelocity float64) {

	for _, biofilm := range c.biofilms {
		biofilmCenter := biofilm.FindBiofilmCenter()

		// Rotate each cell in the biofilm around the biofilm center
		for _, cell := range biofilm.cells {
			// Calculate the relative position of the cell to the original biofilm center
			relativeX := cell.position.x - biofilmCenter.x
			relativeY := cell.position.y - biofilmCenter.y

			// Calculate the angle and distance from the biofilm center to the cell
			cellAngle := math.Atan2(relativeY, relativeX)
			cellDistance := math.Sqrt(relativeX*relativeX + relativeY*relativeY)

			// Update the cell's position
			// Rotate the cell around the biofilm center
			newCellAngle := cellAngle + angularVelocity
			cell.position.x = biofilmCenter.x + cellDistance*math.Cos(newCellAngle)
			cell.position.y = biofilmCenter.y + cellDistance*math.Sin(newCellAngle)
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
