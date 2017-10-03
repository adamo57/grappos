package grappos

import (
	"errors"
	"fmt"
)

type location struct {
	DisplayName string `json:"displayName"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lon"`
	ZipCode     string `json:"zip"`
}

// LocationAPIResponse Holds the API response.
type LocationAPIResponse struct {
	Locations []location `json:"locations"`
}

var locationDataRetriever = NewDataRetriever("locate")

// GetLocations Returns all locations.
func GetLocations(n int) (LocationAPIResponse, error) {

	var s = new(LocationAPIResponse)

	if n >= 0 {
		m := map[string]string{
			"limit": fmt.Sprintf("%d", n),
		}
		locationDataRetriever.addQueryParams(m)
	} else {
		return *s, errors.New("Limit should be a positive int")
	}

	err := locationDataRetriever.getData(s)

	return *s, err
}

// SearchForLocation Postal Code or City Name (ex: “13066”, “San Francisco”).
func SearchForLocation(l string) (LocationAPIResponse, error) {

	var s = new(LocationAPIResponse)

	if len(l) == 5 {
		m := map[string]string{
			"locate": l,
		}
		locationDataRetriever.addQueryParams(m)
	} else {
		return *s, errors.New("invalid Postal Code")
	}

	err := locationDataRetriever.getData(s)

	return *s, err
}
