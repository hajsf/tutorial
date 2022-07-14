package main

import (
	"encoding/csv"
	"os"
)

func writeCSVFile(persons []Person, outputPath string) {

	// Define header row
	headerRow := []string{
		"Firstname", "Lastname", "Country",
	}

	// Data array to write to CSV
	data := [][]string{
		headerRow,
	}

	// Add persons to output data
	for _, person := range persons {
		data = append(data, []string{
			// Make sure the property order here matches
			// the one from 'headerRow' !!!
			person.Firstname,
			person.Lastname,
			person.Country,
		})
	}

	// Create file
	file, err := os.Create(outputPath)
	checkError("Cannot create file", err)
	defer file.Close()

	// Create writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write rows into file
	for _, value := range data {
		err = writer.Write(value)
		checkError("Cannot write to file", err)
	}
}
