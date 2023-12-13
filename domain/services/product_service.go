package services

import (
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
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{productRepo: repositories.NewProductRepository(db)}
}

func (service *productService) Create(product models.Product) utils.Response {
	var response utils.Response
	if err := service.productRepo.Create(product); err != nil {
		response.Status = 400
		response.Messages = "Failed to create a new product"
	} else {
		response.Status = 200
		response.Messages = "Success to create a new product"
	}
	return response
}
