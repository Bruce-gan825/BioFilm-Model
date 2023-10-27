package main

import (
	"canvas"
	"image"
)

// AnimateSystem takes a slice of Culture objects along with a canvas width parameter and frequency parameter
// parameter and generates a slice of images corresponding to drawing each Culture
// on a canvasWidth x canvasWidth canvas only if its index in the Culture slice is divisible by drawing frequency
func AnimateSystem(timePoints []Culture, canvasWidth, drawingFrequency int) []image.Image {
	//Create empty slice of images representing the frames of the animation
	images := make([]image.Image, 0)
	//Iterate over each frame of the simulation
	for i := range timePoints {
		//If i is divisible by drawingFrequency, draw it and append image to slice of images
		if i%drawingFrequency == 0 {
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}
	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing a Culture
// object's cells on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(culture Culture, canvasWidth int) image.Image {
	// Create a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// Set black background of designated canvasWidth
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// Iterate over each individual Boid and draw as a circle of pixel radius 5
	for _, b := range culture.cells {
		//Retrieve the randomized colour of each Boid from its fields
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		centreX := (b.position.x / culture.width) * float64(canvasWidth)
		centreY := (b.position.y / culture.width) * float64(canvasWidth)
		r := 5.0
		c.Circle(centreX, centreY, r)
		c.Fill()
	}
	// Return the image created
	return c.GetImage()
}
