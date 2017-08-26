package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	DisplayName string `json:"displayName"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lon"`
	ZipCode     string `json:"zip"`
}

type LocationsAPIResponse struct {
	Locations []Location `json:"locations"`
}

func getLocations(body []byte) (*LocationsAPIResponse, error) {
	var s = new(LocationsAPIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func main() {
	res, err := http.Get("http://www.grappos.com/api2/locate.php?format=json")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	s, err := getLocations([]byte(body))

	fmt.Println(s)
}
