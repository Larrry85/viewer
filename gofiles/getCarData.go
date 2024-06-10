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

type carResult struct {
	car Car
	err error
}

type manufacturerResult struct {
	manufacturer Manufacturer
	err          error
}

type categoryResult struct {
	category Category
	err      error
}

// GetCarDetails utilizes goroutines and channels to fetch car details, manufacturer
// details, and category details concurrently, improving the performance of the 
// function by making asynchronous API calls

// fetches detailed information about a car from an API using channels for concurrency
func GetCarDetails(idStr string) (Car, error) {
	// idStr as input, representing the ID of the car to fetch
	id, err := strconv.Atoi(idStr) // idStr to an integer (id
	if err != nil {
		return Car{}, errors.New("invalid car ID")
	}

	carCh := make(chan carResult) //  car details
	manufacturerCh := make(chan manufacturerResult) // manufacturer details
	categoryCh := make(chan categoryResult) // category details

	go func() {
		// GET request to the API endpoint for car details
		resp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/models/%d", id))
		if err != nil {
			carCh <- carResult{Car{}, err} // error result to the carCh channel
			return
		}
		defer resp.Body.Close()

		var car Car
		if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
			carCh <- carResult{Car{}, err}
			return
		}

		carCh <- carResult{car, nil}
	}()

	var car Car
	var manufacturer Manufacturer
	var category Category

	// Car details are received from the carCh channel
	carResult := <-carCh
	if carResult.err != nil {
		return Car{}, carResult.err
	}
	car = carResult.car // it assigns the received car details to the car variable

	go func() {
		// Fetch manufacturer details
		resp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/manufacturers/%d", car.ManufacturerID))
		if err != nil {
			manufacturerCh <- manufacturerResult{Manufacturer{}, err}
			return
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&manufacturer); err != nil {
			manufacturerCh <- manufacturerResult{Manufacturer{}, err}
			return
		}

		manufacturerCh <- manufacturerResult{manufacturer, nil}
	}()

	go func() {
		// Fetch category details
		resp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/categories/%d", car.CategoryID))
		if err != nil {
			categoryCh <- categoryResult{Category{}, err}
			return
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
			categoryCh <- categoryResult{Category{}, err}
			return
		}

		categoryCh <- categoryResult{category, nil}
	}()

	// Receive manufacturer details
	manufacturerResult := <-manufacturerCh
	if manufacturerResult.err != nil {
		return Car{}, manufacturerResult.err
	}
	manufacturer = manufacturerResult.manufacturer

	// Receive category details
	categoryResult := <-categoryCh
	if categoryResult.err != nil {
		return Car{}, categoryResult.err
	}
	category = categoryResult.category

	car.ManufacturerName = manufacturer.Name
	car.ManufacturerCountry = manufacturer.Country
	car.ManufacturerYear = manufacturer.Year
	car.CategoryName = category.Name

	// function returns the fetched car details along with any error encountered during the process

	return car, nil
}

// handles HTTP requests for car details, fetches the details, and renders them using an HTML template
func CarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// 	// extracts the id from the query parameters
	id := r.URL.Query().Get("id")

	// fetch car details based on the ID
	car, err := GetCarDetails(id) // // id is passed to the GetCarDetails()
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound) // 404
		return // //returned values are assigned to the variables car
	}

	// 	// parses the carDetails.html template
	tmpl, err := template.ParseFiles("static/carDetails.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
		return
	}

	// 	// executes the template with the fetched car data
	if err := tmpl.Execute(w, car); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
		return
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// finds a manufacturer by its ID from a slice of manufacturers
func FindManufacturerByID(id int, manufacturers []Manufacturer) (Manufacturer, error) {
	// 	// iterates through the slice of manufacturers
	for _, manufacturer := range manufacturers {
		if manufacturer.ID == id {
			return manufacturer, nil
		}
	} // // returns the manufacturer if the ID matches
	return Manufacturer{}, fmt.Errorf("manufacturer not found")
}

// finds a category by its ID from a slice of categories
func FindCategoryByID(id int, categories []Category) (Category, error) {
	// 	// iterates through the slice of categories
	for _, category := range categories {
		if category.ID == id {
			return category, nil
		}
	} // // returns the category if the ID matches
	return Category{}, fmt.Errorf("category not found")
}

// populates the ManufacturerName and CategoryName fields of a car
func PopulateCarDetails(car *Car, data CarsData) error {

	// uses FindManufacturerByID and FindCategoryByID to get the names
	manufacturer, err := FindManufacturerByID(car.ManufacturerID, data.Manufacturers)
	if err != nil {
		return err
	}
	// sets manufacturer info to variables
	car.ManufacturerName = manufacturer.Name
	car.ManufacturerCountry = manufacturer.Country
	car.ManufacturerYear = manufacturer.Year

	category, err := FindCategoryByID(car.CategoryID, data.Categories)
	if err != nil {
		return err
	}
	// // sets car.CategoryName to the name of the found category
	car.CategoryName = category.Name

	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// handles filter requests, filters car data based on query parameters, and renders the filtered data using an HTML template
func FilterPage(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(filePath) // // reads car data from a file
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file %s: %v", filePath, err), http.StatusInternalServerError) // 500
		return
	}

	var carData CarsData // // unmarshals the JSON data into a CarsData struct
	if err := json.Unmarshal(data, &carData); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse JSON: %v", err), http.StatusInternalServerError) // 500
		return
	}

	// extracts filter parameters from the query string
	manufacturer := r.URL.Query().Get("manufacturer")
	year := r.URL.Query().Get("year")
	category := r.URL.Query().Get("category")

	// 	filters car data using the filterCars()
	filteredCars := filterCars(carData, manufacturer, year, category)
	uniqueYears := getUniqueYears(carData.CarModels)

	// if no matches while filtering cars
	var message string
	if len(filteredCars) == 0 {
		message = "No cars found, try again!"
	}

	// 	// prepares the data for the template and...
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

	//  ...renders the filtered.html template with the filtered data
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
	// 	// iterates through the car models
	for _, car := range data.CarModels { // checks if each car matches the filter criteria
		if (manufacturer == "" || strings.EqualFold(data.Manufacturers[car.ManufacturerID-1].Name, manufacturer)) &&
			(year == "" || fmt.Sprintf("%d", car.Year) == year) &&
			(category == "" || strings.EqualFold(data.Categories[car.CategoryID-1].Name, category)) {
			PopulateCarDetails(&car, data)
			filtered = append(filtered, car)
		}
	} // // populates the car details and adds the matching cars to the filtered list
	return filtered
}

// extracts unique years from a list of cars.
func getUniqueYears(cars []Car) []int {
	// 	// map where the keys are the years from the cars slice
	yearMap := make(map[int]struct{})
	// 	// iterates over each car in the cars slice
	for _, car := range cars {
		yearMap[car.Year] = struct{}{}
	} // for each car, it adds the car.Year to yearMap. If the year is already in the map,
	// it won't add a duplicate because map keys are unique

	var years []int
	// 	// iterates over the keys in yearMap
	for year := range yearMap {
		years = append(years, year) // appends each unique year to the years slice
	}
	sort.Ints(years) // sorts the years slice in ascending order
	return years // // sorted years slice is returned
}
