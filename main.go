package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/peidrao/eulabs-api/config"
	"github.com/peidrao/eulabs-api/domain/controllers"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	apiV1 := e.Group("api/v1/")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	productController := controllers.NewProductController(db)

	apiV1.POST("product/", productController.Create)
	apiV1.PUT("product/:id", productController.Update)
	apiV1.DELETE("product/:id", productController.Delete)
	apiV1.GET("product/:id", productController.GetById)

	e.Logger.Fatal(e.Start(":8080"))
}
