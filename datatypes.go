package main

// RodCell represents out rod-shaped cells. Length is the length of the rectangle part.
// Width is the width of the rectangle part and also width/2 is the radius of circle part.
// Position is the position of the bottom-left corner of the rectangle of the cell.
type RodCell struct {
	length, width, angle float64
	red, green, blue     uint8
	position             OrderedPair
}

// OrderedPair contains two float64 fields corresponding to
// the x and y coordinates of a point or vector in two-dimensional space.
type OrderedPair struct {
	x, y float64
}

// Culture represents all the cells in the biofilm at the time.
type Culture struct {
	cells []*RodCell

	//Width represents the width of the "canvas" where cells grow
	width float64
}
