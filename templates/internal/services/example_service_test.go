package services

import (
	"testing"

	"{{.ProjectName}}/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExampleRepository struct {
	mock.Mock
}

func (m *MockExampleRepository) FindAll() ([]*models.Example, error) {
	args := m.Called()
	return args.Get(0).([]*models.Example), args.Error(1)
}

func (m *MockExampleRepository) FindByID(id uint) (*models.Example, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Example), args.Error(1)
}

func (m *MockExampleRepository) Create(example *models.Example) (*models.Example, error) {
	args := m.Called(example)
	return args.Get(0).(*models.Example), args.Error(1)
}

func (m *MockExampleRepository) Update(example *models.Example) (*models.Example, error) {
	args := m.Called(example)
	return args.Get(0).(*models.Example), args.Error(1)
}

func (m *MockExampleRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestExampleService_GetExample(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockExampleRepository)

	// Create service
	service := NewExampleService(mockRepo)

	// Test GetExample
	result := service.GetExample()

	// Assertions
	assert.NotNil(t, result)
	assert.Equal(t, "Hello from LupettoGo! 🐺", result["message"])
	assert.Equal(t, "success", result["status"])
}

func TestExampleService_GetAllExamples(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockExampleRepository)

	// Set up mock expectations
	expectedExamples := []*models.Example{
		{ID: 1, Name: "Test 1", Email: "test1@example.com"},
		{ID: 2, Name: "Test 2", Email: "test2@example.com"},
	}
	mockRepo.On("FindAll").Return(expectedExamples, nil)

	// Create service
	service := NewExampleService(mockRepo)

	// Test GetAllExamples
	result, err := service.GetAllExamples()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedExamples, result)
	mockRepo.AssertExpectations(t)
}