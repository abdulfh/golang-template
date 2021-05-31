package services

import "golangbe/repository"

type BookServices interface{}

type bookServices struct {
	repository repository.BookRepository
}

func InitBookServices(repository repository.BookRepository) *bookServices {
	return &bookServices{repository}
}
