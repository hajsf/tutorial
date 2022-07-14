package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func main() {
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("error: ", err)
	}

	fmt.Print(records)

	// Read csv file headers
	header := readCSVHeader("./input.csv")
	fmt.Println(header)

	// Read from file using the provided handlers

	// Read persons from 'input.csv'
	persons := readCSVFile("./input.csv")

	// Modify persons a bit and write into new array.
	var modifiedPersons []Person
	for _, person := range persons {
		person.Country = "AnotherCountry"
		modifiedPersons = append(modifiedPersons, person)
	}

	// Write modified persons into 'output.csv'
	writeCSVFile(modifiedPersons, "./output.csv")

}
