package repositories

import (
	"github.com/peidrao/eulabs-api/domain/models"
	"gorm.io/gorm"
)

type dbProduct struct {
	Conn *gorm.DB
}

type ProductRepository interface {
	Create(product models.Product) (models.Product, error)
	Delete(productId int) error
	Update(productId int, product models.Product) error
	Get(productId int) (models.Product, error)
}

func NewProductRepository(Conn *gorm.DB) ProductRepository {
	return &dbProduct{Conn: Conn}
}

func (db *dbProduct) Create(product models.Product) (models.Product, error) {
	result := db.Conn.Create(&product)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

func (db *dbProduct) Delete(productId int) error {
	var existingProduct models.Product
	err := db.Conn.First(&existingProduct, productId).Error
	if err != nil {
		return err
	}

	return db.Conn.Delete(&models.Product{ID: productId}).Error
}

func (db *dbProduct) Update(productId int, product models.Product) error {
	return db.Conn.Where("id", productId).Updates(product).Error
}

func (db *dbProduct) Get(productId int) (models.Product, error) {
	var data models.Product
	result := db.Conn.Where("id", productId).First(&data)
	return data, result.Error
}
