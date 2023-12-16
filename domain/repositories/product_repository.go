package repositories

import (
	"errors"

	"github.com/peidrao/eulabs-api/domain/models"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	Create(product models.Product) (models.Product, error)
	Delete(productId int) error
	Update(productId int, product models.Product) error
	Get(productId int) (models.Product, error)
	HandleDBError(err error, defaultMessage string) error
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) Create(product models.Product) (models.Product, error) {
	err := p.db.Create(&product).Error
	return product, p.HandleDBError(err, "Erro ao criar um novo produto.")
}

func (p *productRepository) Delete(productId int) error {
	var existingProduct models.Product
	err := p.findProductByID(productId, &existingProduct)
	if err != nil {
		return p.HandleDBError(err, "Produto não foi encontrado")
	}

	err = p.db.Delete(&models.Product{ID: productId}).Error
	return p.HandleDBError(err, "Erro ao remover produto.")
}

func (p *productRepository) Update(productId int, product models.Product) error {
	var existingProduct models.Product
	err := p.findProductByID(productId, &existingProduct)
	if err != nil {
		return p.HandleDBError(err, "Produto não foi encontrado")
	}

	err = p.db.Where("id", productId).Updates(product).Error
	return p.HandleDBError(err, "Erro ao atualizar produto.")
}

func (p *productRepository) Get(productId int) (models.Product, error) {
	var product models.Product
	err := p.findProductByID(productId, &product)
	return product, p.HandleDBError(err, "Produto não foi encontrado")
}

func (p *productRepository) findProductByID(productId int, product *models.Product) error {
	err := p.db.Where("id", productId).First(product).Error
	return err
}

func (p *productRepository) HandleDBError(err error, defaultMessage string) error {
	if err != nil {
		return errors.New(defaultMessage)
	}
	return nil
}
