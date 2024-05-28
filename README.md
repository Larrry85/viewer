# Cars Viewer

This is task number four.


## How To Use

### Before starting the program:

1. Download 'api cars directory' inside 'cars' project from kood/sisu's page - if it is not in 'cars' repo already.
Read the README.md in 'api' directory:

Install NodeJS

Install NPM: npm install

### Program starts node server and go server:

1. Type in terminal window:

```
cd api
node main.js
```

2. Type in another terminal window:

```
cd ..
go run main.go
```


### Open browser:

3. Type in addres bar:

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

1. index.html:

*  gallery
*  detail button
*  
*  
*  

2. carDetails.html:

*  car detail window
*  
*  
*  
* 


### Close the servers:

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
│   ├── getCarData.go         // starts go server 8080
│   └── handlers.go  
│
├── static/
│   ├── carDetails.html
│   ├── index.html  
│   └── styles.css   
│
├── go.mod
├── main.go    
└── README.md
             
```

### Coders

Laura Levistö - Jonathan Dahl - 5/24
