package book

import (
	"time"
)

type Book struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
