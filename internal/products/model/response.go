package model

type ProductResponse struct {
	ID string `json:"id"`
}

type UpdateProductResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
