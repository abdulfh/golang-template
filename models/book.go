package models

import "time"

type Book struct {
	ID          int
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	CreatedAt   time.Time
}

func (t *Book) TableName() string {
	return "books"
}

type BookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
