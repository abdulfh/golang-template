package services

import (
	"golangbe/models"
	"golangbe/repository"
)

type BookServices interface {
	List() ([]models.Book, error)
	Store(models.Book) (bool, error)
}

type bookServices struct {
	repository repository.BookRepository
}

func InitBookServices(repository repository.BookRepository) *bookServices {
	return &bookServices{repository}
}

func (services *bookServices) List() ([]models.Book, error) {
	result, err := services.repository.List()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (services *bookServices) Store(dataInput models.Book) (bool, error) {
	result, err := services.repository.Store(dataInput)
	if err != nil {
		return result, err
	}
	return result, nil
}
