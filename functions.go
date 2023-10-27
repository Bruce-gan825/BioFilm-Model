package main

// SimulateBiofilm() takes as input a Culture object initialCulture, a number of generations parameter numGens, and a time interval timeStep.
// It simulate the biofilm over numGens generations starting with initialCulture in which the system is updated every timeStep units of time, and as a result it should return a slice of numGens + 1 Culture objects
func SimulateBiofilm(initialCulture Culture, numGens int, time float64) []Culture {
	timePoints := make([]Culture, numGens+1)
	timePoints[0] = initialCulture
	for i := 1; i < numGens+1; i++ {
		timePoints[i] = UpdateCulture(timePoints[i-1])
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
	}

	return newCulture
}

//test
