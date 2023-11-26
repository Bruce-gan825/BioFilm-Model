package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNutritionBoardFromFile(filename string) [][]int {
	var nutritionBoard [][]int

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into individual numbers based on tab delimiter
		numStrings := strings.Fields(line)
		var row []int

		// Convert each string to an integer and append to the row
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return nil
			}
			row = append(row, num)
		}

		// Append the row to the nutritionBoard
		nutritionBoard = append(nutritionBoard, row)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return nutritionBoard
}
