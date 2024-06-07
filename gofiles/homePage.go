package gofiles

import (
	"encoding/json"		// decoding JSON data
	"fmt"				// formatting strings
	"html/template"		// rendering HTML templates
	"net/http"			// handling HTTP requests and responses
	"os"				// file operations
	"path/filepath"		// manipulating file paths
)

// data.json path
var directory = "api"
var fileName = "data.json"
var filePath = filepath.Join(directory, fileName) // full path to the JSON file

// HTTP handler that will serve the homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // checks if the URL path is "/"
		http.Error(w, "404 not found.", http.StatusNotFound) // 404
		return
	}

	// reads the content of the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file %s: %v", filePath, err), http.StatusInternalServerError) // 500
		return
	}

	// parses the JSON data into a CarsData struct
	var carData CarsData
	if err := json.Unmarshal(data, &carData); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse JSON: %v", err), http.StatusInternalServerError) // 500
		return
	}

	// calls the getUniqueYears function to extract unique years from the car models
	uniqueYears := getUniqueYears(carData.CarModels)

	// combines CarsData and unique years into a struct to pass to the template
	dataWithYears := struct {
		CarsData
		Years []int
	}{
		CarsData: carData,
		Years:    uniqueYears,
	}

	// parses the index.html template
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	} // executes the template with the combined data (dataWithYears)
	if err := tmpl.Execute(w, dataWithYears); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}
}