package grappos

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// DataRetriever holds the URL to fetch the data
type DataRetriever struct {
	URL url.URL
}

// NewDataRetriever creates a new DataRetriever to get all of the data
func NewDataRetriever(m string) DataRetriever {
	return DataRetriever{
		URL: url.URL{
			Scheme:   "http",
			Host:     "www.grappos.com",
			Path:     "api2/" + m + ".php",
			RawQuery: "format=json",
		},
	}
}

func (d *DataRetriever) addQueryParams(m map[string]string) {
	d.URL.RawQuery = ""
	p := d.URL.Query()
	p.Add("format", "json")
	for i, x := range m {
		p.Add(i, x)
	}
	d.URL.RawQuery = p.Encode()
}

func (d *DataRetriever) getData(v interface{}) error {

	res, err := http.Get(d.URL.String())
	if err != nil {
		err = errors.New("error fetching data")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = errors.New("error reading data")
	}

	err = json.Unmarshal(body, v)

	return err
}
