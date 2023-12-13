package repositories

import (
	"github.com/peidrao/eulabs-api/domain/models"
	"gorm.io/gorm"
)

type dbProduct struct {
	Conn *gorm.DB
}

type ProductRepository interface {
	Create(product models.Product) error
}

func NewProductRepository(Conn *gorm.DB) ProductRepository {
	return &dbProduct{Conn: Conn}
}

func (db *dbProduct) Create(product models.Product) error {
	return db.Conn.Create(&product).Error
}
