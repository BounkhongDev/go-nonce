package services

import (
	"gorm.io/gorm"
)

type ExampleService interface {
	// Insert your function interface
}

type exampleService struct {
	db *gorm.DB
}

func NewExampleService(db *gorm.DB) ExampleService {
	return &exampleService{db: db}
}