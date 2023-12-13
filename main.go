package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/peidrao/eulabs-api/config"
	"github.com/peidrao/eulabs-api/domain/controlers"
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

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	productController := controlers.NewProductController(db)
	apiV1.POST("product/", productController.Create)

	e.Logger.Fatal(e.Start(":8080"))
}
