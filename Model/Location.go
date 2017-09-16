package Model

// Location The sruct that holds the location data
type Location struct {
	DisplayName string `json:"displayName"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lon"`
	ZipCode     string `json:"zip"`
}

// LocationsAPIResponse Holds the API response
type LocationsAPIResponse struct {
	Locations []Location `json:"locations"`
}
