package services

import (
	"{{.ProjectName}}/internal/repositories"
	"gorm.io/gorm"
)

type Services struct {
	Example *ExampleService
}

func New(db *gorm.DB) *Services {
	repos := repositories.New(db)
	
	return &Services{
		Example: NewExampleService(repos.Example),
	}
}