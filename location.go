//Package grappos an API wrapper for the Grappos API
// (https://www.grappos.com/api-setup/)
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

// GetLocations Returns all locations.
func GetLocations(n int) (LocationAPIResponse, error) {
	var locationDataRetriever = NewDataRetriever("locate")
	var s = new(LocationAPIResponse)

	if n >= 0 {
		m := map[string]string{
			"limit": fmt.Sprintf("%d", n),
		}
		locationDataRetriever.addQueryParams(m)
	} else {
		return *s, errors.New("limit should be a positive int")
	}

	err := locationDataRetriever.getData(s)

	return *s, err
}

// SearchForLocation Postal Code or City Name (ex: “13066”, “San Francisco”).
func SearchForLocation(l string) (LocationAPIResponse, error) {
	var locationDataRetriever = NewDataRetriever("locate")
	var s = new(LocationAPIResponse)

	m := map[string]string{
		"locate": l,
	}

	locationDataRetriever.addQueryParams(m)

	err := locationDataRetriever.getData(s)

	if len(s.Locations) == 0 {
		return *s, errors.New("there aren't any wines at that location")
	}

	return *s, err
}
