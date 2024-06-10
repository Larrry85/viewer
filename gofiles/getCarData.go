package gofiles

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
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

// fetches detailed information about a car from an API using channels for concurrency
func GetCarDetails(idStr string) (Car, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Car{}, errors.New("invalid car ID")
	}

	carCh := make(chan carResult)
	manufacturerCh := make(chan manufacturerResult)
	categoryCh := make(chan categoryResult)

	go func() {
		// Fetch car details
		resp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/models/%d", id))
		if err != nil {
			carCh <- carResult{Car{}, err}
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

	// Receive car details
	carResult := <-carCh
	if carResult.err != nil {
		return Car{}, carResult.err
	}
	car = carResult.car

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

	return car, nil
}

// handles HTTP requests for car details, fetches the details, and renders them using an HTML template
func CarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	car, err := GetCarDetails(id)
	if err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("static/carDetails.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, car); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// finds a manufacturer by its ID from a slice of manufacturers
func FindManufacturerByID(id int, manufacturers []Manufacturer) (Manufacturer, error) {
	for _, manufacturer := range manufacturers {
		if manufacturer.ID == id {
			return manufacturer, nil
		}
	}
	return Manufacturer{}, fmt.Errorf("manufacturer not found")
}

// finds a category by its ID from a slice of categories
func FindCategoryByID(id int, categories []Category) (Category, error) {
	for _, category := range categories {
		if category.ID == id {
			return category, nil
		}
	}
	return Category{}, fmt.Errorf("category not found")
}

// populates the ManufacturerName and CategoryName fields of a car
func PopulateCarDetails(car *Car, data CarsData) error {
	manufacturer, err := FindManufacturerByID(car.ManufacturerID, data.Manufacturers)
	if err != nil {
		return err
	}
	car.ManufacturerName = manufacturer.Name
	car.ManufacturerCountry = manufacturer.Country
	car.ManufacturerYear = manufacturer.Year

	category, err := FindCategoryByID(car.CategoryID, data.Categories)
	if err != nil {
		return err
	}
	car.CategoryName = category.Name

	return nil
}

// handles filter requests, filters car data based on query parameters, and renders the filtered data using an HTML template
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

	manufacturer := r.URL.Query().Get("manufacturer")
	year := r.URL.Query().Get("year")
	category := r.URL.Query().Get("category")

	filteredCars := filterCars(carData, manufacturer, year, category)
	uniqueYears := getUniqueYears(carData.CarModels)

	var message string
	if len(filteredCars) == 0 {
		message = "No cars found, try again!"
	}

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

	tmpl, err := template.ParseFiles("static/filtered.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, filteredData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// filters car models based on manufacturer, year, and category
func filterCars(data CarsData, manufacturer, year, category string) []Car {
	var filtered []Car
	for _, car := range data.CarModels {
		if (manufacturer == "" || strings.EqualFold(data.Manufacturers[car.ManufacturerID-1].Name, manufacturer)) &&
			(year == "" || fmt.Sprintf("%d", car.Year) == year) &&
			(category == "" || strings.EqualFold(data.Categories[car.CategoryID-1].Name, category)) {
			PopulateCarDetails(&car, data)
			filtered = append(filtered, car)
		}
	}
	return filtered
}

// extracts unique years from a list of cars.
func getUniqueYears(cars []Car) []int {
	yearMap := make(map[int]struct{})
	for _, car := range cars {
		yearMap[car.Year] = struct{}{}
	}

	var years []int
	for year := range yearMap {
		years = append(years, year)
	}
	sort.Ints(years)
	return years
}
