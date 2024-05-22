package web

type StoreProductRequest struct {
	SellerIDFK int    `validate:"required" json:"seller_id_fk"`
	Name       string `validate:"required" json:"name"`
	Quantity   int    `validate:"required" json:"quantity"`
	Price      int    `validate:"required" json:"price"`
}
