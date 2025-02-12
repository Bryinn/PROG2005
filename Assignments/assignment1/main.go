package main

import (
	"assignment1/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	// Override port with default port if not provided (e.g. local deployment)
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	router := http.NewServeMux()

	router.HandleFunc("/countryinfo/v1/info/{p1}", handlers.CountryInfoHandler)
	router.HandleFunc("/countryinfo/v1/population/{p1}", handlers.CountryPopulationHandler)
	router.HandleFunc("/countryinfo/v1/status/", handlers.CountryStatusHandler)

	log.Println("Starting server on port " + port + " ...")
	log.Println(http.ListenAndServe(":"+port, router))

}
