package grappos_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var productBaseURL = "http://www.grappos.com/api2/subscriber.php?1=1&format=json"

type product struct {
	ProductID   string `json:"ProductID"`
	Description string `json:"description"`
	BrandID     string `json:"BrandID"`
	IsToppick   int    `json:"is_toppick"`
}

type productsAPIResponse struct {
	Products []product `json:"products"`
}

func productDataRetriever(m *productsAPIResponse, q string) error {
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
func GetProducts(u string) (*productsAPIResponse, error) {
	var s = new(productsAPIResponse)
	queryParams := fmt.Sprintf("&uid=%s", u)

	err := productDataRetriever(s, productBaseURL+queryParams)

	if len(s.Products) == 0 {
		err = fmt.Errorf("Product not found with id: %s", u)
	}

	return s, err
}
