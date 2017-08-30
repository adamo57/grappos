package Controller

import (
	"encoding/json"
	"fmt"
	"hvíla/Model"
	"io/ioutil"
	"log"
	"net/http"
)

var retailerBaseURL = "http://www.grappos.com/api2/retailer.php?1=1&format=json"

//DataRetriever
//Make calls to api for data
func RetailerDataRetriever(m *Model.RetailersAPIResponse, q string) error {
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

func SearchCoordinates(lat float32, lon float32) (*Model.RetailersAPIResponse, error) {
	var s = new(Model.RetailersAPIResponse)
	queryString := fmt.Sprintf("?lat=%d?lon=%d", lat, lon)

	err := RetailerDataRetriever(s, retailerBaseURL+queryString)

	return s, err
}

func SearchProductByID(id int, st string) (*Model.RetailersAPIResponse, error) {
	var s = new(Model.RetailersAPIResponse)
	queryString := ""

	if CheckStoreType(st) {
		queryString = fmt.Sprintf("?product_id=%d?store_type=%s", id, st)
	} else {
		queryString = fmt.Sprintf("?product_id=%d?store_type=All", id)
	}

	err := RetailerDataRetriever(s, retailerBaseURL+queryString)

	return s, err
}

func SearchBrandByID(id int, st string) (*Model.RetailersAPIResponse, error) {
	var s = new(Model.RetailersAPIResponse)
	queryString := ""

	if CheckStoreType(st) {
		queryString = fmt.Sprintf("?brand_id=%d?store_type=%s", id, st)
	} else {
		queryString = fmt.Sprintf("?brand_id=%d?store_type=All", id)
	}

	err := RetailerDataRetriever(s, retailerBaseURL+queryString)

	return s, err
}

func CheckStoreType(st string) bool {
	if st == "All" || st == "Wine Shops" || st == "Restaurants" {
		return true
	}

	return false
}
