package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type PageData struct {
	Title string
}

func main() {
	// Define handler functions for different routes
	http.HandleFunc("/", homePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("starting server at port 8080")
	//starts server on port 3000 and uses nil as handler = http.DefaultServeMux
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { //checks if requested url path is right
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	var HomeData = PageData{
		Title: "Cars Api",
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = tmpl.Execute(w, HomeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
