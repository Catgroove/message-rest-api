package main

import (
	"backend/pkg/api"
	"log"
	"net/http"
)

func main() {
	API := api.CreateAPI()
	log.Fatal(http.ListenAndServe(":8080", API))
}
