package router

import (
	"github.com/fabrianivan21/erajaya-test/api/controllers"
	"github.com/fabrianivan21/erajaya-test/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	pdb := utils.NewProductDB(db)
	pc := controllers.NewProductController(pdb)

	//tambah product
	e.POST("/products", pc.AddProduct)
	e.GET("/products", pc.GetProductByAtoZ)
	e.GET("/products/desc", pc.GetProductByZtoA)
	e.GET("/products/latest", pc.GetLatestProduct)

	e.GET("/products/lowest", pc.GetProductByLowestPrice)
	e.GET("/products/highest", pc.GetProductByHighestPrice)

	return e
}
