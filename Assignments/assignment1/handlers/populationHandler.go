package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

/*
Handles the population endpoint
returns ApiResponsePopulation to client
*/
func CountryPopulationHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Error reading request. Missing argument or argument too short", http.StatusBadRequest)
		log.Println(w, "Error reading request.  Missing argument or argument too short", http.StatusBadRequest)
		return
	}
	if len(r.PathValue(argName)) > 2 {
		http.Error(w, "Error reading request. Invalid argument", http.StatusBadRequest)
		log.Println(w, "Error reading request.  Invalid argument", http.StatusBadRequest)
		return
	}
	// get country data
	country, err := getCountry(r.PathValue(argName), w)
	if err != nil {
		http.Error(w, "Failed to get country information", http.StatusInternalServerError)
		log.Println("Failed to get country information:", err)
		return
	}
	// get population API data
	var postRequest apiPostRequest
	postRequest.Country = country.Name

	jsonPost, err := json.Marshal(postRequest)
	if err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		log.Println("JSON Marshal Error:", err)
		return
	}

	jsonResponse, err := api_post(COUNTRIES_NOW_API+"/countries/population", jsonPost)
	if err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		log.Println("JSON Unmarshal Error:", err)
		return
	}

	// format country API data
	var apiResponse ApiResponsePopulation
	err = json.Unmarshal(jsonResponse, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		log.Println("JSON Unmarshal Error:", err)
		return
	}

	// filter ouput according to limit query
	if r.URL.Query().Get(queryName) != "" {
		queryParams := strings.Split(r.URL.Query().Get(queryName), "-")

		//split query parameters
		if len(queryParams) != 2 {
			http.Error(w, "Bad query, aborting request", http.StatusBadRequest)
			log.Println("Bad query, request aborted:")
			return
		}

		// convert query parameters to int
		queryParamsInt := []int{0, 0}
		for i := 0; i < len(queryParams); i++ {
			queryParamsInt[i], err = strconv.Atoi(queryParams[i])
			if err != nil {
				http.Error(w, "Bad request, query must be numbers seperated by a hyphen", http.StatusBadRequest)
				log.Println(w, "Bad request, query must be numbers seperated by a hyphen", http.StatusBadRequest)
				return
			}
		}

		// make sure query is in acending order
		sort.Ints(queryParamsInt)

		// filter output
		var tempPopulationArr []ApiPopulationData
		for i := 0; i < len(apiResponse.Data.PopulationData); i++ {
			if apiResponse.Data.PopulationData[i].Year >= queryParamsInt[0] && apiResponse.Data.PopulationData[i].Year <= queryParamsInt[1] {
				tempPopulationArr = append(tempPopulationArr, apiResponse.Data.PopulationData[i])
			}
		}
		apiResponse.Data.PopulationData = tempPopulationArr
	}

	// create response object
	returnJSON, err := json.Marshal(apiResponse)
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
