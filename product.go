package grappos

import (
	"errors"
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

// GetProducts Returns a list of Products.
func GetProducts(u string) (ProductAPIResponse, error) {
	var productDataRetriever = NewDataRetriever("subscriber")
	var s = new(ProductAPIResponse)

	m := map[string]string{
		"uid": u,
	}
	productDataRetriever.addQueryParams(m)

	err := productDataRetriever.getData(s)
	if err != nil {
		err = errors.New("searching for location went wrong")
	}

	if len(s.Products) == 0 {
		err = errors.New("product not found with id")
	}

	return *s, err
}
