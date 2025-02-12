package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/**
* Api get request function
* Makes a get request to specified string, and returns byte array
 */
func api_get(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	res, _ := io.ReadAll(response.Body)
	// Read response body

	return res
}

/**
* Api post request function
* Makes a post request to specified string, and returns byte array
 */
func api_post(url string, body []byte) ([]byte, error) {
	reqBody := bytes.NewBuffer(body)
	response, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		err = fmt.Errorf("error posting request")
		return nil, err
	}

	defer response.Body.Close()
	r, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("error reading request response after post")
		return nil, err
	}
	return r, nil
}

/**
* Makes an API call to the country API
* Takes the target country two letter country code, and the writer object
* Returns the response of the API call, and an error if anything went wrong.
* Writer is only taken to make fitting errors
 */
func getCountry(two_letter_country_code string, w http.ResponseWriter) (countryAPIResponse, error) {

	//var countries countryAPIResponse
	var country countryAPIResponse
	// get country API data
	requestUrl := REST_COUNTIRES_API + "/" + two_letter_country_code
	res := api_get(requestUrl)

	// format country API data
	var apiResponses []apiResponseCountry
	err := json.Unmarshal(res, &apiResponses)
	if err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		log.Println("JSON Unmarshal Error:", err)
		return country, err
	}
	// Extract country API data
	country = countryAPIResponse{
		Name:       apiResponses[0].Name.Common,
		Continents: apiResponses[0].Continents,
		Population: apiResponses[0].Population,
		Languages:  apiResponses[0].Languages,
		Borders:    apiResponses[0].Borders,
		Flag:       apiResponses[0].Flags.Png,
		Capital:    "",
	}
	// extract the first capital if available
	if len(apiResponses[0].Capital) > 0 {
		country.Capital = apiResponses[0].Capital[0]
	}

	return country, nil
}
