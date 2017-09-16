package Controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var productBaseURL = "http://www.grappos.com/api2/subscriber.php?1=1&format=json"

// Product The sruct that holds the product data
type Product struct {
	ProductID   string `json:"ProductID"`
	Description string `json:"description"`
	BrandID     string `json:"BrandID"`
	IsToppick   int    `json:"is_toppick"`
}

// ProductsAPIResponse Holds the API response
type ProductsAPIResponse struct {
	Products []Product `json:"products"`
}

// ProductDataRetriever Make calls to api for data.
func ProductDataRetriever(m *ProductsAPIResponse, q string) error {
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

// GetProducts Returns a list of Products.
func GetProducts(u string) (*ProductsAPIResponse, error) {
	var s = new(ProductsAPIResponse)
	queryParams := fmt.Sprintf("&uid=%s", u)

	err := ProductDataRetriever(s, productBaseURL+queryParams)

	if len(s.Products) == 0 {
		err = fmt.Errorf("Product not found with id: %s", u)
	}

	return s, err
}
