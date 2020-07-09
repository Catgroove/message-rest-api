package main

import (
	"log"
	"net/http"
	"backend/pkg/api"
)

func main() {
	API := api.CreateAPI()
	log.Fatal(http.ListenAndServe(":8080", API))
}
