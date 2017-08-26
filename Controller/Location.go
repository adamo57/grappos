package Controller

import (
	"encoding/json"
	"fmt"
	"hvíla/Model"
	"io/ioutil"
	"net/http"
)

var baseURL = "http://www.grappos.com/api2/locate.php?format=json"

func GetLocations() (*Model.LocationsAPIResponse, error) {

	var s = new(Model.LocationsAPIResponse)

	res, err := http.Get(baseURL)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal([]byte(body), &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	return s, err
}
