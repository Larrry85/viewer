# Cars Viewer

This is task number four.


## How To Use

### Before starting the program:

1. Download 'api cars directory' inside 'cars' project from kood/sisu's page - if it is not in 'cars' repository already.
Read the README.md in 'api' directory:

Install NodeJS

Install NPM:
```
npm install
```

### Program starts node server and go server:

1. Type in terminal window:

```
go build main.go
./main
```
2. Close server at the end:

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

*  car gallery of all the cars
*  details button for enxtra info
*  user can filter cars by:
    - manufacturers
    - year
    - category
*  
*  

2. CarDetails.html:

*  car detail window showing all data of one car
*  styled to fit different sized windowd
*  
*  
* 

3. filtered.html:

*  this page shows cars filtered by user
*  if no matches, it shows a message
*  
*  
* 


## What we needed to to in this task

User-friendly Website:

* website that effectively presents the car data retrieved from the API.
* data visualization techniques to make the information easy to understand and navigate.
* communication with the server, triggered by user action, (fetching specific data or initiating a certain process on the server.)

Go backend:

* a stable and well-structured backend for the website.
* requests and errors handled gracefully to ensure a smooth user experience.
* best practices and coding conventions for readability and maintainability.

Extra requirements

Advanced Filtering:

* advanced filtering and search options for users to easily find specific car models or manufacturers.

Comparisons:

* users are allowed to compare different car models side-by-side, highlighting their key features and specifications.

Personalized Recommendations:

* personalized recommendations for users based on their preferences or past interactions with the website.


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
│   ├── filtered.html       // filtered cars page
│   ├── index.html          // main page
│   └── styles.css   
│
├── go.mod
├── main
├── main.go                 // starts node server and go server
└── README.md
             
```

### Coders

Laura Levistö - Jonathan Dahl - 6/24
