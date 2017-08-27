package Model

type Product struct {
	ProductID   string `json:"ProductID"`
	Description string `json:"description"`
	BrandID     string `json:"BrandID"`
	IsToppick   int    `json:"is_toppick"`
}

type ProductsAPIResponse struct {
	Products []Product `json:"products"`
}
