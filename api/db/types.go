package db

import (
	"time"
)

type Product struct {
	Name  string  `json:"name" form:"name" query:"name"`
	Price float32 `json:"price" form:"price" query:"price"`
}

type ProductDTO struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	Price     float32   `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
