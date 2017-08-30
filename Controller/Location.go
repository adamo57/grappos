package Controller

import (
	"encoding/json"
	"fmt"
	"hvíla/Model"
	"io/ioutil"
	"log"
	"net/http"
)

var baseURL = "http://www.grappos.com/api2/locate.php?1=1&format=json"

//LocationDataRetriever
//Make calls to api for data
func LocationDataRetriever(m *Model.LocationsAPIResponse, q string) error {
	res, err := http.Get(q)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		log.Println("whoops:", err)
	}

	return err
}

// GetLocations
// returns all locations
func GetLocations(n int) (*Model.LocationsAPIResponse, error) {

	var s = new(Model.LocationsAPIResponse)
	queryParams := ""

	if n > 0 {
		queryParams = fmt.Sprintf("&limit=%d", n)
	}

	err := LocationDataRetriever(s, baseURL+queryParams)

	return s, err
}

// SearchForLocation
// Postal Code or City Name (ex: “13066”, “San Francisco”)
func SearchForLocation(l string) (*Model.LocationsAPIResponse, error) {
	queryParams := fmt.Sprintf("&locate=%s", l)

	var s = new(Model.LocationsAPIResponse)

	err := LocationDataRetriever(s, baseURL+queryParams)

	return s, err
}
