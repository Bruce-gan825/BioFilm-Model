package main

/*
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

*/
