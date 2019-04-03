# Analyzing NYC Building footprints dataset

Back-end engineer assignment for topos inc.

Exract data, load into database with an api written using Go. Includes a simple frontend

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

* Install Go
* Install mongodb

MongoDB Driver

```
go get gopkg.in/mgo.v2
```

Gorilla/mux - URL router

```
go get github.com/gorilla/mux
```

### Running the code

Open two windows of terminal/cmd

Run mongodb server on one

```
mongod
```

Change directory to the project folder and build

```
go build
```

Run the app

```
go run app.go
```

## Testing

Open localhost:3000/index in a browser
