package main

// RodCell represents out rod-shaped cells. Length is the length of the rectangle part.
// Width is the width of the rectangle part and also width/2 is the radius of circle part.
// Position is the position of the bottom-left corner of the rectangle of the cell.
// type RodCell struct {
// 	length, width, angle, maxLength float64
// 	red, green, blue                uint8
// 	position                        OrderedPair
// }

// SphereCell represents a single spherical bacterial cell. It contains fields representing the movement of
// a single cell in 2D space such as position, velocity, acceleration. It also contains the radius of the cell
// necessary for drawing and collision detection. It also contains a cellID that will allow certain "clusters" of
// spherical cells to form and communicate.
type SphereCell struct {
	cellID                           int
	radius                           float64
	position, velocity, acceleration OrderedPair
	red, green, blue                 uint8
	cellNutrition                    float64
}

// OrderedPair contains two float64 fields corresponding to
// the x and y coordinates of a point or vector in two-dimensional space.
type OrderedPair struct {
	x, y float64
}

// Culture represents all the cells in the biofilm at the time.
type Culture struct {
	//cells []*RodCell
	cells []*SphereCell

	//Width represents the width of the "canvas" where cells grow
	width float64

	//nutrition represents the nutrition board of the culture
	nutrition NutritionBoard //square 2D board with length = width

	//particles represents all signal particles
	particles []*SignalParticle
}

// SignalPaticle represents the signal particles released by cells, cells tend to move toward the direction they
// received the paticles
type SignalParticle struct {
	position, velocity OrderedPair
}

type NutritionBoard [][]int
