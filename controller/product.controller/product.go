package productcontroller

import (
	"minimarket_mikti/helper"
	"minimarket_mikti/model/web"
	orderservice "minimarket_mikti/service/order.service"
	productservice "minimarket_mikti/service/product.service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service      productservice.ProductServiceInterface
	orderService orderservice.OrderServiceInterface
}

func NewProductService(service productservice.ProductServiceInterface, orderService orderservice.OrderServiceInterface) *ProductController {
	return &ProductController{
		service:      service,
		orderService: orderService,
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

func (pC *ProductController) BuyProduct(c echo.Context) error {
	newOrder := new(web.OrderRequest)
	id, _ := strconv.Atoi(c.Param("id"))
	userToken := c.Get("user").(*jwt.Token)
	claims, _ := userToken.Claims.(*helper.CustomClaims)

	if err := c.Bind(newOrder); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newOrder); err != nil {
		return err
	}

	createOrder, errOrder := pC.orderService.Create(claims.UserID, id, *newOrder)
	_, errSave := pC.service.BuyProduct(id, *newOrder)

	if errSave != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errSave.Error(), nil))
	}

	if errOrder != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errOrder.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success order", createOrder))
}
