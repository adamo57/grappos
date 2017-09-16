package Controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var retailerBaseURL = "http://www.grappos.com/api2/retailer.php?1=1&format=json"

// Retailer The sruct that holds the retailer data
type Retailer struct {
	RetailerID  string  `json:"RetailerID"`
	Name        string  `json:"name"`
	Address1    string  `json:"addr1"`
	Address2    string  `json:"addr2"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	StateAbbr   string  `json:"state_abbr"`
	ZipCode     string  `json:"zip"`
	Phone       string  `json:"phone"`
	Latitude    string  `json:"lat"`
	Longitude   string  `json:"lon"`
	IsRetail    string  `json:"is_retail"`
	IsOnPremise string  `json:"is_on_premise"`
	Distance    float32 `json:"distance"`
}

// RetailersAPIResponse Holds the API response
type RetailersAPIResponse struct {
	Retailers []Retailer `json:"retailers"`
}

// RetailerDataRetriever Make calls to api for data.
func RetailerDataRetriever(m *RetailersAPIResponse, q string) error {
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
func SearchCoordinates(lat string, lon string) (*RetailersAPIResponse, error) {
	var s = new(RetailersAPIResponse)
	queryString := fmt.Sprintf("&lat=%s&lon=%s", lat, lon)

	err := RetailerDataRetriever(s, retailerBaseURL+queryString)

	return s, err
}

// SearchProductByID Searches for a retailer that has a given ProductID.
func SearchProductByID(id int, st string) (*RetailersAPIResponse, error) {
	var s = new(RetailersAPIResponse)
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
func SearchBrandByID(id int, st string) (*RetailersAPIResponse, error) {
	var s = new(RetailersAPIResponse)
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
