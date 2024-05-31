# Cars Viewer

This is task number four.


## How To Use

### Before starting the program:

1. Download 'api cars directory' inside 'cars' project from kood/sisu's page - if it is not in 'cars' repo already.
Read the README.md in 'api' directory:

Install NodeJS

Install NPM:
```
npm install
```

### Program starts node server and go server:

1. Type in terminal window:

```
go run main.go
```
2. Close server:

```
Ctrl+C
```

### Open browser:

1. Type in addres bar:

index.html
```
localhost:8080
```
raw car data
```
localhost:3000/api
localhost:3000/api/models
localhost:3000/api/manufacturers
localhost:3000/api/categories
```
car id
```
localhost:3000/api/models/{id}
localhost:3000/api/manufacturers/{id}
localhost:3000/api/categories/{id}
```

example id
```
localhost:3000/api/models/8
```

### Page functionality:

1. Index.html:

*  gallery
*  details button
*  filtering cars by year / category
*  
*  

2. CarDetails.html:

*  car detail window
*  
*  
*  
* 


## TO DO

Build a User-friendly Website:

* Use HTML and CSS to create a user-friendly website.
* Fetch and display information from the Cars API endpoints.
* Present the information using appropriate data visualization techniques.

Server Communication:

* Implement user-triggered actions that communicate with the Go backend.
* For example, clicking on a car model fetches additional details from the server.

Advanced Filtering:

* Implement advanced filtering and search options for users to easily find specific car models or manufacturers.

Comparisons:

* Allow users to compare different car models side-by-side, highlighting their key features and specifications.

Personalized Recommendations:

* Create personalized recommendations for users based on their preferences or past interactions with the website.


## Cars-viewer structure

```
cars/
│
├── api/
│   ├── img/
│   ├── node_modules/
│   ├── data.json
│   ├── main.js             // node server
│   ├── Makefile
│   ├── package-lock.json
│   ├── package.json
│   └── README.md
│
├── gofiles/
│   ├── CarData.go          // car data structs
│   ├── getCarData.go       // fetching car data
│   ├── homePage.go         // render page, parse data
│   └── serveImg.go  
│
├── static/
│   ├── carDetails.html     // car details page
│   ├── index.html          // main page
│   └── styles.css   
│
├── go.mod
├── main.go                 // starts node server and go server
└── README.md
             
```

### Coders

Laura Levistö - Jonathan Dahl - 5/24
