package services

import (
	"fmt"

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
	GetById(productId int) utils.ProductResponse
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{productRepo: repositories.NewProductRepository(db)}
}

func (service *productService) Create(product models.Product) utils.Response {
	var response utils.Response
	if err := service.productRepo.Create(product); err != nil {
		response.Status = 200
		response.Messages = "Success to create a new product"
	}
	return response
}

func (service *productService) Delete(productId int) utils.Response {
	var response utils.Response
	if err := service.productRepo.Delete(productId); err != nil {
		response.Status = 400
		response.Messages = fmt.Sprint("Failed to delete product: ", productId)
	} else {
		response.Status = 200
		response.Messages = "Success to delete product"
	}
	return response
}

func (service *productService) Update(productId int, product models.Product) utils.Response {
	var response utils.Response
	if err := service.productRepo.Update(productId, product); err != nil {
		response.Status = 400
		response.Messages = fmt.Sprint("Failed to update product: ", productId)
	} else {
		response.Status = 200
		response.Messages = "Success to update product"
	}
	return response
}

func (service *productService) GetById(productId int) utils.ProductResponse {
	var response utils.ProductResponse
	data, _ := service.productRepo.Get(productId)
	response.Product = data
	return response
}
