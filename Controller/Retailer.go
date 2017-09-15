package Controller

import (
	"encoding/json"
	"fmt"
	"grappos_api/Model"
	"io/ioutil"
	"log"
	"net/http"
)

var retailerBaseURL = "http://www.grappos.com/api2/retailer.php?1=1&format=json"

// RetailerDataRetriever Make calls to api for data.
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

// SearchCoordinates Searches a retailer based of their latitude and longitude.
func SearchCoordinates(lat string, lon string) (*Model.RetailersAPIResponse, error) {
	var s = new(Model.RetailersAPIResponse)
	queryString := fmt.Sprintf("&lat=%s&lon=%s", lat, lon)

	err := RetailerDataRetriever(s, retailerBaseURL+queryString)

	return s, err
}

// SearchProductByID Searches for a retailer that has a given ProductID.
func SearchProductByID(id int, st string) (*Model.RetailersAPIResponse, error) {
	var s = new(Model.RetailersAPIResponse)
	queryString := ""

	if CheckStoreType(st) {
		queryString = fmt.Sprintf("&product_id=%d&store_type=%s", id, st)
	} else {
		queryString = fmt.Sprintf("&product_id=%d&store_type=All", id)
	}

	err := RetailerDataRetriever(s, retailerBaseURL+queryString)

	return s, err
}

// SearchBrandByID Searches for a retailer that has product(s) with a given BrandID.
func SearchBrandByID(id int, st string) (*Model.RetailersAPIResponse, error) {
	var s = new(Model.RetailersAPIResponse)
	queryString := ""

	if CheckStoreType(st) {
		queryString = fmt.Sprintf("&brand_id=%d&store_type=%s", id, st)
	} else {
		queryString = fmt.Sprintf("&brand_id=%d&store_type=All", id)
	}

	err := RetailerDataRetriever(s, retailerBaseURL+queryString)

	return s, err
}

// CheckStoreType Checks the given input to determine which type of store
// the retailer is.
func CheckStoreType(st string) bool {
	if st == "All" || st == "Wine Shops" || st == "Restaurants" {
		return true
	}

	return false
}
