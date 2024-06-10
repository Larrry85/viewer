// carData.go
package gofiles

// This file defines the core data structures used in the application for representing cars,
// manufacturers, and categories. These structs are equipped with JSON tags to facilitate
// easy serialization and deserialization of JSON data. The CarsData struct aggregates these
// entities and includes endpoint information, making it a central point for managing
// car-related data in the application.

// struct that represents a car manufacturer.
type Manufacturer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Year    int    `json:"foundingYear"`
}

// struct that represents a category of cars
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// struct that represents a car model
type Car struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ManufacturerID int    `json:"manufacturerId"`
	CategoryID     int    `json:"categoryId"`
	Year           int    `json:"year"`
	Specifications struct {
		Engine       string `json:"engine"`
		Horsepower   int    `json:"horsepower"`
		Transmission string `json:"transmission"`
		Drivetrain   string `json:"drivetrain"`
	} `json:"specifications"`
	Image               string `json:"image"`
	ManufacturerName    string `json:"manufacturerName"`
	ManufacturerCountry string `json:"manufacturerCountry"`
	ManufacturerYear    int    `json:"manufacturerYear"`
	CategoryName        string `json:"categoryName"`
}

// struct that contains slices of Car, Manufacturer, and Category.
type CarsData struct {
	CarModels             []Car          `json:"carModels"`
	Manufacturers         []Manufacturer `json:"manufacturers"`
	Categories            []Category     `json:"categories"`
	ModelEndpoint         string         `json:"modelsEndpoint"`        //
	CategoriesEndpoint    string         `json:"categoriesEndpoint"`    // endpoints
	ManufacturersEndpoint string         `json:"manufacturersEndpoint"` //
}
