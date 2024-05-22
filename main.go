package main

import (
	"log"
	"minimarket_mikti/app"
	"minimarket_mikti/controller"
	"minimarket_mikti/helper"
	"minimarket_mikti/repository"
	"minimarket_mikti/service"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file!")
	}

	r := echo.New()

	r.Validator = &CustomValidator{validator: validator.New()}
	r.HTTPErrorHandler = helper.BindAndValidate

	db := app.DBConnection()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	r.POST("/product", productController.CreateProduct)

	r.Logger.Fatal(r.Start(":" + os.Getenv("PORT")))
}
