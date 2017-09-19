package grappos

import (
	"fmt"
)

type retailer struct {
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

// RetailerAPIResponse Holds the API response.
type RetailerAPIResponse struct {
	Retailers []retailer `json:"retailers"`
}

var retailerDataRetriever = NewDataRetriever("retailer")

// SearchCoordinates Searches a retailer based of their latitude and longitude.
func SearchCoordinates(lat string, lon string) (RetailerAPIResponse, error) {
	var s = new(RetailerAPIResponse)

	m := map[string]string{
		"lat": lat,
		"lon": lon,
	}
	retailerDataRetriever.addQueryParams(m)

	err := retailerDataRetriever.getData(s)
	if err != nil {
		err = fmt.Errorf("something went wrong: %s", err)
	}

	return *s, err
}

// SearchProductByID Searches for a retailer that has a given ProductID.
func SearchProductByID(id int, st string) (RetailerAPIResponse, error) {
	var s = new(RetailerAPIResponse)
	m := map[string]string{}

	if checkStoreType(st) {
		m["product_id"] = fmt.Sprintf("%d", id)
		m["store_type"] = st
	} else {
		m["product_id"] = fmt.Sprintf("%d", id)
		m["store_type"] = "All"
	}

	retailerDataRetriever.addQueryParams(m)

	err := retailerDataRetriever.getData(s)

	return *s, err
}

// SearchBrandByID Searches for a retailer that has product(s) with a given BrandID.
func SearchBrandByID(id int, st string) (RetailerAPIResponse, error) {
	var s = new(RetailerAPIResponse)
	m := map[string]string{}

	if checkStoreType(st) {
		m["brand_id"] = fmt.Sprintf("%d", id)
		m["store_type"] = st
	} else {
		m["brand_id"] = fmt.Sprintf("%d", id)
		m["store_type"] = "All"
	}

	retailerDataRetriever.addQueryParams(m)

	err := retailerDataRetriever.getData(s)

	return *s, err
}

// CheckStoreType Checks the given input to determine which type of store
// the retailer is.
func checkStoreType(st string) bool {
	if st == "All" || st == "Wine Shops" || st == "Restaurants" {
		return true
	}

	return false
}
