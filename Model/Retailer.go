package Model

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
