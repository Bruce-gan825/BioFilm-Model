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

	// Iterate over each individual RodCell and draw as a rounded rectangle
	// Each rounded rectangle has width, height, and rounded corners with radius width
	for _, b := range culture.cells {
		//Retrieve the randomized colour of each RodCell from its fields
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))

		//Draw each RodCell's central rectangle first
		//Apply rotation based on angle of cell, assuming position refers to center of cell
		/*
			topX := ((b.position.x - b.length/2) / culture.width) * float64(canvasWidth)
			topY := ((b.position.y - b.width/2) / culture.width) * float64(canvasWidth)
			bottomX := ((b.position.x + b.length/2) / culture.width) * float64(canvasWidth)
			bottomY := ((b.position.y + b.width/2) / culture.width) * float64(canvasWidth)

			rotateTopX := (topX-b.position.x)*math.Cos(b.angle) - (topY-b.position.y)*math.Sin(b.angle)
			rotateTopY := (topX-b.position.x)*math.Sin(b.angle) - (topY-b.position.y)*math.Cos(b.angle)

			rotateBottomX := (bottomX-b.position.x)*math.Cos(b.angle) - (bottomY-b.position.y)*math.Sin(b.angle)
			rotateBottomY := (bottomX-b.position.x)*math.Sin(b.angle) - (bottomY-b.position.y)*math.Cos(b.angle)
			c.ClearRect(int(rotateTopX+b.position.x), int(rotateTopY+b.position.y), int(rotateBottomX+b.position.x), int(rotateBottomY+b.position.y))
				c.ClearRect(200, 200, 100, 400)
				//Draw a circle at the ends of the central rectangle
				//REMEMBER TO ADJUST CIRCLE X/Y POSITION WHEN ROTATING
				c.Circle(topX, topY+(b.width/2), b.width/2)
				c.Circle(topX+b.length, topY+(b.width/2), b.width/2)

		*/
		//Draw the base rectangle
		vertices := GetRectPoints(b.position, b.width, b.length, b.angle)
		c.MoveTo(vertices[0].x, vertices[0].y)
		c.LineTo(vertices[1].x, vertices[1].y)
		c.LineTo(vertices[3].x, vertices[3].y)
		c.LineTo(vertices[2].x, vertices[2].y)
		c.LineTo(vertices[0].x, vertices[0].y)
		c.Fill()

		//Draw two circles at the end of the base rectangle
		c.Circle(GetMidPoint(vertices[0], vertices[2]).x, GetMidPoint(vertices[0], vertices[2]).y, b.width/2)
		c.Circle(GetMidPoint(vertices[1], vertices[3]).x, GetMidPoint(vertices[1], vertices[3]).y, b.width/2)
		c.Fill()

	}
	// Return the image created
	return c.GetImage()
}
