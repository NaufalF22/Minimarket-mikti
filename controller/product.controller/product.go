package productcontroller

import (
	"minimarket_mikti/helper"
	productservice "minimarket_mikti/service/product.service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service productservice.ProductServiceInterface
}

func NewProductService(service productservice.ProductServiceInterface) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (pC *ProductController) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, errProduct := pC.service.GetProductID(id)

	if errProduct != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errProduct.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Product", product))
}

func (pC *ProductController) GetProducts(c echo.Context) error {
	products, errProducts := pC.service.GetAllProduct()

	if errProducts != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errProducts.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Product", products))
}
