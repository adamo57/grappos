package Model

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
