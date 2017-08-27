package Controller

//NOTE: this doesnt actually work on the API, but im going to write it anyways ¯\_(ツ)_/¯

import (
	"encoding/json"
	"fmt"
	"hvíla/Model"
	"io/ioutil"
	"log"
	"net/http"
)

var productBaseURL = "http://www.grappos.com/api2/subscriber.php?1=1&format=json"

//DataRetriever
//Make calls to api for data
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

func GetProducts(u int) (*Model.ProductsAPIResponse, error) {
	var s = new(Model.ProductsAPIResponse)
	queryParams := fmt.Sprintf("?uid=%d", u)

	err := ProductDataRetriever(s, productBaseURL+queryParams)

	return s, err
}
