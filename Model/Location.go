package Model

type Location struct {
	DisplayName string `json:"displayName"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lon"`
	ZipCode     string `json:"zip"`
}

type LocationsAPIResponse struct {
	Locations []Location `json:"locations"`
}
