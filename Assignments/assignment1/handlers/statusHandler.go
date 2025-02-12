package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
Handles the status endpoint
returns StatusResponse to client
*/
func CountryStatusHandler(w http.ResponseWriter, r *http.Request) {
	var output string
	if DEBUG {
		output += "URL (Path): " + r.URL.Path + LINEBREAK
		output += "Method: " + r.Method + LINEBREAK
	}
	// Input sanitization
	if r.Method != http.MethodGet {
		http.Error(w, "Error reading method must be"+http.MethodGet, http.StatusBadRequest)
		log.Println(w, "Error reading method must be"+http.MethodGet, http.StatusBadRequest)
		return
	}

	var StatusResponse StatusResponse

	// make POST request to determine status
	var postRequest apiPostRequest
	postRequest.Country = "norway"

	jsonPost, err := json.Marshal(postRequest)
	if err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		log.Println("JSON Marshal Error:", err)
		return
	}
	reqBody := bytes.NewBuffer(jsonPost)
	countriesNowResponse, err := http.Post(COUNTRIES_NOW_API+"/countries/population", "application/json", reqBody)
	if err != nil {
		http.Error(w, "Error in response", http.StatusInternalServerError)
		log.Println("Error in response:", err)
		return
	}
	defer countriesNowResponse.Body.Close()

	// Make the HTTP GET request to determine status
	restCountriesResponse, err := http.Get(REST_COUNTIRES_API + "/no")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer restCountriesResponse.Body.Close()

	// populate response
	StatusResponse.Countriesnowapi = countriesNowResponse.StatusCode
	StatusResponse.Restcountriesapi = restCountriesResponse.StatusCode
	StatusResponse.Version = "v1"
	StatusResponse.Uptime = time.Now().Unix() - startTime.Unix()

	// create response object
	returnJSON, err := json.Marshal(StatusResponse)
	if err != nil {
		http.Error(w, "Error outputting JSON", http.StatusInternalServerError)
		log.Println(w, "Error outputting JSON", http.StatusInternalServerError)
		return
	}

	output += string(returnJSON) + LINEBREAK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "%v", output)
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
		log.Println(w, "Error when returning output", http.StatusInternalServerError)
		return
	}
}
