package main

/*
func IsInside(position OrderedPair) bool {

	// if the point falls inside the rectangle part of RodCell
	if isInsideRectangle(position) {
		return true
	}

	// if the point falls inside the semi-circle part of RodCell
	if IsInsideSemiCircle {

	}

	return false
}

func IsInsideSemiCircle(p OrderedPair) bool {

}

func IsInsideRectangle(p OrderedPair, rodcell RodCell) bool {
	// the half length of RodCell
	halfLength := 0.5 * RodCell.length
	// radius = half of the width
	radius := 0.5 * RodCell.width

	recCenter = RotatedRectangleCenter(rodcell)

	// check whether the point is in the rectangle
	return p.x >= recCenter.x-halfLength && p.x <= rectCenter.x+halfLength && p.y >= rectCenter.y-radius && p.y <= rectCenter.y+radius
}

func RotatedRectangleCenter(rodcell RodCell) OrderedPair {
	// radius = half of the width
	radius := 0.5 * rodcell.width

	unrotatedCenterX := rodcell.position.x + radius
	unrotatedCenterY := rodcell.position.y + radius

	// Translate center for rotation about the origin
	centerX := unrotatedCenterX - rodcell.position.x
	centerY := unrotatedCenterY - rodcell.position.y

	// Convert rotation angle from degrees to radians
	radians := rodcell.angle * math.Pi / 180

	// Apply rotation matrix to the translated center point
	rotatedCenterX := centerX*math.Cos(radians) - centerY*math.Sin(radians)
	rotatedCenterY := centerX*math.Sin(radians) + centerY*math.Cos(radians)

	rotatedCenterX += position.x
	rotatedCenterY += position.y

	return OrderedPair{X: rotatedCenterX, Y: rotatedCenterY}
}
*/
