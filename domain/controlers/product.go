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

	response := controller.productService.Create(
		models.Product{
			Name:  payloadValidator.Name,
			Price: payloadValidator.Price,
		},
	)

	return c.JSON(response.Status, response.Data)
}

func (controller ProductController) Update(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))

	type payload struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	response := controller.productService.Update(
		productId,
		models.Product{
			Name:  payloadValidator.Name,
			Price: payloadValidator.Price,
		},
	)

	return c.JSON(response.Status, response.Data)
}

func (controller ProductController) Delete(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	result := controller.productService.Delete(productId)

	return c.JSON(result.Status, result)
}

func (controller ProductController) GetById(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))

	response := controller.productService.GetById(productId)

	return c.JSON(response.Status, response.Data)
}
