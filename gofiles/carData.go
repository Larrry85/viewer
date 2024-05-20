// carData.go
package gofiles

// Manufacturer represents a car manufacturer
type Manufacturer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Year    int    `json:"foundingYear"`
}

// Category represents a car category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Car represents a car model
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
	Image string `json:"image"`
}

// CarsData: Represents the overall data structure retrieved from the API.
//
//	It contains slices of Car, Manufacturer, and Category
type CarsData struct {
	CarModels     []Car          `json:"carModels"`
	Manufacturers []Manufacturer `json:"manufacturers"`
	Categories    []Category     `json:"categories"`
}
