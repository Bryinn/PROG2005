package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
)

/*
Country Info Endpoint: Return general country infos
The initial endpoint focuses returns general information for a given country, 2-letter country codes (ISO 3166-2).

Method: GET
Path: info/{:two_letter_country_code}{?limit=10}

eks: info/no
*/
func CountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	var output string
	argName := "p1"
	queryName := "limit"
	if DEBUG {
		output += "URL (Path): " + r.URL.Path + LINEBREAK
		output += "Path value: " + r.PathValue(argName) + LINEBREAK
		output += "HTTP argum: " + r.URL.Query().Get(queryName) + LINEBREAK
		output += "Method: " + r.Method + LINEBREAK
	}
	// Input santitization
	if r.Method != http.MethodGet {
		http.Error(w, "Error reading method must be"+http.MethodGet, http.StatusBadRequest)
		log.Println(w, "Error reading method must be"+http.MethodGet, http.StatusBadRequest)
		return
	}
	if len(r.PathValue(argName)) < 2 {
		http.Error(w, "Error reading request. Missing argument or argument too short. Two letter country code needed", http.StatusBadRequest)
		log.Println(w, "Error reading request.  Missing argument or argument too short. Two letter country code needed", http.StatusBadRequest)
		return
	}
	if len(r.PathValue(argName)) > 2 {
		http.Error(w, "Error reading request. Invalid argument, argument too log. Two letter country code needed", http.StatusBadRequest)
		log.Println(w, "Error reading request.  Invalid argument, argument too log. Two letter country code needed", http.StatusBadRequest)
		return
	}

	country, err := getCountry(r.PathValue(argName), w)
	if err != nil {
		http.Error(w, "Failed to get country information", http.StatusInternalServerError)
		log.Println("Failed to get country information:", err)
		return
	}

	// get cities api data
	var postRequest apiPostRequest
	postRequest.Country = country.Name

	jsonPost, err := json.Marshal(postRequest)
	if err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		log.Println("JSON Marshal Error:", err)
		return
	}

	jsonResponse, err := api_post(COUNTRIES_NOW_API+"/countries/cities", jsonPost)
	if err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		log.Println("JSON Unmarshal Error:", err)
		return
	}

	// format cities api data
	var postResponse apiResponseCity
	err = json.Unmarshal(jsonResponse, &postResponse)
	if err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		log.Println("JSON Unmarshal Error:", err)
		return
	}
	// Sorts cities in acending order
	sort.Strings(postResponse.Cities)

	str_limit := r.URL.Query().Get(queryName)
	if str_limit == "" {
		str_limit = "10"
	}
	if str_limit == "0" {
		country.Cities = postResponse.Cities
	} else {
		int_limit, err := strconv.Atoi(str_limit)
		if err != nil {
			http.Error(w, "Bad request, limit must be a number", http.StatusBadRequest)
			log.Println(w, "Bad request, limit must be a number", http.StatusBadRequest)
			return
		}

		for i := 0; i < int_limit; i++ {
			country.Cities = append(country.Cities, postResponse.Cities[i])
		}
	}

	// create response object
	returnJSON, err := json.Marshal(country)
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
