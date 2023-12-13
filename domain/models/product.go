package models

type Product struct {
	ID    int     `gorm:"primaryKey" json:"id"`
	Price float64 `gorm:"column:price" json:"price"`
	Name  string  `gorm:"column:name" json:"name"`
}
