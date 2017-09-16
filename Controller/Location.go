package Controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var baseURL = "http://www.grappos.com/api2/locate.php?1=1&format=json"

// Location The sruct that holds the location data
type Location struct {
	DisplayName string `json:"displayName"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lon"`
	ZipCode     string `json:"zip"`
}

// LocationsAPIResponse Holds the API response
type LocationsAPIResponse struct {
	Locations []Location `json:"locations"`
}

// LocationDataRetriever Make calls to api for data.
func LocationDataRetriever(m *LocationsAPIResponse, q string) error {
	res, err := http.Get(q)

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal([]byte(body), &m)

	return err
}

// GetLocations Returns all locations.
func GetLocations(n int) (*LocationsAPIResponse, error) {

	var s = new(LocationsAPIResponse)
	queryParams := ""

	if n >= 0 {
		queryParams = fmt.Sprintf("&limit=%d", n)
	} else {
		return s, errors.New("Limit should be a positive int")
	}

	err := LocationDataRetriever(s, baseURL+queryParams)
	if err != nil {
		log.Fatalf("Searching for location went wrong: %s", err)
	}

	return s, err
}

// SearchForLocation Postal Code or City Name (ex: “13066”, “San Francisco”).
func SearchForLocation(l string) (*LocationsAPIResponse, error) {
	queryParams := ""

	var s = new(LocationsAPIResponse)

	if len(l) == 5 {
		queryParams = fmt.Sprintf("&locate=%s", l)
	} else {
		return s, errors.New("Invalid Postal Code")
	}

	err := LocationDataRetriever(s, baseURL+queryParams)
	if err != nil {
		log.Fatalf("Searching for location went wrong: %s", err)
	}

	return s, err
}
