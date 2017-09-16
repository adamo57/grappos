package Controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adamo57/grappos_api/Model"
)

var productBaseURL = "http://www.grappos.com/api2/subscriber.php?1=1&format=json"

// ProductDataRetriever Make calls to api for data.
func ProductDataRetriever(m *Model.ProductsAPIResponse, q string) error {
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
func GetProducts(u string) (*Model.ProductsAPIResponse, error) {
	var s = new(Model.ProductsAPIResponse)
	queryParams := fmt.Sprintf("&uid=%s", u)

	err := ProductDataRetriever(s, productBaseURL+queryParams)

	if len(s.Products) == 0 {
		err = fmt.Errorf("Product not found with id: %s", u)
	}

	return s, err
}
