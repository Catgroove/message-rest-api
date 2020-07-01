# REST API for a simple message service 

# Frontend + Backend Stack

## Requirements

* [Docker](https://www.docker.com/)

## Installation

Clone the repo:

`git clone https://github.com/Catgroove/message-rest-api.git`

Run docker-compose:

`docker-compose up`

The site is available at [http://localhost:3000/](http://localhost:3000/).


# Backend Only


## Installation

Clone the repo:

`git clone https://github.com/Catgroove/message-rest-api.git`

cd into backend:

`cd backend`

Install all the dependencies:

`go get -u ./...`

Run the server:

`go run main.go`

The server is available at [http://localhost:8080/](http://localhost:8080/).


## Tests

Run tests:

`go test ./...`