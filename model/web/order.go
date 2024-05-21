package web

type OrderRequest struct {
	Quantity int `validate:"required" json:"quantity"`
	Price    int `validate:"required" json:"price"`
}
