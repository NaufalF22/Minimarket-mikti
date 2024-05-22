package controller

import "github.com/labstack/echo/v4"

type ProductController interface {
	CreateProduct(c echo.Context) error
}
