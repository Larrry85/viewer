package gofiles

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// data.json path
var directory = "api"
var fileName = "data.json"
var filePath = filepath.Join(directory, fileName)

// HomePage reads data from data.json, parses it into CarsData,
// and renders it using index.html
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Reads and parses JSON data from data.json
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file %s: %v", filePath, err), http.StatusInternalServerError)
		return
	}

	// Parse json data into variables
	var carData CarsData
	if err := json.Unmarshal(data, &carData); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse JSON: %v", err), http.StatusInternalServerError)
		return
	}

	// Extract unique years
	uniqueYears := getUniqueYears(carData.CarModels)

	// Combine car data and unique years into a single struct
	dataWithYears := struct {
		CarsData
		Years []int
	}{
		CarsData: carData,
		Years:    uniqueYears,
	}

	// Render the HTML template with the car data
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, dataWithYears); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
