package repository

import (
	"golangbe/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	List() ([]models.Book, error)
	Store(models.Book) (bool, error)
}

type bookRepository struct {
	db *gorm.DB
}

func InitBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (repository *bookRepository) List() ([]models.Book, error) {
	books := []models.Book{}

	err := repository.db.Find(&books).Error

	if err != nil {
		return books, err
	}

	return books, nil
}

func (repository *bookRepository) Store(dataInput models.Book) (bool, error) {
	err := repository.db.Create(&dataInput).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
