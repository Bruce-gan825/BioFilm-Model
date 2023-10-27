package main

// rodCell represents out rod-shaped cells. Length is the length of the rectangle part.
// Width is the width of the rectangle part and also width/2 is the radius of circle part.
// Position is the position of the bottom-left corner of the rectangle of the cell.
type rodCell struct {
	length, width, angle float64
	position             OrderedPair
}

// OrderedPair contains two float64 fields corresponding to
// the x and y coordinates of a point or vector in two-dimensional space.
type OrderedPair struct {
	x, y float64
}
