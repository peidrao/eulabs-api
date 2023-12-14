package services

import (
	"net/http"

	"github.com/peidrao/eulabs-api/domain/models"
	"github.com/peidrao/eulabs-api/domain/repositories"
	"github.com/peidrao/eulabs-api/utils"
	"gorm.io/gorm"
)

type productService struct {
	productRepo repositories.ProductRepository
}

type ProductService interface {
	Create(product models.Product) utils.Response
	Delete(productId int) utils.Response
	Update(productId int, product models.Product) utils.Response
	GetById(productId int) utils.Response
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{productRepo: repositories.NewProductRepository(db)}
}

func (service *productService) Create(product models.Product) utils.Response {
	var response utils.Response
	data, err := service.productRepo.Create(product)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
	} else {
		response.Data = data
		response.Status = http.StatusCreated

	}

	return response
}

func (service *productService) Delete(productId int) utils.Response {
	var response utils.Response

	if err := service.productRepo.Delete(productId); err != nil {
		response.Status = http.StatusNotFound
		response.Data = err.Error()
	} else {
		response.Status = http.StatusNoContent
	}
	return response
}

func (service *productService) Update(productId int, product models.Product) utils.Response {
	var response utils.Response
	if err := service.productRepo.Update(productId, product); err != nil {
		response.Status = http.StatusBadRequest
		response.Data = err.Error()
	} else {
		productUpdated := service.GetById(productId)
		response.Status = http.StatusOK
		response.Data = productUpdated.Data
	}
	return response
}

func (service *productService) GetById(productId int) utils.Response {
	var response utils.Response
	product, err := service.productRepo.Get(productId)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusNotFound
	} else {
		response.Data = product
		response.Status = http.StatusOK

	}
	return response
}
