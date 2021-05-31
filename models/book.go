package models

import "time"

type Book struct {
	ID          int
	Title       string `gorm:"title" json:"title"`
	Description string `gorm:"description" json:"description"`
	CreatedAt   time.Time
}

func (t *Book) TableName() string {
	return "books"
}
