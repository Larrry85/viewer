// OLD VERSION

/*
// goServer.go
package gofiles

import (
	"html/template"
	"log"
	"net/http"
)

// Open your browser and navigate to http://localhost:8080/
// You should see the car data displayed in the HTML template
func GoServer() error {
	http.HandleFunc("/", homePage)
	// !!!!!!!!!!!!!!!!
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/models", modelsHandler)
	http.HandleFunc("/api/categories", categoriesHandler)
	http.HandleFunc("/api/manufacturers", manufacturersHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/api/images/", http.StripPrefix("/api/images", http.FileServer(http.Dir("img"))))

	log.Println("starting server at port 8080")
	return http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound) // 404
		return
	}

	// Fetch car data from the API server
	carsData, err := GetCarData()
	if err != nil {
		http.Error(w, "Failed to fetch car data", http.StatusInternalServerError) // 500
		return
	}

// !!!!!!!!!!!!!!!!
	// Prepare a map for quick lookup of manufacturers by ID
	manufacturerMap := make(map[int]string)
	for _, manufacturer := range carsData.Manufacturers {
		manufacturerMap[manufacturer.ID] = manufacturer.Name
	}

// !!!!!!!!!!!!!!!!
	// Add Manufacturer names to CarModels based on ManufacturerID
	for i, car := range carsData.CarModels {
		carsData.CarModels[i].ManufacturerName = manufacturerMap[car.ManufacturerID]
	}


	// Render the HTML template with the car data
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}
	err = tmpl.Execute(w, carsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}
}
*/