// getCarData.go
package gofiles

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

// Handles detailed car data fetching and rendering.

// Fetches car details from an API endpoint.
// GetCarDetails fetches car details from another API endpoint.
func GetCarDetails(idStr string) (Car, error) {
	// Convert the string ID to an integer
	id, err := strconv.Atoi(idStr) // tämän pitäis muuttaa string to int
	if err != nil {
		return Car{}, errors.New("invalid car ID")
	}

	// Make a request to the Go server endpoint to fetch car details
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/models/%d", id))
	if err != nil {
		return Car{}, err
	}
	defer resp.Body.Close()

	// Decode the response body into a Car struct
	var car Car
	if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
		return Car{}, err
	}

	return car, nil // ...lähettää sen auton data carDetails.html sivulle
}

// CarDetailsHandler extracts the car ID from query parameters,
// fetches car details, and renders them using carDetails.html.
// Handles requests for car details and renders an HTML template with the fetched data.
// Tämä sitä varten kun avaa details napin ja carDetails.html sivun
func CarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the car ID from the query parameters
	id := r.URL.Query().Get("id") // tämä id tulee html linkistä

	// Perform any necessary validation on the ID

	// Fetch car details based on the ID
	car, err := GetCarDetails(id) // tämä kutsuu GetCarDeatailsByID(), mukana id
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	// Render the car details page using a template
	tmpl, err := template.ParseFiles("static/carDetails.html") // oikean auton data tällä sivulla
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the car data
	if err := tmpl.Execute(w, car); err != nil { // oikean auton data tällä sivulla
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

/*
// Sends an HTTP GET request to http://localhost:3000/api
// to fetch the car data served by the ApiServer
func GetCarData() ([]Car, error) {
	resp, err := http.Get("http://localhost:3000/api")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var carsData []Car
	err = json.NewDecoder(resp.Body).Decode(&carsData)
	if err != nil {
		return nil, err
	}
	return carsData, nil
}*/

// FILTER CARS ///////////////////
// FilterPage handler processes filter requests and renders filtered data.
func FilterPage(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file %s: %v", filePath, err), http.StatusInternalServerError)
		return
	}

	var carData CarsData
	if err := json.Unmarshal(data, &carData); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse JSON: %v", err), http.StatusInternalServerError)
		return
	}

	year := r.URL.Query().Get("year")
	category := r.URL.Query().Get("category")

	filteredCars := filterCars(carData.CarModels, year, category)
	filteredData := CarsData{
		CarModels:     filteredCars,
		Categories:    carData.Categories,
		Manufacturers: carData.Manufacturers,
	}

	// Parse and execute the filtered template
	tmpl, err := template.ParseFiles("static/filtered.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, filteredData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// filterCars filters car models based on year and category.
func filterCars(cars []Car, year, category string) []Car {
	var filtered []Car
	for _, car := range cars {
		if (year == "" || fmt.Sprintf("%d", car.Year) == year) &&
			(category == "" || car.Category == category) {
			filtered = append(filtered, car)
		}
	}
	return filtered
}
