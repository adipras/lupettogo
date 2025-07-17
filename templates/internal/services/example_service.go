package services

import (
	"{{.ProjectName}}/internal/models"
	"{{.ProjectName}}/internal/repositories"
)

type ExampleService struct {
	exampleRepo *repositories.ExampleRepository
}

func NewExampleService(exampleRepo *repositories.ExampleRepository) *ExampleService {
	return &ExampleService{
		exampleRepo: exampleRepo,
	}
}

func (s *ExampleService) GetExample() map[string]interface{} {
	return map[string]interface{}{
		"message": "Hello from LupettoGo! 🐺",
		"status":  "success",
		"data": map[string]interface{}{
			"example": "This is an example response from the service layer",
			"tips":    "Replace this service with your business logic",
		},
	}
}

func (s *ExampleService) GetAllExamples() ([]*models.Example, error) {
	return s.exampleRepo.FindAll()
}

func (s *ExampleService) GetExampleByID(id uint) (*models.Example, error) {
	return s.exampleRepo.FindByID(id)
}