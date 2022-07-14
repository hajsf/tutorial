package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

func readCSVHeader(filePath string) (headerMap map[string]int) {
	headerMap = make(map[string]int)

	// Load a csv file.
	f, _ := os.Open(filePath)

	// Create a new reader.
	r := csv.NewReader(f)

	// Read first row only
	record, err := r.Read()

	checkError("Some other error occurred", err)

	// Add mapping: Column/property name --> record index
	for i, v := range record {
		headerMap[v] = i
	}

	// The map will be sorted by header name alphapatically, we need to sort it based on header index, i.e. by value
	temp := map[int][]string{}
	var a []int
	for k, v := range headerMap {
		temp[v] = append(temp[v], k)
	}
	for k := range temp {
		a = append(a, k)
	}

	// sort in increasing order, if required to sort in decreasing order use: sort.Sort(sort.Reverse(sort.IntSlice(a))) if need reverse sorting
	sort.Ints(a)
	for _, k := range a {
		for _, s := range temp[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	}

	return
}
