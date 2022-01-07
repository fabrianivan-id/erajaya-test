package controllers

import (
	"github.com/fabrianivan21/erajaya-test/utils"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/fabrianivan21/erajaya-test/models"
)

type P map[string]interface{}

type ProductController struct {
	db utils.ProductModel
}

func NewProductController(db utils.ProductModel) *ProductController {
	return &ProductController{db: db}
}

func (pc *ProductController) AddProduct(c echo.Context) error {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	product, err := pc.db.AddProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusCreated, P{
		"status": "success",
		"data":   product,
	})
}
func (pc *ProductController) GetLatestProduct(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetLatestProduct()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, P{
		"status": "success",
		"data":   products,
	})
}
func (pc *ProductController) GetProductByLowestPrice(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetProductLowestPrice()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, P{
		"status": "success",
		"data":   products,
	})

}
func (pc *ProductController) GetProductByHighestPrice(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetProductHighestPrice()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, P{
		"status": "success",
		"data":   products,
	})

}

func (pc *ProductController) GetProductByAtoZ(c echo.Context) error {

	var products []models.Product

	products, err := pc.db.GetProductASC()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, P{
		"status": "success",
		"data":   products,
	})

}
func (pc *ProductController) GetProductByZtoA(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetProductDESC()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, P{
		"status": "success",
		"data":   products,
	})

}