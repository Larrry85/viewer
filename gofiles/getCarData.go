// getCarData.go
package gofiles

import (
	"encoding/json"		// encoding and decoding JSON data
	"errors"			// error handling
	"fmt"				// formatting strings
	"html/template"		// rendering HTML templates
	"net/http"			// handling HTTP requests and responses
	"os"				// file operations
	"sort"				// sorting slices
	"strconv"			// string conversions
	"strings"			// string manipulation
)

// This file handling car data, fetching car details from an API,
// and filtering cars based on certain criteria

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// fetches detailed information about a car from an API
func GetCarDetails(idStr string) (Car, error) {
	// converts the car ID from a string to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Car{}, errors.New("invalid car ID")
	}

	// Make a request to the Go server endpoint to fetch car details
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/models/%d", id))
	if err != nil {
		return Car{}, err
	}
	defer resp.Body.Close()

	// decode the response body into a Car struct
	var car Car
	if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
		return Car{}, err
	}
		// populates the ManufacturerName and CategoryName fields of the Car struct
		// Fetch manufacturer details
		manufacturerResp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/manufacturers/%d", car.ManufacturerID))
		if err != nil {
			return Car{}, err
		}
		defer manufacturerResp.Body.Close()
	
		var manufacturer Manufacturer
		if err := json.NewDecoder(manufacturerResp.Body).Decode(&manufacturer); err != nil {
			return Car{}, err
		}
		car.ManufacturerName = manufacturer.Name
	
		// Fetch category details
		categoryResp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/categories/%d", car.CategoryID))
		if err != nil {
			return Car{}, err
		}
		defer categoryResp.Body.Close()
	
		var category Category
		if err := json.NewDecoder(categoryResp.Body).Decode(&category); err != nil {
			return Car{}, err
		}
		car.CategoryName = category.Name
	
	return car, nil
	// returns a Car struct with detailed information about a specific car model,
	// including the manufacturer's name and the category's name
}

// handles HTTP requests for car details, fetches the details, and renders them using an HTML template
func CarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// extracts the id from the query parameters
	id := r.URL.Query().Get("id")

	// fetch car details based on the ID
	car, err := GetCarDetails(id) // id is passed to the GetCarDetails()
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound) // 404
		return //r eturned values are assigned to the variables car
	}

	// parses the carDetails.html template
	tmpl, err := template.ParseFiles("static/carDetails.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
		return
	}

	// executes the template with the fetched car data
	if err := tmpl.Execute(w, car); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
		return
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// finds a manufacturer by its ID from a slice of manufacturers
func FindManufacturerByID(id int, manufacturers []Manufacturer) (Manufacturer, error) {
	// iterates through the slice of manufacturers
	for _, manufacturer := range manufacturers {
		if manufacturer.ID == id {
			return manufacturer, nil
		}
	} // returns the manufacturer if the ID matches
	return Manufacturer{}, fmt.Errorf("manufacturer not found")
}

// finds a category by its ID from a slice of categories
func FindCategoryByID(id int, categories []Category) (Category, error) {
	// iterates through the slice of categories
	for _, category := range categories {
		if category.ID == id {
			return category, nil
		}
	} // returns the category if the ID matches
	return Category{}, fmt.Errorf("category not found")
}

// populates the ManufacturerName and CategoryName fields of a car
func PopulateCarDetails(car *Car, data CarsData) error {

	// uses FindManufacturerByID and FindCategoryByID to get the names
	manufacturer, err := FindManufacturerByID(car.ManufacturerID, data.Manufacturers)
	if err != nil {
		return err
	} //  sets car.ManufacturerName to the name of the found manufacturer
	car.ManufacturerName = manufacturer.Name

	category, err := FindCategoryByID(car.CategoryID, data.Categories)
	if err != nil {
		return err
	} // sets car.CategoryName to the name of the found category
	car.CategoryName = category.Name

	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// handles filter requests, filters car data based on query parameters, and renders the filtered data using an HTML template
func FilterPage(w http.ResponseWriter, r *http.Request) {
    data, err := os.ReadFile(filePath) // reads car data from a file
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to read file %s: %v", filePath, err), http.StatusInternalServerError) // 500
        return
    }

    var carData CarsData // unmarshals the JSON data into a CarsData struct
    if err := json.Unmarshal(data, &carData); err != nil {
        http.Error(w, fmt.Sprintf("Failed to parse JSON: %v", err), http.StatusInternalServerError) // 500
        return
    }

	// extracts filter parameters from the query string
    manufacturer := r.URL.Query().Get("manufacturer")
    year := r.URL.Query().Get("year")
    category := r.URL.Query().Get("category")

	// filters car data using the filterCars()
    filteredCars := filterCars(carData, manufacturer, year, category)
    uniqueYears := getUniqueYears(carData.CarModels)

	// if no matches while filtering cars
    var message string
    if len(filteredCars) == 0 {
        message = "No cars found, try again!"
    }

	// prepares the data for the template and...
    filteredData := struct {
        CarModels     []Car
        Categories    []Category
        Manufacturers []Manufacturer
        Years         []int
        Message       string
    }{
        CarModels:     filteredCars,
        Categories:    carData.Categories,
        Manufacturers: carData.Manufacturers,
        Years:         uniqueYears,
        Message:       message,
	}

	// ...renders the filtered.html template with the filtered data
    tmpl, err := template.ParseFiles("static/filtered.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) // 500
        return
    }
    if err := tmpl.Execute(w, filteredData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) // 500
    }
}

// filters car models based on manufacturer, year, and category
func filterCars(data CarsData, manufacturer, year, category string) []Car {
	var filtered []Car
	// iterates through the car models
	for _, car := range data.CarModels { // checks if each car matches the filter criteria
		if (manufacturer == "" || strings.EqualFold(data.Manufacturers[car.ManufacturerID-1].Name, manufacturer)) &&
			(year == "" || fmt.Sprintf("%d", car.Year) == year) &&
			(category == "" || strings.EqualFold(data.Categories[car.CategoryID-1].Name, category)) {
			PopulateCarDetails(&car, data)
			filtered = append(filtered, car)
		}
	} // populates the car details and adds the matching cars to the filtered list
	return filtered
}

// extracts unique years from a list of cars.
func getUniqueYears(cars []Car) []int {
	// map where the keys are the years from the cars slice
	yearMap := make(map[int]struct{})
	// iterates over each car in the cars slice
	for _, car := range cars {
		yearMap[car.Year] = struct{}{}
	} // for each car, it adds the car.Year to yearMap. If the year is already in the map,
	// it won't add a duplicate because map keys are unique

	var years []int
	// iterates over the keys in yearMap
	for year := range yearMap {
		years = append(years, year) // appends each unique year to the years slice
	}
	sort.Ints(years) // sorts the years slice in ascending order
	return years // sorted years slice is returned
}