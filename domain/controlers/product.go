package controlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/peidrao/eulabs-api/domain/models"
	"github.com/peidrao/eulabs-api/domain/services"
	"gorm.io/gorm"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(db *gorm.DB) ProductController {
	service := services.NewProductService(db)
	controller := ProductController{
		productService: service,
	}

	return controller
}

func (controller ProductController) Create(c echo.Context) error {
	type payload struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if payloadValidator.Name == "" || payloadValidator.Price == 0.0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Name and Price cannot be empty",
		})
	}

	result := controller.productService.Create(
		models.Product{
			Name:  payloadValidator.Name,
			Price: payloadValidator.Price,
		},
	)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": result,
	})
}
