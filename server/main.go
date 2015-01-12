package main

import (
	"github.com/liamgiles/api"
	"log"
	"net/http"
)

func main() {

	api, err := api.New()
	if err != nil {

		log.Fatalf("Failed to create API: %s", err.Error)
	}
	err = http.ListenAndServe(":8081", api)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
