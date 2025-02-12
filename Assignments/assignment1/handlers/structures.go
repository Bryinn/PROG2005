package handlers

//status object return data
type StatusResponse struct {
	Countriesnowapi  int
	Restcountriesapi int
	Version          string
	Uptime           int64
}

//Population object return data
type ApiPopulationData struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type ApiPopCountryData struct {
	Country        string              `json:"country"`
	Code           string              `json:"code"`
	Iso3           string              `json:"iso3"`
	PopulationData []ApiPopulationData `json:"populationCounts"`
}

type ApiResponsePopulation struct {
	Error   bool              `json:"error"`
	Message string            `json:"msg"`
	Data    ApiPopCountryData `json:"data"`
}

//data from country api
type countryAPIResponse struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

type apiResponseCountry struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		Png string `json:"png"`
	} `json:"flags"`
	Capital []string `json:"capital"`
}
type apiResponseCity struct {
	Err     bool     `json:"error"`
	Message string   `json:"msg"`
	Cities  []string `json:"data"`
}
type apiPostRequest struct {
	Country string `json:"country"`
}
