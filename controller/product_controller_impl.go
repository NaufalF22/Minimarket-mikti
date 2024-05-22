package controller

import (
	"minimarket_mikti/model"
	"minimarket_mikti/model/web"
	"minimarket_mikti/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductControllerImpl struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		service: service,
	}
}

func (controller *ProductControllerImpl) CreateProduct(c echo.Context) error {
	product := new(web.StoreProductRequest)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(product); err != nil {
		return err
	}

	saveProduct, errSaveProduct := controller.service.CreateProduct(*product)

	if errSaveProduct != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveProduct.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Product created successfully", saveProduct))
}
