package grappos

import (
	"fmt"
	"log"
)

type product struct {
	ProductID   string `json:"ProductID"`
	Description string `json:"description"`
	BrandID     string `json:"BrandID"`
	IsToppick   int    `json:"is_toppick"`
}

// ProductAPIResponse Holds the API response.
type ProductAPIResponse struct {
	Products []product `json:"products"`
}

var productDataRetriever = NewDataRetriever("subscriber")

// GetProducts Returns a list of Products.
func GetProducts(u string) (ProductAPIResponse, error) {

	var s = new(ProductAPIResponse)

	m := map[string]string{
		"uid": u,
	}
	productDataRetriever.addQueryParams(m)

	err := productDataRetriever.getData(s)
	if err != nil {
		log.Fatalf("Searching for location went wrong: %s", err)
	}

	if len(s.Products) == 0 {
		err = fmt.Errorf("product not found with id: %s", u)
	}

	return *s, err
}
