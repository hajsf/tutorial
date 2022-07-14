package main

import (
	"encoding/csv"
	"io"
	"os"
)

func readCSVFile(filePath string) (persons []Person) {
	isFirstRow := true
	headerMap := make(map[string]int)

	// Load a csv file.
	f, _ := os.Open(filePath)

	// Create a new reader.
	r := csv.NewReader(f)
	for {
		// Read row
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		checkError("Some other error occurred", err)

		// Handle first row case
		if isFirstRow {
			isFirstRow = false

			// Add mapping: Column/property name --> record index
			for i, v := range record {
				headerMap[v] = i
			}

			// Skip next code
			continue
		}

		// Create new person and add to persons array
		persons = append(persons, Person{
			Firstname: record[headerMap["Firstname"]],
			Lastname:  record[headerMap["Lastname"]],
			Country:   record[headerMap["Country"]],
		})
	}
	return
}
