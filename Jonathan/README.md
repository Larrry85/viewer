# Cars Viewer

This is task number four.


## How To Use

1. Program starts api server and go server:

```
go build main.go
./main
```

2.  Type in address bar:

```
localhost:8080
localhost:3000/api
```

3. Index.html page functionality:

*  
*  
*  
*  
*  

4. Close the server:

```
Ctrl+C
```

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
│   ├── node_modules
│   ├── data.json
│   ├── main.js
│   ├── Makefile
│   ├── package-lock.json
│   ├── package.json
│   └── README.md
│
├── gofiles/
│   ├── apiServer.go          // starts api server 3000
│   ├── CarData.go            // car data struct
│   ├── getCarData.go         // how to get data to html
│   └── goServer.go           // starts go server 8080
│
├── static/
│   ├── index.html  
│   └── styles.css   
│
├── go.mog
├── main.go    
└── README.md
             
```

### Coders

Laura Levistö - Jonathan Dahl - 5/24