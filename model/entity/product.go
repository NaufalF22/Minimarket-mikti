package entity

import "minimarket_mikti/model/domain"

type ProductEntity struct {
	ProductID int    `json:"product_id"`
	SellerID  int    `json:"seller_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
}

func ToProduct(product domain.Products) ProductEntity {
	return ProductEntity{
		ProductID: product.ProductID,
		SellerID:  product.SellerIDFK,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Price:     product.Price,
	}
}

func ToProductList(products []domain.Products) []ProductEntity {
	dataProduct := []ProductEntity{}

	for _, v := range products {
		dataProduct = append(dataProduct, ToProduct(v))
	}

	return dataProduct
}
