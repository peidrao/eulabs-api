package controlers

import (
	"net/http"
	"strconv"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/peidrao/eulabs-api/domain/models"
	"github.com/peidrao/eulabs-api/domain/services"
	"gorm.io/gorm"
)

type ProductController struct {
	productService services.ProductService
	validate       vl.Validate
}

func NewProductController(db *gorm.DB) ProductController {
	service := services.NewProductService(db)
	controller := ProductController{
		productService: service,
		validate:       *vl.New(),
	}

	return controller
}

func (controller ProductController) Create(c echo.Context) error {
	type payload struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	controller.productService.Create(
		models.Product{
			Name:  payloadValidator.Name,
			Price: payloadValidator.Price,
		},
	)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": "",
	})
}

func (controller ProductController) Update(c echo.Context) error {
	type payload struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	idItem, _ := strconv.Atoi(c.Param("id"))
	result := controller.productService.Update(
		idItem,
		models.Product{
			Name:  payloadValidator.Name,
			Price: payloadValidator.Price,
		},
	)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

func (controller ProductController) Delete(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	result := controller.productService.Delete(productId)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

func (controller ProductController) GetById(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))

	result := controller.productService.GetById(productId)

	return c.JSON(http.StatusOK, result)
}
